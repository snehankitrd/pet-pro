package dataprovider

import (
	"github.com/snehankitrd/pet-pro/internal/pet-pro/usecase"
)

type DataProvider struct {
	menu usecase.Menu
}

func NewService(db map[string]*usecase.Item) usecase.DataProviderService {
	dataProvider := DataProvider{menu: usecase.Menu{Items: db}}
	return &dataProvider
}

func (dp *DataProvider) GetMenu() usecase.Menu {
	return dp.menu
}

func (dp *DataProvider) GetMenuItem(id string) usecase.Item {
	return *(dp.menu.Items)[id]
}

func (dp *DataProvider) DeleteMenuItem(id string) bool {
	if _, ok := (dp.menu.Items)[id]; ok {
		delete(dp.menu.Items, id)
		return ok
	}
	return false
}

func (dp *DataProvider) AddMenuItem(item usecase.Item) bool {
	if _, ok := dp.menu.Items[item.Id]; !ok {
		dp.menu.Items[item.Id] = &item
		return true
	}
	return false
}

func (dp *DataProvider) UpdateMenuItem(item usecase.Item) (usecase.Item, bool) {
	if _, ok := dp.menu.Items[item.Id]; ok {
		if item.Name != "" {
			dp.menu.Items[item.Id].Name = item.Name
		}
		if item.Note != "" {
			dp.menu.Items[item.Id].Note = item.Note
		}
		if item.Price != 0 {
			dp.menu.Items[item.Id].Price = item.Price
		}
		if item.Type != "" {
			dp.menu.Items[item.Id].Type = item.Type
		}

		return *dp.menu.Items[item.Id], true
	}
	return usecase.Item{}, false
}
