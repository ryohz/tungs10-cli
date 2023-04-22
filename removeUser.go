package main

import (
	"fmt"
	"ryozk/tungs10-cli/extra"
	"ryozk/tungs10-cli/user"

	"github.com/manifoldco/promptui"
)

func removeUser() {
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
	username, rpouErr := promptOfUsername.Run()
	if rpouErr != nil {
		fmt.Printf("Prompt failed %v\n", rpouErr)
		return
	}

	promptOfPassword := promptui.Prompt{
		Label: "Password",
		Mask:  '*',
	}
	password, rpopErr := promptOfPassword.Run()
	if rpopErr != nil {
		fmt.Printf("Prompt failed %v\n", rpopErr)
		return
	}

	var loginUser *user.User = user.New()
	loginUser.SetUsername(username)
	loginUser.SetPassword(password)
	var isLogined bool = loginUser.Login()
	if !isLogined {
		fmt.Println("failed to login")
		return
	}
	fmt.Println("Do you really want to remove this user?")
	promptOfVerify := promptui.Prompt{
		Label: "type 'remove' to remove this user",
	}
	removeVerify, rpovErr := promptOfVerify.Run()
	if rpovErr != nil {
		fmt.Printf("Prompt failed %v\n", rpovErr)
		return
	}

	if removeVerify == "remove" {
		var ruErr error = loginUser.RemoveUser()
		if ruErr != nil {
			fmt.Println(ruErr)
			return
		}
		fmt.Printf("%s is removed successfully\n", loginUser.GetUsername())
		return
	}
	fmt.Println("removing is canceled")
	return
}
