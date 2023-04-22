package credential

import (
	"ryozk/tungs10-cli/crypto"
	"ryozk/tungs10-cli/user"
)

func (credential *Credential) EncryptAndSetCredentialInformations() error {
	encryptedServiceName, err := encrypt(credential.GetServiceName(), credential.user)
	if err != nil {
		return err
	}
	encryptedUserName, err := encrypt(credential.GetUsername(), credential.user)
	if err != nil {
		return err
	}
	encryptedEmail, err := encrypt(credential.GetEmail(), credential.user)
	if err != nil {
		return err
	}
	encryptedPassword, err := encrypt(credential.GetPassword(), credential.user)
	if err != nil {
		return err
	}
	credential.SetEncryptedServiceName(encryptedServiceName)
	credential.SetEncryptedUserName(encryptedUserName)
	credential.SetEncryptedEmail(encryptedEmail)
	credential.SetEncryptedPassword(encryptedPassword)
	return nil
}

func (credential *Credential) DecryptAndSetCredentialInformations() error {
	decryptedServiceName, dsErr := decrypt(credential.GetEncryptedServiceName(), credential.user)
	if dsErr != nil {
		return dsErr
	}
	decryptedUserName, duErr := decrypt(credential.GetEncryptedUserName(), credential.user)
	if duErr != nil {
		return duErr
	}
	decryptedEmail, deErr := decrypt(credential.GetEncryptedEmail(), credential.user)
	if deErr != nil {
		return deErr
	}
	decryptedPassword, dpErr := decrypt(credential.GetEncryptedPassword(), credential.user)
	if dpErr != nil {
		return dpErr
	}
	credential.SetServiceName(string(decryptedServiceName))
	credential.SetUsername(string(decryptedUserName))
	credential.SetEmail(string(decryptedEmail))
	credential.SetPassword(string(decryptedPassword))
	return nil
}

func encrypt(data string, user user.User) (string, error) {
	encrypted, err := crypto.Encrypt([]byte(data), []byte(user.GetKey()), user.GetIv())
	if err != nil {
		return "", err
	}
	return string(encrypted), nil
}

func decrypt(data string, user user.User) (string, error) {
	var iv []byte = user.GetIv()
	var key []byte = user.GetKey()
	decrypted, dErr := crypto.Decrypt([]byte(data), key, iv)
	if dErr != nil {
		return "", dErr
	}
	return string(decrypted), nil
}
