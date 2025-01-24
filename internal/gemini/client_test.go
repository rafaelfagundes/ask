package gemini

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/google/generative-ai-go/genai"
	"github.com/rafaelfagundes/ask/internal/osinfo"
)

type mockGenerativeModel struct {
	response string
	err      error
}

func (m *mockGenerativeModel) GenerateContent(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error) {
	if m.err != nil {
		return nil, m.err
	}
	return &genai.GenerateContentResponse{
		Candidates: []*genai.Candidate{
			{
				Content: &genai.Content{
					Parts: []genai.Part{
						genai.Text(m.response),
					},
				},
			},
		},
	}, nil
}

func TestNewClient_NoAPIKey(t *testing.T) {
	os.Setenv("GEMINI_API_KEY", "")
	defer os.Unsetenv("GEMINI_API_KEY")

	_, err := NewClient(osinfo.Get())
	if err == nil {
		t.Error("Expected error when GEMINI_API_KEY is not set")
	}
}

func TestNewClient(t *testing.T) {
	os.Setenv("GEMINI_API_KEY", "test-api-key")
	defer os.Unsetenv("GEMINI_API_KEY")

	client, err := NewClient(osinfo.Get())
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if client == nil {
		t.Error("Expected client to be created")
	}
}

func TestGenerateContent(t *testing.T) {
	mockResponse := "This is a test response"
	client := &Client{
		client: nil,
		model: &mockGenerativeModel{
			response: mockResponse,
		},
	}

	tests := []struct {
		name    string
		prompt  string
		want    string
		wantErr bool
		err     error
	}{
		{
			name:    "Valid prompt",
			prompt:  "Hello, world!",
			want:    mockResponse,
			wantErr: false,
		},
		{
			name:    "Error case",
			prompt:  "error",
			wantErr: true,
			err:     errors.New("test error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				client.model = &mockGenerativeModel{err: tt.err}
			} else {
				client.model = &mockGenerativeModel{response: tt.want}
			}

			ctx := context.Background()
			response, err := client.GenerateContent(ctx, tt.prompt)

			if tt.wantErr {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if response != tt.want {
				t.Errorf("Got %q, want %q", response, tt.want)
			}
		})
	}
}
