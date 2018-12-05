package control

import (
	"os"
	"net/http"
	"github.com/YCLiang95/CS160Group1OFS/backend/utils"
	"github.com/YCLiang95/CS160Group1OFS/backend/control/api"
	"github.com/YCLiang95/CS160Group1OFS/backend/control/filter"
)

type HttpServer struct {
	hs   *http.Server
	Done chan struct{}
}

func NewHttpServer(port string) *HttpServer {

	http.HandleFunc("/cs160/user/register", filter.BrowserWapper(api.Register))
	http.HandleFunc("/cs160/user/login", filter.BrowserWapper(api.Login))
	http.HandleFunc("/cs160/mechandise/list",filter.BrowserWapper(api.List))
	http.HandleFunc("/cs160/mechandise/save",filter.BrowserWapper(api.SaveMerchandis))
	http.HandleFunc("/cs160/mechandise/update",filter.BrowserWapper(api.UpdateMerchandis))
	http.HandleFunc("/cs160/mechandise/get",filter.BrowserWapper(api.GetMerchandise))
    http.HandleFunc("/cs160/user/order/placeOrder",filter.BrowserWapper(api.PlaceOrder))
	http.HandleFunc("/cs160/user/order/getOrders",filter.BrowserWapper(api.GetOrders))
	http.HandleFunc("/cs160/user/order/getDelivery",filter.BrowserWapper(api.HandlerDelivery))

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
