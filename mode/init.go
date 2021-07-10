package mode

import (
	"fmt"
	"os"

	"github.com/Pauloo27/gommit/config"
	"github.com/Pauloo27/gommit/prefix"
	"github.com/Pauloo27/gommit/utils"
	"github.com/c-bata/go-prompt"
)

func initCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	for _, prefixPack := range prefix.Packs {
		s = append(s, prompt.Suggest{
			Text: prefixPack.Name,
			Description: fmt.Sprintf("%s %s %s",
				prefixPack.GetPrefix("feat"), prefixPack.GetPrefix("fix"), prefixPack.GetPrefix("ci"),
			),
		})
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func Init() {
	name := utils.Prompt("Prefix pack: ", initCompleter)

	pack := prefix.GetPrefixPack(name)
	if pack == nil {
		fmt.Println("Prefix pack not found =(")
		os.Exit(-1)
	}
	// TODO: store config
	c := config.Config{
		PrefixPack: pack.Name,
	}
	err := config.StoreConfig(&c, "./.gommitrc.json")
	if err != nil {
		fmt.Println("Something went wrong while storing the config", err)
		os.Exit(-1)
	}
}
