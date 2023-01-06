package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Jaromu88/GoTwitt/bd"
	"github.com/Jaromu88/GoTwitt/models"
)

/*GuardarTweet permite grabar el tweet en la BD */
func GuardarTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GuardarTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertarTweetBD(registro)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el tweet "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
