package cat2nd

const (
	INPUTE_RROR = "#input format Error#"
)

type ICat2nd interface {
	GetCat2ndItems(name Cat2ndRequest) []Cat2nd
}

type Cat2nd struct {
	CategoryId   int    `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	Cat2ndId     int    `json:"cat2ndId"`
	Cat2ndName   string `json:"cat2ndName"`
}

func (s Cat2nd) GetCat2ndItems(c2nd Cat2ndRequest) []Cat2nd {
	ss := []Cat2nd{
		{CategoryName: "品牌", CategoryId: 2, Cat2ndId: 1, Cat2ndName: "内野"},
		{CategoryName: "品牌", CategoryId: 2, Cat2ndId: 2, Cat2ndName: "414"},
	}
	return ss
}
