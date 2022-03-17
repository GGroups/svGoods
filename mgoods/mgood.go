package mgoods

import (
	"errors"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/logoove/sqlite"
)

type MGood struct {
	GoodId   int     `json:"goodId"`
	GoodKey  string  `json:"goodKey"`
	GoodName string  `json:"goodName"`
	Category string  `json:"category"`
	Cat2nd   string  `json:"cat2nd"`
	MinPrice float32 `json:"minPrice"`
	MaxPrice float32 `json:"maxPrice"`
	OnShelf  bool    `json:"onShelf"`
}

const (
	SQL_CRE = `CREATE TABLE MGood ("goodId" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	  "goodKey" char(200) NOT NULL, 
	  "goodName" char(200) NOT NULL, 
	  "category" char(100) NOT NULL, 
	  "cat2nd" char(100) NOT NULL, 
	  "minPrice" FLOAT ,
	  "maxPrice" FLOAT ,
	  "onShelf" BOOLEAN );`

	SQL_INS = `insert into MGood (
	"goodId",	"goodKey", 	"goodName", "category", "cat2nd", "minPrice","maxPrice","onShelf") 
	 VALUES (?,?,?,?,?,?,?,?)`

	DB_FILE = `./lite.db`

	LITE3 = "sqlite3"
)

type IMGood interface {
	FindGoodById(gnd *MGood, goodid int) error
	GetGoodPix(goodid int) []string
	//input or update goods info
	InGood(MGood) error
	UpdateGood(MGood) error
	UpdateGoodPix(goodid int, picList []string) error
	//init func
	InitTable() error
	InitLoadAllForUpdate() error
}

func (r MGood) InitTable() error {
	db, err := sqlx.Open(LITE3, DB_FILE)
	if err != nil {
		return err
	}
	_, err2 := db.Exec(SQL_CRE)
	if err2 != nil {
		return err2
	}
	db.Close()

	return nil
}

func (r MGood) InitLoadAllForUpdate() error {
	db, err := sqlx.Open(LITE3, DB_FILE)
	if err != nil {
		return err
	}
	rows := []MGood{}
	db.Select(&rows, "SELECT goodId, goodKey FROM MGood ORDER BY goodId")

	db.Close()
	return nil
}

func (r MGood) InGood(ngd MGood) error {
	db, err := sqlx.Open(LITE3, DB_FILE)
	if err != nil {
		return err
	}
	_, err2 := db.Exec(SQL_INS, ngd.GoodId, ngd.GoodKey, ngd.GoodName, ngd.Category, ngd.Cat2nd, ngd.MinPrice, ngd.MaxPrice, ngd.OnShelf)
	if err2 != nil {
		return err2
	}
	return nil
}

func (r MGood) UpdateGood(MGood) error {
	return nil
}

func (r MGood) FindGoodById(gnd *MGood, goodid int) error {
	db, err := sqlx.Open(LITE3, DB_FILE)
	if err != nil {
		return err
	}
	rows := []MGood{}
	err2 := db.Select(&rows, "SELECT * FROM MGood where goodId="+strconv.Itoa(goodid)+" ORDER BY id")
	if err2 != nil {
		return err2
	}
	if len(rows) == 1 {
		*gnd = rows[0]
		return nil
	} else if len(rows) == 1 {
		return errors.New("未找到商品id=" + strconv.Itoa(goodid))
	} else if len(rows) > 1 {
		return errors.New("找到商品" + strconv.Itoa(len(rows)) + ",id=" + strconv.Itoa(goodid))
	}
	return nil
}
