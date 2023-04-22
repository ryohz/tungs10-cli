package user

func (user *User) GetIv() []byte {
	return user.iv
}
func (user *User) SetIv(iv []byte) []byte {
	user.iv = iv
	return user.iv
}
func (user *User) GetPassword() string {
	return user.password
}

func (user *User) SetPassword(password string) string {
	user.password = password
	return user.password
}

func (user *User) SetUsername(username string) string {
	user.username = username
	return user.username
}

func (user *User) GetUsername() string {
	return user.username
}

func (user *User) SetHashedPassword(hashedPassword []byte) {
	user.hashedPassword = hashedPassword
	return
}

func (user *User) GetHashedPassword() []byte {
	return user.hashedPassword
}

func (user *User) SetKey(key []byte) {
	user.key = key
	return
}

func (user *User) GetKey() []byte {
	return user.key
}
