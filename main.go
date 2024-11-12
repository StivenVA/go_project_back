package main

import (
	"fmt"
	"log"
	"net/http"
	"proyecto_go/controllers"
)

func main() {

	mux := http.NewServeMux()

	authEndPoints := controllers.GetEndPoints()

	for _, endPoint := range authEndPoints {
		path, handler := endPoint()
		fmt.Println(path)
		mux.HandleFunc(path, handler)
	}

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Servidor ejecutándose en http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(mux)))
}
