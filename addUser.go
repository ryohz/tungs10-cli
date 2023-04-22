package main

import (
	"errors"
	"fmt"
	"ryozk/tungs10-cli/crypto"
	"ryozk/tungs10-cli/user"

	"github.com/manifoldco/promptui"
)

func addUser() {
	fmt.Println("creating a new user...")
	var newUser *user.User = user.New()
	var username string
	var password string
input:
	for {
		usernameValidate := func(input string) error {
			if len(input) < 1 {
				return errors.New("Username must have more than a characters")
			}
			return nil
		}
		prompt := promptui.Prompt{
			Label:    "Username",
			Validate: usernameValidate,
		}
		usernameResult, err := prompt.Run()
		username = usernameResult
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		passwordValidate := func(input string) error {
			if len(input) < 5 {
				return errors.New("Password must have more than 5 characters")
			}
			if len(input) > 32 {
				return errors.New("Password must have less than 32 characters")
			}
			return nil
		}
		PasswordPrompt := promptui.Prompt{
			Label:    "Password",
			Validate: passwordValidate,
			Mask:     '*',
		}
		passwordResult, err := PasswordPrompt.Run()
		password = passwordResult
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
	invalidAllowRe:
		for {
			prompt := promptui.Prompt{
				Label:     "Are you sure you want to add this user?",
				IsConfirm: true,
			}
			isSure, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			if isSure == "n" || isSure == "N" {
				continue input
			}
			if isSure == "y" || isSure == "Y" {
				break input
			}
			continue invalidAllowRe
		}
	}
	newUser.SetUsername(username)
	newUser.SetPassword(password)
	iv, err := crypto.GenerateIV()
	if err != nil {
		fmt.Println(err)
		return
	}
	newUser.SetIv(iv)
	var cuErr error = newUser.CreateUser()
	if cuErr != nil {
		fmt.Println(cuErr)
		return
	}
	return
}
