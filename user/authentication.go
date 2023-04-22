package user

import (
	"fmt"
	"os"
	constant "ryozk/tungs10-cli/constant"
)

func (user *User) Login() bool {
	userDir, err := user.GetUserDir()
	if err != nil {
		return false
	}
	var passwordFile string = fmt.Sprintf("%s/%s", userDir, constant.PASSWD_FILENAME)
	passwordByte, err := os.ReadFile(passwordFile)
	if err != nil {
		return false
	}
	user.hashAndSetPassword()
	if string(passwordByte) == string(user.GetHashedPassword()) {
		return true
	}
	return false
}
