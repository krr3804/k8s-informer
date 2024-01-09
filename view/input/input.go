package input

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func ScreenPrompt(contexts []string) (string, string, string) {

	selectedProvider := getProvider()

	selectedContext := getContext(contexts)

	fileName := getFileName()

	return selectedProvider, selectedContext, fileName
}

func getProvider() string {
	prompt := promptui.Select{
		Label: "Select Provider",
		Items: []string{"EKS[AWS]", "Local"}}

	_, provider, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return provider
}

func getContext(contexts []string) string {
	prompt := promptui.Select{
		Label: "Select Context",
		Items: contexts,
	}

	_, context, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return context
}

func getFileName() string {
	prompt := promptui.Prompt{
		Label:   "Write file name",
		Default: "output",
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}
