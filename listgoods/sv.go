package listgoods

import (
	"strconv"

	GD "github.com/GGroups/svGoods/mgoods"
	log "github.com/cihub/seelog"
	"github.com/jmoiron/sqlx"
)

type IGoods interface {
	ListGoodsRA(pageNum int, pageSize int) []Good
	ListGoodsSearchName(searchtxt string, pageNum int, pageSize int) []Good
	ListGoodsInCat2nd(searchtxt string, category int, cat2nd int, pageNum int, pageSize int) []Good
}

type Good struct {
	GoodId   int     `json:"goodId"`
	GoodName string  `json:"goodName"`
	MinPrice float32 `json:"minPrice"`
	MaxPrice float32 `json:"maxPrice"`
	Pic      string  `json:"pic"`
}

func (s Good) ListGoodsRA(pageNum int, pageSize int) []Good {
	db, err := sqlx.Open(GD.LITE3, GD.DB_FILE)
	if err != nil {
		log.Error("##", err.Error())
		return []Good{}
	}
	rows := []GD.MGood{}
	sqlstr := GD.SQL_SEL_MGOOD + GD.SQL_SEL_MGOOD_CND1 +
		GD.SQL_SEL_MGOOD_LIMIT + strconv.Itoa(pageSize) + GD.SQL_SEL_MGOOD_OFFSET + strconv.Itoa(pageSize*pageNum)
	err = db.Select(&rows, sqlstr)
	if err != nil {
		log.Error("##", err.Error())
		return []Good{}
	}
	log.Info("->", sqlstr)
	epgood := make([]Good, len(rows))
	for c, g := range rows {
		epgood[c].GoodId = g.GoodId
		epgood[c].GoodName = g.GoodName
		epgood[c].MaxPrice = g.MaxPrice
		epgood[c].MinPrice = g.MinPrice
		//epgood[c].Pic
	}
	db.Close()
	return epgood
}

func (s Good) ListGoodsSearchName(searchtxt string, pageNum int, pageSize int) []Good {
	db, err := sqlx.Open(GD.LITE3, GD.DB_FILE)
	if err != nil {
		log.Error("##", err.Error())
		return []Good{}
	}
	rows := []GD.MGood{}
	sqlstr := GD.SQL_SEL_MGOOD + GD.SQL_SEL_MGOOD_CND1 + `and goodName="` + searchtxt + `"` +
		GD.SQL_SEL_MGOOD_LIMIT + strconv.Itoa(pageSize) + GD.SQL_SEL_MGOOD_OFFSET + strconv.Itoa(pageSize*pageNum)
	err = db.Select(&rows, sqlstr)
	if err != nil {
		log.Error("##", err.Error())
		return []Good{}
	}
	log.Info("->", sqlstr)
	epgood := make([]Good, len(rows))
	for c, g := range rows {
		epgood[c].GoodId = g.GoodId
		epgood[c].GoodName = g.GoodName
		epgood[c].MaxPrice = g.MaxPrice
		epgood[c].MinPrice = g.MinPrice
		//epgood[c].Pic
	}
	db.Close()
	return epgood
}

func (s Good) ListGoodsInCat2nd(searchtxt string, category int, cat2nd int, pageNum int, pageSize int) []Good {
	db, err := sqlx.Open(GD.LITE3, GD.DB_FILE)
	if err != nil {
		log.Error("##", err.Error())
		return []Good{}
	}
	rows := []GD.MGood{}
	var stxt string
	if len(searchtxt) > 0 {
		stxt = `and goodName="` + searchtxt + `" `
	}
	sqlstr := GD.SQL_SEL_MGOOD + GD.SQL_SEL_MGOOD_CND1 + stxt + `and category=` + strconv.Itoa(category) + ` and cat2nd=` + strconv.Itoa(cat2nd) +
		GD.SQL_SEL_MGOOD_LIMIT + strconv.Itoa(pageSize) + GD.SQL_SEL_MGOOD_OFFSET + strconv.Itoa(pageSize*pageNum)
	err = db.Select(&rows, sqlstr)
	if err != nil {
		log.Error("##", err.Error())
		return []Good{}
	}
	log.Info("ListGoodsInCat2nd->", len(rows), " ", sqlstr)
	epgood := make([]Good, len(rows))
	for c, g := range rows {
		epgood[c].GoodId = g.GoodId
		epgood[c].GoodName = g.GoodName
		epgood[c].MaxPrice = g.MaxPrice
		epgood[c].MinPrice = g.MinPrice
		//epgood[c].Pic
	}
	db.Close()
	return epgood
}
