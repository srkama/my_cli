package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

type changeType struct {
	Code        string
	Description string
}

func main() {
	flag.Parse()
	message := flag.Args()
	var components []string

	changeTypes := []changeType{
		{Code: "feat", Description: "change is related to new feature"},
		{Code: "fix", Description: "change is related to a bug fix"},
		{Code: "docs", Description: "changes are related to documentation"},
		{Code: "refactor", Description: "changes are related to refactoring of the code"},
		{Code: "tests", Description: "unit tests changes"},
		{Code: "style", Description: "format related changes"},
		{Code: "perf", Description: "performance improvements"},
	}

	promptTemplates :=
		&promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "\U0001F336 {{ .Code | cyan }} ({{ .Description | red }})",
			Inactive: "  {{ .Code | cyan }} ({{ .Description | red }})",
			Selected: "\U0001F336 {{ .Code | red | cyan }}",
		}

	if len(message) == 0 {
		log.Fatal("message is mandatory")
	}

	prompt := promptui.Select{
		Label:     "Select Type",
		Items:     changeTypes,
		Size:      10,
		Templates: promptTemplates,
	}

	is_present := strings.Index(message[0], "feat")

	commitMessage := strings.Join(message, " ")

	if is_present == -1 {
		i, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		commitMessage = changeTypes[i].Code + " : " + strings.Join(message, " ")
	}
	components = append(components, "commit", "-m")
	components = append(components, commitMessage)
	fmt.Println("this from my cli git", commitMessage, components)
	// Execute the command
	out, err := exec.Command("git", components...).CombinedOutput()
	fmt.Println("output", string(out), err)
}
