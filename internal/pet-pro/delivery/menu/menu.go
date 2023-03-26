package menu

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snehankitrd/pet-pro/internal/pet-pro/usecase"
)

type MenuService struct {
	dataProvider usecase.DataProviderService
}

func NewService(dataProvider usecase.DataProviderService) MenuService {
	return MenuService{dataProvider: dataProvider}
}

func (ms *MenuService) GetMenu(c *gin.Context) {

	c.JSON(http.StatusOK, ms.dataProvider.GetMenu())

}

func (ms *MenuService) GetMenuItem(c *gin.Context) {
	id := c.Param("id")
	if id != "" {
		c.JSON(http.StatusOK, ms.dataProvider.GetMenuItem(id))
	} else {
		c.JSON(http.StatusUnprocessableEntity, errors.New("invalid id"))
	}
}
