package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	A int `json:"a"`
	B int `json:"b"`
}

func HandleRequest(ctx context.Context, e *MyEvent) (*int, error) {
	if e == nil {
		return nil, fmt.Errorf("received nil event %v", e)
	}

	res := add(e.A, e.B)
	return &res, nil
}

func add(a, b int) int {
	return a + b
}

func main() {
	lambda.Start(HandleRequest)
}
