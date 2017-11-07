package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

//Encrypt a string
func Encrypt(toEncrypt []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(toEncrypt, bcrypt.DefaultCost)
}

//CompareHases returns an error if the hashes arnt equal
func CompareHashes(toCompare1 []byte, toCompare2 []byte) error {
	return bcrypt.CompareHashAndPassword(toCompare1, toCompare2)
}
