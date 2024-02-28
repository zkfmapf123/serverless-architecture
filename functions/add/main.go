package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Response struct {
	Result int `json:"result"`
}

func HandleRequest(ctx context.Context, req Request) (Response, error) {
	resp := Response{
		Result: req.A + req.B,
	}

	return resp, nil
}

func main() {
	lambda.Start(HandleRequest)
}
