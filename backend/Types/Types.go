package Types

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
