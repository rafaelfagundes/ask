package gemini

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

	"github.com/rafaelfagundes/ask/internal/osinfo"
)

type GenerativeModelInterface interface {
	GenerateContent(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error)
}

type Client struct {
	client *genai.Client
	model  GenerativeModelInterface
}

func NewClient(osInfo *osinfo.OSInfo) (*Client, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY environment variable not set")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	model := client.GenerativeModel("gemini-2.0-flash-thinking-exp-01-21")
	model.SetTemperature(0.7)
	model.SetTopK(32)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(4096)

	systemPrompt := fmt.Sprintf(`You are a terminal-based assistant. Current environment:
- OS: %s %s
- Shell: %s
- Terminal: %s

Provide concise answers with:
- Clear line breaks at 80 characters max
- No unnecessary intros/outros
- REQUIRED: all response should be formatted in markdown
- Do not wrap the response with markdown code fences 
- Separate sections with clear headings
- Wrap long lines smartly
- When showing command examples, use the appropriate syntax for the current OS/shell`,
		osInfo.OS, osInfo.Version, osInfo.Shell, osInfo.Terminal)

	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(systemPrompt),
		},
	}

	return &Client{
		client: client,
		model:  model,
	}, nil
}

func (c *Client) GenerateContent(ctx context.Context, question string) (string, error) {
	resp, err := c.model.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		return "", err
	}

	responseText := buildResponseText(resp)
	return cleanMarkdownFences(responseText), nil
}

func buildResponseText(resp *genai.GenerateContentResponse) string {
	var sb strings.Builder
	for _, cand := range resp.Candidates {
		for _, part := range cand.Content.Parts {
			sb.WriteString(fmt.Sprintf("%v", part))
		}
	}
	return sb.String()
}

func cleanMarkdownFences(response string) string {
	lines := strings.Split(response, "\n")
	if len(lines) < 2 {
		return response
	}

	firstLine := strings.TrimSpace(lines[0])
	lastLine := strings.TrimSpace(lines[len(lines)-1])

	if (firstLine == "```markdown" || firstLine == "```md") && lastLine == "```" {
		return strings.Join(lines[1:len(lines)-1], "\n")
	}
	return response
}

func (c *Client) Close() error {
	return c.client.Close()
}
