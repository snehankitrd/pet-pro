package usecase

type DataProviderService interface {
	GetMenu() Menu
	GetMenuItem(id string) Item
}
