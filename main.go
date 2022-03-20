package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	ND "github.com/GGroups/svGoods/category2nd"
	LGDS "github.com/GGroups/svGoods/listgoods"
	log "github.com/cihub/seelog"

	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	log.Info("#good service started!")

	nd := ND.Cat2nd{}
	epnd := ND.MakeWCat2ndEndPoint(nd) //cat2nd写服务，重新设置二级分类，全量重写。
	sgdn := LGDS.Good{}
	epgd := LGDS.MakeCouponsEndPoint(sgdn) //index页面列出所有商品

	mysvr := httpTransport.NewServer(epnd, ND.WCat2ndDecodeRequest, ND.CommEncodeResponse)
	mysgdsvr := httpTransport.NewServer(epgd, LGDS.GoodsListDecodeRequest, LGDS.CommEncodeResponse)

	routeSvr := mux.NewRouter()

	routeSvr.Handle(`/gpwm/goods/setCat2nd`, mysvr).Methods("POST")
	routeSvr.Handle(`/gpwm/goods/getGoodList`, mysgdsvr).Methods("POST")

	//main loop
	ch := make(chan error, 2)
	go func() {
		log.Info("0.0.0.0:8007", `/gpwm/goods/**`)
		ch <- http.ListenAndServe("0.0.0.0:8007", routeSvr)
	}()
	go func() {
		log.Info("##", "wait for exit sigint...")
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		ch <- fmt.Errorf("%s", <-c)
	}()

	log.Info("MainSvr Terminated", <-ch)
}
