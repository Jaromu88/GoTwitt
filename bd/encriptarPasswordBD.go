package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword es la función que me encriptará la password*/
func EncriptarPassword(pass string) (string, error) {
	coste := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), coste)
	return string(bytes), err
}
