package filter

import (
	"os"
	"net/http"
	"github.com/YCLiang95/CS160Group1OFS/backend/utils"
	"github.com/YCLiang95/CS160Group1OFS/backend/control/api"
)

type HttpServer struct {
	hs   *http.Server
	Done chan struct{}
}

func NewHttpServer(port string) *HttpServer {

	http.HandleFunc("/cs160/user/register", BrowserWapper(api.Register))
	http.HandleFunc("/cs160/user/login", BrowserWapper(api.Login))
	http.HandleFunc("/cs160/mechandise/list",BrowserWapper(api.List))
	http.HandleFunc("/cs160/mechandise/save",BrowserWapper(api.SaveMerchandis))
	http.HandleFunc("/cs160/mechandise/update",BrowserWapper(api.UpdateMerchandis))
	http.HandleFunc("/cs160/mechandise/get",BrowserWapper(api.GetMerchandise))

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
