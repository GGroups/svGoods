package listgoods

import (
	"context"
	"errors"

	ERR "github.com/GGroups/svGoods/comm_err"
	"github.com/go-kit/kit/endpoint"
)

type GoodBizInfo struct {
	SearchText string `json:"searchText"`
	CategoryId int    `json:"categoryId"`
	Cat2ndId   int    `json:"cat2ndId"`
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
			return nil, errors.New(ERR.INPUTE_RROR + `not "wx"`)
		}
		bi := r.BizInfo
		var listgood []Good

		if bi.CategoryId > 0 {
			listgood = sv.ListGoodsInCat2nd(bi.SearchText, bi.CategoryId, bi.Cat2ndId, bi.PageNum, bi.PageSize)
			return GoodsListResponse{Goods: listgood, Msg: "ok", RetCode: "0"}, nil
		}

		if len(bi.SearchText) == 0 { //查询所有,带推荐算法
			listgood = sv.ListGoodsRA(bi.PageNum, bi.PageSize)
		} else {
			listgood = sv.ListGoodsSearchName(bi.SearchText, bi.PageNum, bi.PageSize)
		}
		return GoodsListResponse{Goods: listgood, Msg: "ok", RetCode: "0"}, nil
	}
}
