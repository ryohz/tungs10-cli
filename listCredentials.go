package main

import (
	"fmt"
	"ryozk/tungs10-cli/credential"
	"ryozk/tungs10-cli/extra"
	"ryozk/tungs10-cli/user"

	"github.com/manifoldco/promptui"
)

func listCredentials() {
	userList, gulErr := extra.GetUserList()
	if gulErr != nil {
		fmt.Println(gulErr)
		return
	}
	if len(userList) == 0 {
		fmt.Println("you don't have any user, yet")
		return
	}

	promptOfUsername := promptui.Prompt{
		Label: "Username",
	}
	username, rpouErr := promptOfUsername.Run()
	if rpouErr != nil {
		fmt.Printf("Prompt failed: %v\n", rpouErr)
		return
	}

	promptOfPassword := promptui.Prompt{
		Label: "Password",
		Mask:  '*',
	}
	password, rpopErr := promptOfPassword.Run()
	if rpopErr != nil {
		fmt.Printf("Prompt failed: %v\n", rpopErr)
		return
	}

	var loginUser *user.User = user.New()
	loginUser.SetUsername(username)
	loginUser.SetPassword(password)
	loginUser.CalcAndSetKey()
	var rasiErr error = loginUser.ReadAndSetIv()
	if rasiErr != nil {
		fmt.Println(rasiErr)
		return
	}
	var allowedLogin bool = loginUser.Login()
	if !allowedLogin {
		fmt.Println("failed to login")
		return
	}
	var isCredentialExists bool = loginUser.IsCredentialExists()
	if !isCredentialExists {
		fmt.Println("you don't have any credential")
		return
	}
	lastId, gliErr := credential.GetLastId(*loginUser)
	if gliErr != nil {
		fmt.Println(gliErr)
		return
	}
	for i := 1; i <= lastId; i++ {
		var newCredential *credential.Credential = credential.New()
		newCredential.SetId(i)
		newCredential.SetUser(*loginUser)
		var recErr error = newCredential.ReadAndSetEncryptedCredentialInformations()
		if recErr != nil {
			continue
		}
		var dErr error = newCredential.DecryptAndSetCredentialInformations()
		if dErr != nil {
			fmt.Println(dErr)
			return
		}
		fmt.Println(newCredential.GetId(), newCredential.GetServiceName(), newCredential.GetUsername(), newCredential.GetEmail(), newCredential.GetPassword())
	}
	return
}
