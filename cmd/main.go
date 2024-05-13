package main

import (
	"fmt"

	"github.com/badsector998/go-rbac/internal/app"
	"github.com/badsector998/go-rbac/internal/http"
)

func main() {
	fmt.Println("golang rbac")

	s := http.NewServer(app.NewApp())

	s.Run()
}
