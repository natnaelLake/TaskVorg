package middleware

import (
	"backend/Types"
	"backend/Utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendGraphQLRequest(method, query string, variables map[string]interface{}) (*Types.GraphQLResponse, error) {
	requestBody, err := json.Marshal(Types.GraphQLRequest{
		Query:     query,
		Variables: variables,
	})
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	req, err := http.NewRequest(method, Utils.GraphqlEndpoint, bytes.NewBuffer(requestBody))
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

	var response Types.GraphQLResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if len(response.Errors) > 0 {
		return &response, fmt.Errorf("graphql errors: %v", response.Errors)
	}

	return &response, nil
}
