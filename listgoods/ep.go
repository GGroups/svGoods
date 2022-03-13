package listgoods

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type GoodBizInfo struct {
	SearchText string `json:"searchText"`
	PageSize   int    `json:"pageSize"`
	PageNum    int    `json:"pageNum"`
}

type GoodsListRequest struct {
	BizInfo GoodBizInfo `json:"bizInfo"`
	Type    string      `json:"type"`
}

type GoodsListResponse struct {
	Goods   []Good `json:"datas"`
	Msg     string `json:"msg"`
	RetCode string `json:"retode"`
}

func MakeCouponsEndPoint(sv IGoods) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(GoodsListRequest)
		if !ok {
			return GoodsListResponse{}, nil
		}
		if r.Type != "wx" {
			return nil, errors.New(INPUTE_RROR + `not "wx"`)
		}
		return GoodsListResponse{Goods: sv.ListGoods(r.Type), Msg: "ok", RetCode: "0"}, nil
	}
}
