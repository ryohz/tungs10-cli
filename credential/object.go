package credential

import (
	"ryozk/tungs10-cli/user"

	_ "github.com/mattn/go-sqlite3"
)

type Credential struct {
	id                   int
	serviceName          string
	username             string
	email                string
	password             string
	encryptedServiceName string
	encryptedUserName    string
	encryptedEmail       string
	encryptedPassword    string
	user                 user.User
}

func New() *Credential {
	return &Credential{}
}
