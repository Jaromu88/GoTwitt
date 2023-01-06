package main

import (
	"log"

	"github.com/Jaromu88/GoTwitt/bd"
	"github.com/Jaromu88/GoTwitt/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Error al conectar con la BD")
		return
	}
	handlers.Handlers()
}
