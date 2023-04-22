package user

import (
	"bytes"
	"crypto/sha512"
	"errors"
	"fmt"
	constant "ryozk/tungs10-cli/constant"
	"ryozk/tungs10-cli/extra"

	"os"
)

type User struct {
	username       string
	password       string
	key            []byte
	iv             []byte
	hashedPassword []byte
}

func New() *User {
	return &User{}
}

func (user *User) CreateUser() error {
	homeDir, ghdErr := os.UserHomeDir()
	if ghdErr != nil {
		return ghdErr
	}
	userlist, getUserlistErr := extra.GetUserList()
	if getUserlistErr != nil {
		return getUserlistErr
	}
	var sameNameErr error = errors.New("User that has this username is already existing")
	for i := 0; i < len(userlist); i++ {
		if user.username == userlist[i] {
			return sameNameErr
		}
	}
	var userDirPath string = fmt.Sprintf("%s/%s/%s", homeDir, constant.APP_ROOTDIR_NAME, user.username)
	var createUserDirError error = os.Mkdir(userDirPath, 0700)
	if createUserDirError != nil {
		return createUserDirError
	}
	user.hashAndSetPassword()
	spErr := user.storeHashedPassword()
	if spErr != nil {
		return spErr
	}
	var siErr error = user.StoreIv()
	if siErr != nil {
		return siErr
	}
	return nil
}

func (user *User) CalcAndSetKey() {
	if len(user.GetPassword()) > 16 {
		user.SetKey([]byte(user.GetPassword()[:16]))
	}
	var paddedPassword []byte = padPassword([]byte(user.GetPassword()))
	user.SetKey(paddedPassword)
	return
}

func (user *User) ReadAndSetIv() error {
	userDir, gudErr := user.GetUserDir()
	if gudErr != nil {
		return gudErr
	}
	var ivFile string = fmt.Sprintf("%s/%s", userDir, constant.IV_FILENAME)
	iv, rifErr := os.ReadFile(ivFile)
	if rifErr != nil {
		return rifErr
	}
	user.SetIv(iv)
	return nil
}

func (user *User) hashAndSetPassword() {
	var hasedPassword [64]byte = sha512.Sum512([]byte(user.GetPassword()))
	user.SetHashedPassword(hasedPassword[:])
	return
}

func (user *User) storeHashedPassword() error {
	userDir, getUserDirError := user.GetUserDir()
	if getUserDirError != nil {
		return getUserDirError
	}
	var passwordFile string = fmt.Sprintf("%s/%s", userDir, constant.PASSWD_FILENAME)
	os.WriteFile(passwordFile, []byte(user.GetHashedPassword()), 0700)
	return nil
}

func (user *User) StoreIv() error {
	var iv []byte = user.iv
	userDir, err := user.GetUserDir()
	if err != nil {
		return err
	}
	var ivFile string = fmt.Sprintf("%s/%s", userDir, constant.IV_FILENAME)
	var wifErr error = os.WriteFile(ivFile, iv, 0700)
	if wifErr != nil {
		return wifErr
	}
	return nil
}

func (user *User) RemoveUser() error {
	userDir, gudErr := user.GetUserDir()
	if gudErr != nil {
		return gudErr
	}
	var rErr error = os.RemoveAll(userDir)
	if rErr != nil {
		return rErr
	}
	return nil
}

func padPassword(data []byte) []byte {
	var length int = constant.KEY_LENGTH - (len(data) % constant.KEY_LENGTH)
	var pad []byte = bytes.Repeat([]byte{byte(length)}, length)
	return append(data, pad...)
}
