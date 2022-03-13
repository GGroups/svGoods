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
		{},
	}
	return ss
}
