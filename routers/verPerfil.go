package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Jaromu88/GoTwitt/bd"
)

/*VerPerfil me muestra un perfil de usuario con todos sus valores*/
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscarPerfil(ID)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar el registro"+err.Error(), 400)
		return
	}

	w.Header().Set("Context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)

}
