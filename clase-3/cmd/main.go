package main

import (
	"curso_apis_go/clase-3/handler"
	"curso_apis_go/clase-3/storage"
	"log"
	"net/http"
)

func main() {
	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)

	log.Println("Servidor iniciado en el puerto 8888")
	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}
