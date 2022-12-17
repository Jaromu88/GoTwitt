package routers

import (
	"errors"
	"strings"

	"github.com/Jaromu88/GoTwitt/bd"
	"github.com/Jaromu88/GoTwitt/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email valor de email que voy a usar en todos mis endpoints */
var Email string

/*IDUsuario id devuelto del modelo, usado en todos mis endpoints */
var IDUsuario string

/*ProcesarToken procesa el token para extraer sus valores */
func ProcesarToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("EstaEsMiClave")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.CheckUsuarioYaExiste(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err

}
