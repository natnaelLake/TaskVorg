package main

import (
    "encoding/json"
    "net/http"
    "bytes"
    "fmt"
    "log"
    "github.com/gorilla/handlers"
)

const graphqlEndpoint = "http://localhost:8080/v1/graphql"

type GraphQLRequest struct {
    Query     string                 `json:"query"`
    Variables map[string]interface{} `json:"variables"`
}

type GraphQLResponse struct {
    Data   interface{} `json:"data"`
    Errors []struct {
        Message string `json:"message"`
    } `json:"errors"`
}

func sendGraphQLRequest(query string, variables map[string]interface{}) (*GraphQLResponse, error) {
    requestBody, err := json.Marshal(GraphQLRequest{
        Query:     query,
        Variables: variables,
    })
    if err != nil {
        return nil, fmt.Errorf("error marshaling request body: %w", err)
    }

    req, err := http.NewRequest("POST", graphqlEndpoint, bytes.NewBuffer(requestBody))
    if err != nil {
        return nil, fmt.Errorf("error creating request: %w", err)
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("x-hasura-admin-secret", "myadminsecretkey")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("error sending request: %w", err)
    }
    defer resp.Body.Close()

    var response GraphQLResponse
    err = json.NewDecoder(resp.Body).Decode(&response)
    if err != nil {
        return nil, fmt.Errorf("error decoding response: %w", err)
    }

    if len(response.Errors) > 0 {
        return &response, fmt.Errorf("graphql errors: %v", response.Errors)
    }

    return &response, nil
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
    var req GraphQLRequest

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    response, err := sendGraphQLRequest(req.Query, req.Variables)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(map[string]interface{}{
        "data": response.Data,
    }); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/graphql", tasksHandler)

    // Apply CORS middleware
    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Allow only this origin
        handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}), // Allow specific methods
        handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}), // Allow specific headers
    )(mux)

    log.Fatal(http.ListenAndServe(":8081", corsHandler)) // Ensure this port matches your frontend configuration
}