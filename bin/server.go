package main

import (
	"net/http"

	"github.com/AlexVagrant/gmc"
)

func main() {
	mux := gmc.DefineApp()
	mux.Mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello"))
	})
	mux.Start(":8001")
}
