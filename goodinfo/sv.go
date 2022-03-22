package goodinfo

import (
	"encoding/json"
	"strconv"

	CMM "github.com/GGroups/svGoods/comm"
	log "github.com/cihub/seelog"

	redis "github.com/gomodule/redigo/redis"
)

const (
	GOODINFO_REDIS_KEY = "gpwm_goodinfos"
	REDIS_URL          = "redis://127.0.0.1:5101"
)

type IGoodInf interface {
	GetGoodInf(goodId int, ret_c2nds *GoodInf) error
	SetGoodInfItems(newlist []GoodInf) error
}

type PicInf struct {
	Id  int    `json:"id"`
	Pic string `json:"pic"`
}

type GoodInf struct {
	GoodID     int      `json:"goodID"`
	GoodName   string   `json:"goodName"`
	GoodDesc   string   `json:"goodDesc"`
	MinPrice   float32  `json:"minPrice"`
	SizeList   []string `json:"sizeList"`
	ColorList  []string `json:"colorList"`
	PicInfList []PicInf `json:"picInfList"`
}

func (s GoodInf) GetGoodInf(goodId int, goodinfo *GoodInf) error {
	c2, err := redis.DialURL(REDIS_URL)
	if err != nil {
		return err
	}
	defer c2.Close()
	redis_key := GOODINFO_REDIS_KEY + strconv.Itoa(goodId)
	val, _ := c2.Do("get", redis_key)
	err = json.Unmarshal(val.([]byte), goodinfo)
	if err != nil {
		log.Error("##format error", err.Error())
	}
	for id, obj := range goodinfo.PicInfList {
		goodinfo.PicInfList[id].Pic = obj.Pic + CMM.GPWM_AUTH
	}

	return nil
}

func (s GoodInf) SetGoodInfItems(newlist []GoodInf) error {
	c2, err := redis.DialURL("redis://127.0.0.1:5101")
	if err != nil {
		return err
	}
	defer c2.Close()

	for _, cat := range newlist {
		strval, e1 := json.Marshal(cat)
		if e1 == nil {
			redis_key := GOODINFO_REDIS_KEY + strconv.Itoa(cat.GoodID)
			c2.Do("set", redis_key, strval)
		} else {
			log.Error("##format error", e1.Error())
		}
	}

	return nil
}
