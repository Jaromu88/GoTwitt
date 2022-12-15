package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Jaromu88/GoTwitt/middlew"
	"github.com/Jaromu88/GoTwitt/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores establece las rutas y el puerto de escucha */
func Manejadores() {
	router := mux.NewRouter()

	//Endpoint de registro
	router.HandleFunc("/registro", middlew.CheckBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
