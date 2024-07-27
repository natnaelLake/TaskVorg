package main

import (
	"backend/Controllers"
	"backend/middleware"
	"fmt"
	"log"
	"net/http"
)

func main() {
    mux := http.NewServeMux();
    mux.HandleFunc("/graphql", Controllers.TasksHandler);
	mux.HandleFunc("/uploads", Controllers.UploadHandler);
	fileServer := http.FileServer(http.Dir("./uploads"))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", fileServer))
	log.Fatal(http.ListenAndServe(":8081", middleware.CORS(mux)));
    fmt.Println("Server is running on localhost:8081");
}