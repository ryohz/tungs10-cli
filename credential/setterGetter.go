package credential

import (
	"ryozk/tungs10-cli/user"
)

func (credential *Credential) SetUser(user user.User) {
	credential.user = user
	return
}

func (credential *Credential) SetId(id int) {
	credential.id = id
	return
}

func (credential *Credential) GetId() int {
	return credential.id
}

func (credential *Credential) SetServiceName(serviceName string) {
	credential.serviceName = serviceName
	return
}

func (credential *Credential) GetServiceName() string {
	return credential.serviceName
}

func (credential *Credential) SetUsername(username string) {
	credential.username = username
	return
}

func (credential *Credential) GetUsername() string {
	return credential.username
}

func (credential *Credential) SetEmail(email string) {
	credential.email = email
	return
}

func (credential *Credential) GetEmail() string {
	return credential.email
}

func (credential *Credential) SetPassword(password string) {
	credential.password = password
	return
}

func (credential *Credential) GetPassword() string {
	return credential.password
}

func (credential *Credential) SetEncryptedServiceName(encryptedServiceName string) {
	credential.encryptedServiceName = encryptedServiceName
	return
}

func (credential *Credential) GetEncryptedServiceName() string {
	return credential.encryptedServiceName
}

func (credential *Credential) SetEncryptedUserName(encryptedUserName string) {
	credential.encryptedUserName = encryptedUserName
	return
}

func (credential *Credential) GetEncryptedUserName() string {
	return credential.encryptedUserName
}

func (credential *Credential) SetEncryptedEmail(encryptedEmail string) {
	credential.encryptedEmail = encryptedEmail
	return
}

func (credential *Credential) GetEncryptedEmail() string {
	return credential.encryptedEmail
}

func (credential *Credential) SetEncryptedPassword(encryptedPassword string) {
	credential.encryptedPassword = encryptedPassword
	return
}

func (credential *Credential) GetEncryptedPassword() string {
	return credential.encryptedPassword
}
