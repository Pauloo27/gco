package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

func Prompt(prefix string, completer prompt.Completer, options ...prompt.Option) string {
	opts := []prompt.Option{
		prompt.OptionDescriptionTextColor(prompt.White),
		prompt.OptionDescriptionBGColor(prompt.Black),

		prompt.OptionSuggestionTextColor(prompt.White),
		prompt.OptionSuggestionBGColor(prompt.Black),

		prompt.OptionSelectedSuggestionTextColor(prompt.LightGray),
		prompt.OptionSelectedSuggestionBGColor(prompt.Black),

		prompt.OptionSelectedDescriptionTextColor(prompt.Black),
		prompt.OptionSelectedDescriptionBGColor(prompt.LightGray),

		prompt.OptionAddKeyBind(prompt.KeyBind{
			Key: prompt.ControlC,
			Fn: func(b *prompt.Buffer) {
				fmt.Println("Ctrl C was pressed. Exiting...")
				os.Exit(-1)
			},
		}),

		prompt.OptionShowCompletionAtStart(),
	}
	opts = append(opts, options...)
	return prompt.Input(
		prefix, completer,
		opts...,
	)
}

func LowerCasedPrompt(prefix string, completer prompt.Completer, options ...prompt.Option) string {
	return strings.ToLower(Prompt(prefix, completer, options...))
}

func EmptyCompleter(d prompt.Document) []prompt.Suggest {
	return nil
}

func YesOrNoCompleter(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{
		{Description: "yes", Text: "y"},
		{Description: "no", Text: "n"},
	}
}

func PromptYesOrNot(prompt string) string {
	input := LowerCasedPrompt(
		" Skip CI? [y/N] ",
		YesOrNoCompleter,
	)
	if input == "yes" || input == "y" {
		return "yes"
	}
	if input == "no" || input == "n" {
		return "no"
	}
	return input

}
