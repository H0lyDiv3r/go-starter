package main

import (
	"fmt"

	"github.com/H0lyDiv3r/go-headless-cms/internal/initializer"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectDB()
}

func main() {
	fmt.Println("working")
}
