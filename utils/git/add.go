package git

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Pauloo27/gco/utils"
	"github.com/c-bata/go-prompt"
)

func parseInstruction(instruction string, filesMap map[string]bool, files []string) (map[string]bool, error) {
	newFilesMap := filesMap
	if instruction == "*" {
		for file := range newFilesMap {
			newFilesMap[file] = true
		}
		return newFilesMap, nil
	}
	splittedString := strings.Split(instruction, "-")
	if len(splittedString) == 1 {
		entry := splittedString[0]
		// not a range
		value := true
		if strings.HasPrefix(splittedString[0], "^") {
			entry = strings.TrimPrefix(entry, "^")
			value = false
		}
		number, err := strconv.Atoi(entry)
		if err != nil {
			return newFilesMap, err
		}
		newFilesMap[files[number-1]] = value
		return newFilesMap, nil
	} else if len(splittedString) == 2 {
		value := true
		entry := splittedString
		if strings.HasPrefix(entry[0], "^") {
			value = false
			entry = []string{strings.TrimPrefix(splittedString[0], "^"), splittedString[1]}
		}
		// range
		startStr, endStr := entry[0], entry[1]
		start, err := strconv.Atoi(startStr)
		if err != nil {
			return newFilesMap, err
		}
		end, err := strconv.Atoi(endStr)
		if err != nil {
			return newFilesMap, err
		}
		for i := start; i <= end; i++ {
			newFilesMap[files[i-1]] = value
		}
		return newFilesMap, nil
	}
	return newFilesMap, errors.New("Invalid instruction")
}

func PromptFilesToAdd(out prompt.ConsoleWriter, filesName []string) {
	for {
		input := utils.Prompt(
			" Files to commit (enter ? for help) ", utils.IndexedCompleter(filesName),
		)

		if input == "?" {
			utils.PrettyPrint(out, "To commit files that are already added: ",
				prompt.Blue, "press enter", prompt.DefaultColor, "\n",
			)
			utils.PrettyPrint(out, "To select files 1 and 3: ", prompt.Blue,
				"1 3", prompt.DefaultColor, "\n",
			)
			utils.PrettyPrint(out,
				"To select files 1, 2, 3 and 4: ", prompt.Blue, "1-4", prompt.DefaultColor,
				" (", prompt.Blue, "1 2 3 4", prompt.DefaultColor, " also works)\n",
			)
			utils.PrettyPrint(out, "To select all files: ", prompt.Blue, "*",
				prompt.DefaultColor, "\n",
			)
			utils.PrettyPrint(out, "To select all files expect 3: ", prompt.Blue,
				"^3", prompt.DefaultColor, "\n",
			)
			continue
		}

		filesToCommit := strings.Split(input, " ")
		if len(filesToCommit) != 0 {
			for _, file := range filesToCommit {
				fmt.Println(file)
			}
		}
		break
	}
}
