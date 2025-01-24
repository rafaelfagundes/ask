package cli

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/rafaelfagundes/ask/internal/app"
	"github.com/rafaelfagundes/ask/internal/spinner"
)

func Run(a *app.App, args []string) error {
	if len(args) == 0 {
		return handleMainCommand(a, args)
	}

	switch args[0] {
	case "history":
		return handleHistoryCommand(a, args[1:])
	case "last":
		return handleLastCommand(a, args[1:])
	case "-h", "--help":
		printUsage()
		return nil
	case "-c":
		fmt.Println(a.Config.Dir())
		return nil
	default:
		return handleMainCommand(a, args)
	}
}

func handleMainCommand(a *app.App, args []string) error {
	fs := flag.NewFlagSet("ask", flag.ExitOnError)
	noPager := fs.Bool("no-pager", false, "Disable paginated output")
	showConfig := fs.Bool("c", false, "Show configuration directory location")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *showConfig {
		fmt.Println(a.Config.Dir())
		return nil
	}

	var question string
	if fs.NArg() > 0 {
		question = strings.Join(fs.Args(), " ")
	} else {
		fmt.Print("Enter your question: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("error reading question: %v", err)
		}
		question = strings.TrimSpace(input)
	}

	ctx := context.Background()
	s := spinner.New()
	s.Start()
	response, err := a.Gemini.GenerateContent(ctx, question)
	s.Stop()
	if err != nil {
		return fmt.Errorf("error generating content: %v", err)
	}

	if err := a.History.Save(question, response); err != nil {
		log.Printf("Warning: failed to save history: %v", err)
	}

	return showResponse(response, *noPager)
}

func handleHistoryCommand(a *app.App, args []string) error {
	if len(args) > 0 && args[0] == "delete" {
		return handleDeleteCommand(a, args[1:])
	}

	switch {
	case len(args) == 0:
		return showHistoryList(a)
	case len(args) == 1:
		if pos, err := strconv.Atoi(args[0]); err == nil {
			return showHistoryEntry(a, pos)
		}
		return fmt.Errorf("invalid history number")
	default:
		return fmt.Errorf("invalid history command")
	}
}

func handleLastCommand(a *app.App, args []string) error {
	fs := flag.NewFlagSet("last", flag.ExitOnError)
	noPager := fs.Bool("no-pager", false, "Disable paginated output")

	if err := fs.Parse(args); err != nil {
		return err
	}

	entry, err := a.History.GetLast()
	if err != nil {
		return fmt.Errorf("error retrieving last response: %v", err)
	}

	if entry == nil {
		fmt.Println("No history entries found")
		return nil
	}

	return showResponse(entry.Response, *noPager)
}

func handleDeleteCommand(a *app.App, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing delete argument. Usage: ask history delete <position|all>")
	}

	switch args[0] {
	case "all":
		return confirmAndDeleteAll(a)
	default:
		pos, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid position number")
		}
		return deletePosition(a, pos)
	}
}

func showHistoryList(a *app.App) error {
	entries, err := a.History.List()
	if err != nil {
		return fmt.Errorf("error retrieving history: %v", err)
	}

	if len(entries) == 0 {
		fmt.Println("No history entries found")
		return nil
	}

	var builder strings.Builder
	for _, entry := range entries {
		fmt.Fprintf(&builder, "%4d. [%s] %s\n",
			entry.Position,
			entry.Timestamp.Format("2006-01-02 15:04"),
			entry.Question)
	}

	cmd := exec.Command("less")
	cmd.Stdin = strings.NewReader(builder.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func showHistoryEntry(a *app.App, pos int) error {
	entry, err := a.History.Get(pos)
	if err != nil {
		return fmt.Errorf("error retrieving history entry: %v", err)
	}

	return showResponse(entry.Response, false)
}

func confirmAndDeleteAll(a *app.App) error {
	fmt.Print("Are you sure you want to delete ALL history? (y/N): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if strings.ToLower(scanner.Text()) != "y" {
		fmt.Println("Deletion canceled")
		return nil
	}

	if err := a.History.DeleteAll(); err != nil {
		return fmt.Errorf("error deleting history: %v", err)
	}

	fmt.Println("All history entries deleted")
	return nil
}

func deletePosition(a *app.App, pos int) error {
	if err := a.History.Delete(pos); err != nil {
		return fmt.Errorf("error deleting entry: %v", err)
	}

	fmt.Printf("Deleted entry at position %d\n", pos)
	return nil
}

func showResponse(response string, noPager bool) error {
	if len(strings.TrimSpace(response)) < 256 {
		noPager = true
	}

	args := []string{"-"}
	if !noPager {
		args = append([]string{"-p"}, args...)
	}

	cmd := exec.Command("glow", args...)
	cmd.Stdin = strings.NewReader(response)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("\nError running glow: %v\n", err)
		fmt.Println("\nRaw response:\n" + response)
	}
	return nil
}

func printUsage() {
	fmt.Println(`Usage: ask [command] [options] [question]

Commands:
  history         Show command history
  history delete  Delete history entries
  last            Show last response
  
Options:`)
	// Add flag descriptions
	fmt.Println(`
Examples:
  ask -c  # Show config directory
  ask "What is Go?"
  ask --no-pager "How do I create a file in Go?"`)
}
