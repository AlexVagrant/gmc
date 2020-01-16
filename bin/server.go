package main

import (
	"net/http"

	"github.com/AlexVagrant/gmc"
)

func main() {
	mux := gmc.DefineApp()
  // 这里应该设置router方法
	mux.ServerMux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello"))
	})
	mux.Start(":8001")
}
