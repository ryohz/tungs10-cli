package extra

import (
	"fmt"
	"os"
	"ryozk/tungs10-cli/constant"
)

func GetAppRootDir() (string, error) {
	home, getHomeErr := os.UserHomeDir()
	if getHomeErr != nil {
		return "", getHomeErr
	}
	var appRootDirPath string = fmt.Sprintf("%s/%s", home, constant.APP_ROOTDIR_NAME)
	return appRootDirPath, nil
}

func GetUserList() ([]string, error) {
	var userList []string
	appRootDir, getAppRootDirError := GetAppRootDir()
	if getAppRootDirError != nil {
		return userList, getAppRootDirError
	}
	userDirList, getUsersErr := os.ReadDir(appRootDir)
	if getUsersErr != nil {
		return userList, getUsersErr
	}
	for i := 0; i < len(userDirList); i++ {
		username := userDirList[i].Name()
		userList = append(userList, username)
	}
	return userList, nil
}

func IsUserExist() (bool, error) {
	appRootDir, getAppRootDirError := GetAppRootDir()
	if getAppRootDirError != nil {
		return false, getAppRootDirError
	}
	userDirList, getUsersErr := os.ReadDir(appRootDir)
	if getUsersErr != nil {
		return false, getUsersErr
	}
	if len(userDirList) > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
