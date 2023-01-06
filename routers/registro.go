package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Jaromu88/GoTwitt/bd"
	"github.com/Jaromu88/GoTwitt/models"
)

/*Registro es la función para crear el registro de usuarios en la BD */
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	//Si el email está vacío
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	//Si la contraseña es demasiado corta
	if len(t.Password) < 6 {
		http.Error(w, "El Password debe tener al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.CheckUserExist(t.Email)

	if encontrado {
		http.Error(w, "Ya existe un usuario registrado con este email", 400)
		return
	}

	_, status, err := bd.InsertoRegistroBD(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado registrar al usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
