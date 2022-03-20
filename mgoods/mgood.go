package mgoods

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/logoove/sqlite"
)

type MGood struct {
	GoodId   int     `json:"goodId" db:"goodId"`
	GoodKey  string  `json:"goodKey" db:"goodKey"`
	GoodName string  `json:"goodName" db:"goodName"`
	Category string  `json:"category" db:"category"`
	Cat2nd   string  `json:"cat2nd" db:"cat2nd"`
	MinPrice float32 `json:"minPrice" db:"minPrice"`
	MaxPrice float32 `json:"maxPrice" db:"maxPrice"`
	OnShelf  bool    `json:"onShelf" db:"onShelf"`
}

const (
	SQL_CRE_MGOOD = `CREATE TABLE MGood ("goodId" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	  "goodKey" char(200) NOT NULL, 
	  "goodName" char(200) NOT NULL, 
	  "category" char(100) NOT NULL, 
	  "cat2nd" char(100) NOT NULL, 
	  "minPrice" FLOAT ,
	  "maxPrice" FLOAT ,
	  "onShelf" BOOLEAN );`

	SQL_SEL_MGOOD        = `SELECT goodId, goodKey, goodName , category, cat2nd , minPrice, maxPrice, onShelf FROM MGood  `
	SQL_SEL_MGOOD_CND1   = ` WHERE onShelf=1 `
	SQL_SEL_MGOOD_LIMIT  = ` ORDER BY goodId LIMIT `
	SQL_SEL_MGOOD_OFFSET = ` OFFSET `

	SQL_CRE_GDINFO = `CREATE TABLE MGoodInfo ("goodId" integer NOT NULL, "pic" char(200) NOT NULL);`

	SQL_INS_GOOD = `insert into MGood (
	"goodId",	"goodKey", 	"goodName", "category", "cat2nd", "minPrice","maxPrice","onShelf") 
	 VALUES (?,?,?,?,?,?,?,?)`

	SQL_INS_GDINFO = `insert into MGoodInfo (
		"goodId",	"pic") 
		 VALUES (?,?)`

	DB_FILE = `./lite.db`

	LITE3 = "sqlite3"
)

type IMGood interface {
	FindGoodById(gnd *MGood, goodid int) error
	GetGoodPix(goodid int) []string
	//input or update goods info
	InGood(gnd *MGood) error
	UpdateGood(gnd *MGood) error
	UpdateGoodPix(piclist []MGoodPic) error
	//init func
	InitTable() error
	InitLoadAllForUpdate() error
}

func (r MGood) InitTable() error {
	db, err := sqlx.Open(LITE3, DB_FILE)
	if err != nil {
		return err
	}
	_, err = db.Exec(SQL_CRE_MGOOD)
	if err != nil && strings.Contains(err.Error(), "already exists") {
		fmt.Printf("##ok:%v\n", err)
	} else if err != nil {
		return err
	}
	_, err = db.Exec(SQL_CRE_GDINFO)
	if err != nil && strings.Contains(err.Error(), "already exists") {
		fmt.Printf("##ok:%v\n", err)
	} else if err != nil {
		return err
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

func (r MGood) InGood(gd *MGood) error {
	db, err := sqlx.Open(LITE3, DB_FILE)
	if err != nil {
		return err
	}
	ngd := *gd
	_, err = db.Exec(SQL_INS_GOOD, ngd.GoodId, ngd.GoodKey, ngd.GoodName, ngd.Category, ngd.Cat2nd, ngd.MinPrice, ngd.MaxPrice, ngd.OnShelf)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

func (r MGood) UpdateGood(gd *MGood) error {
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
