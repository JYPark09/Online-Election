package main

import "crypto/sha512"

type User struct {
	id int
	pw string
}

func getEncryptedPassword(pw string) string {
	hash := sha512.Sum512([]byte(pw))

	return string(hash[:])
}

func checkUserPassword(id int, pw string) bool {
	usr := getUserInfo(id)
	if usr == nil {
		return false
	}

	return getEncryptedPassword(pw) == usr.pw
}
