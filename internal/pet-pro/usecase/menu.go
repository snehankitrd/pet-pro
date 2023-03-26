package usecase

type Menu struct {
	Items map[string]*Item `json:"items"`
}

type Item struct {
	Id    string  `json:"id" required:"true"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Note  string  `json:"note"`
	Type  string  `json:"type"`
}
