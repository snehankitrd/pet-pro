package usecase

type DataProviderService interface {
	GetMenu() Menu
	GetMenuItem(id string) Item
	DeleteMenuItem(id string) bool
	AddMenuItem(item Item) bool
	UpdateMenuItem(item Item) (Item, bool)
}
