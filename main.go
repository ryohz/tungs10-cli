package main

import (
	"fmt"
	"os"
	"ryozk/tungs10-cli/extra"

	"github.com/manifoldco/promptui"
)

var commandLineArgsNum int = len(os.Args) - 1

func main() {
	// if ~/.tungs10 does not exist, create it
	appRootDir, gardErr := extra.GetAppRootDir()
	if gardErr != nil {
		panic(gardErr)
	}
	if _, err := os.Stat(appRootDir); os.IsNotExist(err) {
		if err := os.Mkdir(appRootDir, 0755); err != nil {
			panic(err)
		}
	}
	promptOfMode := promptui.Select{
		Label: "select mode",
		Items: []string{"Add", "List", "Remove", "Information"},
	}
	_, mode, err := promptOfMode.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if mode == "Add" {
		promptOfAdd := promptui.Select{
			Label: "select thing to add",
			Items: []string{"user", "credential"},
		}
		_, targetOfAdd, err := promptOfAdd.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		if targetOfAdd == "user" {
			addUser()
			return
		}
		if targetOfAdd == "credential" {
			addCredential()
			return
		}
	}

	if mode == "List" {
		promptOfList := promptui.Select{
			Label: "select thing to list",
			Items: []string{"user", "credential"},
		}
		_, targetOfList, err := promptOfList.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if targetOfList == "user" {
			listUsers()
			return
		}
		if targetOfList == "credential" {
			listCredentials()
			return
		}
	}

	if mode == "Remove" {
		promptOfRemove := promptui.Select{
			Label: "select thing to remove",
			Items: []string{"user", "credential"},
		}
		_, targetOfRemove, err := promptOfRemove.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		if targetOfRemove == "user" {
			removeUser()
			return
		}
		if targetOfRemove == "credential" {
			removeCredential()
			return
		}
	}

	if mode == "Information" {
		PrintInformations()
		return
	}

	return
}
