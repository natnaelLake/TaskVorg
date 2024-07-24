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
    log.Fatal(http.ListenAndServe(":8081", middleware.CORS(mux)));
    fmt.Println("Server is running on localhost:8081");
}