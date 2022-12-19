package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Jaromu88/GoTwitt/bd"
	"github.com/Jaromu88/GoTwitt/jwt"
	"github.com/Jaromu88/GoTwitt/models"
)

/*Login realiza el login*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o contraseña incorrectas"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email no puede estar vacío", 400)
		return
	}

	//Compruebo si el usuario existe en la base de datos
	documento, existe := bd.IntentoLogin(t.Email, t.Password)

	if !existe {
		http.Error(w, "Usuario y/o contraseña incorrectas", 400)
		return
	}

	//Si existe genero el token
	jwtKey, err := jwt.GenerarToken(documento)

	if err != nil {
		http.Error(w, "Ocurrió un error al generar el token"+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//Por si quiero guardar el token en un cookie
	expirationTime := time.Now().Add(24 * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
