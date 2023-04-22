package main

import (
	"fmt"
	"ryozk/tungs10-cli/credential"
	"ryozk/tungs10-cli/user"

	"github.com/manifoldco/promptui"
)

func addCredential() {
	userList, gulErr := user.GetUserList()
	if gulErr != nil {
		fmt.Println(gulErr)
		return
	}
	if len(userList) == 0 {
		fmt.Println("no user found")
		return
	}

	promptOfLoginUsername := promptui.Prompt{
		Label: "Username",
	}
	loginUsername, rpoluErr := promptOfLoginUsername.Run()
	if rpoluErr != nil {
		fmt.Printf("Prompt failed %v\n", rpoluErr)
		return
	}

	promptOfLoginPassword := promptui.Prompt{
		Label: "Password",
		Mask:  '*',
	}
	loginPassword, rpolpErr := promptOfLoginPassword.Run()
	if rpolpErr != nil {
		fmt.Printf("Prompt failed %v\n", rpolpErr)
		return
	}

	var loginUser *user.User = user.New()
	loginUser.SetUsername(loginUsername)
	loginUser.SetPassword(loginPassword)
	isLogined := loginUser.Login()
	if !isLogined {
		fmt.Println("failed to login")
		return
	}
	fmt.Println("logged in successfully")
	loginUser.CalcAndSetKey()
	var rasiErr error = loginUser.ReadAndSetIv()
	if rasiErr != nil {
		fmt.Println(rasiErr)
		return
	}
	var serviceName string
	var username string
	var email string
	var password string
input:
	for {
		promptOfServiceName := promptui.Prompt{
			Label: "Service Name",
		}
		serviceNameResult, rposErr := promptOfServiceName.Run()
		if rposErr != nil {
			fmt.Printf("Prompt failed %v\n", rposErr)
			return
		}
		serviceName = serviceNameResult

		promptOfUsername := promptui.Prompt{
			Label: "Username",
		}
		usernameResult, rpuErr := promptOfUsername.Run()
		if rpuErr != nil {
			fmt.Printf("Prompt failed %v\n", rpuErr)
			return
		}
		username = usernameResult

		promptOfEmail := promptui.Prompt{
			Label: "Email",
		}
		emailResult, rpeErr := promptOfEmail.Run()
		if rpeErr != nil {
			fmt.Printf("Prompt failed %v\n", rpeErr)
			return
		}
		email = emailResult

		promptOfPassword := promptui.Prompt{
			Label: "Password",
			Mask:  '*',
		}
		passwordResult, rppErr := promptOfPassword.Run()
		if rppErr != nil {
			fmt.Printf("Prompt failed %v\n", rppErr)
			return
		}
		password = passwordResult

	sure:
		for {
			prompt := promptui.Prompt{
				Label:     "Are you sure you want to save this credential",
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
			continue sure
		}
	}
	var newCredential *credential.Credential = credential.New()
	newCredential.SetUser(*loginUser)
	newCredential.SetServiceName(serviceName)
	newCredential.SetUsername(username)
	newCredential.SetEmail(email)
	newCredential.SetPassword(password)
	var eErr error = newCredential.EncryptAndSetCredentialInformations()
	if eErr != nil {
		fmt.Println(eErr)
		return
	}
	var acErr error = newCredential.AddToDB()
	if acErr != nil {
		fmt.Println(acErr)
	}
	return
}
