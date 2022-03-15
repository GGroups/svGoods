package listgoods

const (
	INPUTE_RROR = "#input format Error#"
)

type IGoods interface {
	ListGoods(name string) []Good
}

type Good struct {
	GoodId   int     `json:"goodId"`
	GoodName string  `json:"goodName"`
	MinPrice float32 `json:"minPrice"`
	MaxPrice float32 `json:"maxPrice"`
	Pic      string  `json:"pic"`
}

func (s Good) ListGoods(t string) []Good {
	ss := []Good{
		{GoodId: 1, GoodName: "内野长绒棉", MinPrice: 12.8, MaxPrice: 16.8, Pic: `neiye/neiye长绒棉/1.jpg`},
	}
	return ss
}
