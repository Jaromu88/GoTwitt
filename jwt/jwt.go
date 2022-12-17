package jwt

import (
	"time"

	"github.com/Jaromu88/GoTwitt/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*GenerarToken me va a generar el token con jwt */
func GenerarToken(t models.Usuario) (string, error) {
	miClave := []byte("EstaEsMiClave")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
