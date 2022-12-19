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
	router.HandleFunc("/verperfil", middlew.CheckBD(middlew.ValidarJWT(routers.VerPerfil))).Methods("GET")
	//Endpoint de modificarPerfil
	router.HandleFunc("/modificarPerfil", middlew.CheckBD(middlew.ValidarJWT(routers.ModificarPerfil))).Methods("PUT")
	//Endpoint de tweet
	router.HandleFunc("/tweet", middlew.CheckBD(middlew.ValidarJWT(routers.GraboTweet))).Methods("POST")
	//Endpoint de leoTweets
	router.HandleFunc("/leoTweets", middlew.CheckBD(middlew.ValidarJWT(routers.LeoTweets))).Methods("GET")
	//Endpoint de eliminarTweet
	router.HandleFunc("/eliminarTweet", middlew.CheckBD(middlew.ValidarJWT(routers.EliminarTweet))).Methods("DELETE")

	//Endpoint de subirAvatar
	router.HandleFunc("/subirAvatar", middlew.CheckBD(middlew.ValidarJWT(routers.SubirAvatar))).Methods("POST")
	//Endpoint de obtenerAvatar
	router.HandleFunc("/obtenerAvatar", middlew.CheckBD(routers.ObtenerAvatar)).Methods("GET")

	//Endpoint de subirBanner
	router.HandleFunc("/subirBanner", middlew.CheckBD(middlew.ValidarJWT(routers.SubirBanner))).Methods("POST")
	//Endpoint de obtenerBanner
	router.HandleFunc("/obtenerBanner", middlew.CheckBD(routers.ObtenerBanner)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
