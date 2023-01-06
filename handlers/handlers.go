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

/*Handlers establece las rutas y el puerto de escucha */
func Handlers() {
	router := mux.NewRouter()

	//Endpoint de registro
	router.HandleFunc("/registro", middlew.CheckBD(routers.Registro)).Methods("POST")
	//Endpoint de login
	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")
	//Endpoint de verperfil
	router.HandleFunc("/verperfil", middlew.CheckBD(middlew.ValidateJWT(routers.VerPerfil))).Methods("GET")
	//Endpoint de modificarPerfil
	router.HandleFunc("/modificarPerfil", middlew.CheckBD(middlew.ValidateJWT(routers.ModificarPerfil))).Methods("PUT")
	//Endpoint de tweet
	router.HandleFunc("/tweet", middlew.CheckBD(middlew.ValidateJWT(routers.GuardarTweet))).Methods("POST")
	//Endpoint de leoTweets
	router.HandleFunc("/leoTweets", middlew.CheckBD(middlew.ValidateJWT(routers.LeerTweets))).Methods("GET")
	//Endpoint de eliminarTweet
	router.HandleFunc("/eliminarTweet", middlew.CheckBD(middlew.ValidateJWT(routers.EliminarTweet))).Methods("DELETE")

	//Endpoint de subirAvatar
	router.HandleFunc("/subirAvatar", middlew.CheckBD(middlew.ValidateJWT(routers.SubirAvatar))).Methods("POST")
	//Endpoint de obtenerAvatar
	router.HandleFunc("/obtenerAvatar", middlew.CheckBD(routers.ObtenerAvatar)).Methods("GET")
	//Endpoint de subirBanner
	router.HandleFunc("/subirBanner", middlew.CheckBD(middlew.ValidateJWT(routers.SubirBanner))).Methods("POST")
	//Endpoint de obtenerBanner
	router.HandleFunc("/obtenerBanner", middlew.CheckBD(routers.ObtenerBanner)).Methods("GET")

	//Endpoint de altaRelacion
	router.HandleFunc("/altaRelacion", middlew.CheckBD(middlew.ValidateJWT(routers.AltaRelacion))).Methods("POST")
	//Endpoint de bajaRelacion
	router.HandleFunc("/bajaRelacion", middlew.CheckBD(middlew.ValidateJWT(routers.BajaRelacion))).Methods("DELETE")
	//Endpoint de bajaRelacion
	router.HandleFunc("/consultaRelacion", middlew.CheckBD(middlew.ValidateJWT(routers.ConsultaRelacion))).Methods("GET")

	//Endpoint de listaUsuarios
	router.HandleFunc("/listaUsuarios", middlew.CheckBD(middlew.ValidateJWT(routers.ListaUsuarios))).Methods("GET")
	//Endpoint de leoTweetsSeguidores
	router.HandleFunc("/leoTweetsSeguidores", middlew.CheckBD(middlew.ValidateJWT(routers.LeerTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
