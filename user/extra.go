package user

import (
	"database/sql"
	"fmt"
	"os"
	constant "ryozk/tungs10-cli/constant"
	"ryozk/tungs10-cli/extra"
)

func (user *User) GetUserDir() (string, error) {
	appRootDir, getAppRootDirError := extra.GetAppRootDir()
	if getAppRootDirError != nil {
		return "", getAppRootDirError
	}
	var userDir string = fmt.Sprintf("%s/%s", appRootDir, user.GetUsername())
	return userDir, nil
}

func (user *User) IsCredentialExists() bool {
	userDir, gudErr := user.GetUserDir()
	if gudErr != nil {
		panic(gudErr)
	}
	var dbPath string = fmt.Sprintf("%s/%s", userDir, constant.CREDENTIAL_DATABASE_NAME)
	db, doErr := sql.Open("sqlite3", dbPath)
	if doErr != nil {
		return false
	}
	defer db.Close()
	statement, pErr := db.Prepare("select count(*) from credentials ")
	if pErr != nil {
		return false
	}
	var rowSum int
	statement.QueryRow().Scan(&rowSum)
	if rowSum > 0 {
		return true
	}
	return false
}

func GetUserList() ([]string, error) {
	appRootDir, gardErr := extra.GetAppRootDir()
	if gardErr != nil {
		return nil, gardErr
	}
	userDirs, rudErr := os.ReadDir(appRootDir)
	if rudErr != nil {
		return nil, rudErr
	}
	var userList []string
	for i := 0; i < len(userDirs); i++ {
		userList = append(userList, userDirs[i].Name())
	}
	return userList, nil
}
