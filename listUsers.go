package main

import (
	"fmt"
	"ryozk/tungs10-cli/user"
)

func listUsers() {
	userList, gulErr := user.GetUserList()
	if gulErr != nil {
		fmt.Println(gulErr)
		return
	}
	if len(userList) == 0 {
		fmt.Println("you don't have any users")
		return
	}
	for i := 0; i < len(userList); i++ {
		fmt.Printf("・ %s", userList[i])
	}
	return
}
