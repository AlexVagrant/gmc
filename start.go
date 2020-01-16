package gmc

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Mux struct {
	ServerMux *http.ServeMux
	app *http.Server
}

func (m Mux) quit() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		// 接受退出通知
		if err := m.app.Close(); err != nil {
			log.Fatal("Close server:", err)
		}
	}()

}

func DefineApp() Mux {
	var mux Mux
	mux.ServerMux = http.NewServeMux()
	mux.quit()
	return mux
}

func (m *Mux) Start(addr string) {

	// 配置服务基础信息 后面通过config直接导入进来
	m.app = &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 3,
		Handler:      m.ServerMux,
	}

	if err := m.app.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
