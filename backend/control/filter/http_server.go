package filter

import (
	"os"
	"net/http"
	"github.com/cs160/project/utils"
	"github.com/cs160/project/control/api"
)

type HttpServer struct {
	hs   *http.Server
	Done chan struct{}
}

func NewHttpServer(port string) *HttpServer {
	http.HandleFunc("/cs160/user/register", BrowserWapper(api.Register))
	http.HandleFunc("/cs160/user/login", BrowserWapper(api.Login))
	server := new(http.Server)
	server.Addr = ":" + port
	server.Handler = nil

	return &HttpServer{
		hs:   server,
		Done: make(chan struct{}),
	}
}

func (hs *HttpServer) StartServer() {
	utils.Logger.Warning("Start API Service:", hs.hs.Addr)
	err := hs.hs.ListenAndServe()

	if err != nil {
		utils.Logger.Error("Fail to set up HttpServer: ", err)
		os.Exit(1)
	}
}
