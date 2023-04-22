package main

import (
	"errors"
	"fmt"
	"ryozk/tungs10-cli/credential"
	"ryozk/tungs10-cli/extra"
	"ryozk/tungs10-cli/user"
	"strconv"

	"github.com/manifoldco/promptui"
)

func removeCredential() {
	userList, gulErr := extra.GetUserList()
	if gulErr != nil {
		fmt.Println(gulErr)
		return
	}
	if len(userList) == 0 {
		fmt.Println("you don't have any user")
		return
	}
	promptOfUsername := promptui.Prompt{
		Label: "Username",
	}
	promptOfPassword := promptui.Prompt{
		Label: "Password",
		Mask:  '*',
	}
	username, rpouErr := promptOfUsername.Run()
	if rpouErr != nil {
		fmt.Printf("Prompt failed: %v\n", rpouErr)
		return
	}
	password, rpopErr := promptOfPassword.Run()
	if rpopErr != nil {
		fmt.Printf("Prompt failed: %v\n", rpopErr)
		return
	}
	var loginUser *user.User = user.New()
	loginUser.SetUsername(username)
	loginUser.SetPassword(password)
	var allowedLogin bool = loginUser.Login()
	if !allowedLogin {
		fmt.Println("failed to login")
		return
	}
	var isCredentialExists bool = loginUser.IsCredentialExists()
	if !isCredentialExists {
		fmt.Println("any credential does not exists")
		return
	}
	var rasiErr error = loginUser.ReadAndSetIv()
	if rasiErr != nil {
		fmt.Println(rasiErr)
		return
	}
	loginUser.CalcAndSetKey()
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}
	promptOfId := promptui.Prompt{
		Label:    "Id",
		Validate: validate,
	}
	idStr, rpoiErr := promptOfId.Run()
	if rpoiErr != nil {
		fmt.Printf("Prompt failed: %v\n", rpoiErr)
		return
	}
	var targetCredential *credential.Credential = credential.New()
	targetCredential.SetUser(*loginUser)
	id, is2iErr := strconv.Atoi(idStr)
	if is2iErr != nil {
		fmt.Println(is2iErr)
		return
	}
	targetCredential.SetId(id)
	var raseciErr error = targetCredential.ReadAndSetEncryptedCredentialInformations()
	if raseciErr != nil {
		if raseciErr.Error() == "sql: no rows in result set" {
			fmt.Printf("no credential has id %d is found\n", id)
			return
		}
		fmt.Println(raseciErr)
		return
	}
	var dErr error = targetCredential.DecryptAndSetCredentialInformations()
	if dErr != nil {
		fmt.Println(dErr)
		return
	}
	fmt.Println("information of target credential")
	fmt.Printf("[id] %d\n", targetCredential.GetId())
	fmt.Printf("[service name]: %s\n", targetCredential.GetServiceName())
	fmt.Printf("[email] %s\n", targetCredential.GetEmail())
	fmt.Printf("[username] %s\n", targetCredential.GetUsername())
	fmt.Printf("[password] %s\n", targetCredential.GetPassword())
	fmt.Println("Do you really want to remove this credential?")
	promptOfVerify := promptui.Prompt{
		Label: "type 'remove' to remove this user",
	}
	removeVerify, rpovErr := promptOfVerify.Run()
	if rpovErr != nil {
		fmt.Printf("Prompt failed %v\n", rpovErr)
		return
	}
	if removeVerify != "remove" {
		fmt.Println("removing the credential is cancelled")
		return
	}
	var rErr error = targetCredential.RemoveFromDB()
	if rErr != nil {
		fmt.Println(rErr)
		return
	}
	fmt.Println("successfully removed the credential")
	return
}
