package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	ND "github.com/GGroups/svGoods/category2nd"
	GDI "github.com/GGroups/svGoods/goodinfo"
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
	epgd := LGDS.MakeGoodlistEndPoint(sgdn) //index页面列出所有商品
	gdi := GDI.GoodInf{}
	epgdi := GDI.MakeGoodInfEndPoint(gdi)   //获取单个商品详细信息
	epwgdi := GDI.MakeWGoodInfEndPoint(gdi) //批量写入商品详细信息

	mysvr := httpTransport.NewServer(epnd, ND.WCat2ndDecodeRequest, ND.CommEncodeResponse)
	mysgdsvr := httpTransport.NewServer(epgd, LGDS.GoodsListDecodeRequest, LGDS.CommEncodeResponse)
	mygdisvr := httpTransport.NewServer(epgdi, GDI.GoodInfDecodeRequest, GDI.CommEncodeResponse)
	mywgdisvr := httpTransport.NewServer(epwgdi, GDI.WGoodInfDecodeRequest, GDI.CommEncodeResponse)

	routeSvr := mux.NewRouter()

	routeSvr.Handle(`/gpwm/goods/setCat2nd`, mysvr).Methods("POST")
	routeSvr.Handle(`/gpwm/goods/getGoodList`, mysgdsvr).Methods("POST")
	routeSvr.Handle(`/gpwm/goods/getGoodInfo`, mygdisvr).Methods("POST")
	routeSvr.Handle(`/gpwm/goods/setGoodInfoList`, mywgdisvr).Methods("POST")

	//main loop
	ch := make(chan error, 2)
	go func() {
		log.Info("0.0.0.0:8007", `/gpwm/goods/**`)
		ch <- http.ListenAndServeTLS("0.0.0.0:8007", "../gencert/tenfor.top_bundle.pem", "../gencert/tenfor.top.key", routeSvr)
	}()
	go func() {
		log.Info("##", "wait for exit sigint...")
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		ch <- fmt.Errorf("%s", <-c)
	}()

	log.Info("MainSvr Terminated", <-ch)
}
