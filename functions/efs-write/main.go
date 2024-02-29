package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

const EFS_PATH = "/mnt/msg"

func HandleRequest(ctx context.Context) (string, error) {
	fileName := fmt.Sprintf("%s.txt", time.Now().String())
	content := "My name is Leedonggyu"

	err := os.WriteFile(fmt.Sprintf("%s/%s", EFS_PATH, fileName), []byte(content), 0644)
	if err != nil {
		return "", err
	}

	return "Success", nil
}

func isConnectEFS(path string) {
	if _, err := os.Stat(EFS_PATH); os.IsNotExist(err) {
		log.Fatalln(err)
	}
}

func main() {
	isConnectEFS(EFS_PATH)
	lambda.Start(HandleRequest)
}
