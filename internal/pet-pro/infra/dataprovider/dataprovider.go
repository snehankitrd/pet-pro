package dataprovider

import (
	"github.com/snehankitrd/pet-pro/internal/pet-pro/usecase"
)

type DataProvider struct {
	db map[string]usecase.Item
}

func NewService(db map[string]usecase.Item) usecase.DataProviderService {
	dataProvider := DataProvider{db: db}
	return &dataProvider
}

func (dp *DataProvider) GetMenu() usecase.Menu {
	return usecase.Menu{Items: dp.db}
}

func (dp *DataProvider) GetMenuItem(id string) usecase.Item {
	return dp.db[id]
}
