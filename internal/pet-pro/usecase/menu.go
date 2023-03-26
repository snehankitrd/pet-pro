package usecase

type Menu struct {
	Items map[string]Item
}

type Item struct {
	Id    string
	Name  string
	Price float64
	Note  string
	Type  string
}
