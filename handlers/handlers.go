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
	//Endpoint de login
	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")
	//Endpoint de verperfil
	//router.HandleFunc("/verperfil", middlew.CheckBD(middlew.ValidarJWT(routers.VerPerfil))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
