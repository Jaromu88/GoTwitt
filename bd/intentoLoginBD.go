package bd

import (
	"github.com/Jaromu88/GoTwitt/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLoginBD chequea el intento de login con los datos de la BD */
func IntentoLoginBD(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := CheckUserExist(email)

	if !encontrado {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usu, false
	}

	return usu, true

}
