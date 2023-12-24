package main

import (
	"fmt"
	"net/http"
	"notesServer/controllers/stdhttp"
	"notesServer/gate/psg"
)

func main() {
	s, err := psg.NewPsg("postgres://postgres:Russija13@localhost:5432/postgres", "postgres", "Russija13")
	if err != nil {
		fmt.Println("Error")
	}
	stdhttp.NewController("localhost:8080", s)
	http.ListenAndServe(":8080", nil)
}
