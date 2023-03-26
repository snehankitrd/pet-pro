package menu

import (
	"encoding/json"
	"errors"
	"io"
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

func (ms *MenuService) AddMenuItem(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("error reading request body"))
		return
	}

	var item usecase.Item
	err = json.Unmarshal(body, &item)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, errors.New("error parsing request body"))
		return
	}

	if item.Id != "" {
		ok := ms.dataProvider.AddMenuItem(item)
		if ok {
			c.JSON(http.StatusCreated, item)
			return
		}
		c.JSON(http.StatusInternalServerError, errors.New("id already exists"))
	} else {
		c.JSON(http.StatusUnprocessableEntity, errors.New("invalid id"))
	}
}

func (ms *MenuService) DeleteMenuItem(c *gin.Context) {
	id := c.Param("id")
	if id != "" {
		ok := ms.dataProvider.DeleteMenuItem(id)
		if !ok {
			c.JSON(http.StatusNotFound, errors.New("id not found"))
			return
		}
		c.JSON(http.StatusOK, string("ok"))

	} else {
		c.JSON(http.StatusUnprocessableEntity, errors.New("invalid id"))
	}
}

func (ms *MenuService) UpdateMenuItem(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusUnprocessableEntity, errors.New("invalid id in URL"))
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("error reading request body"))
		return
	}

	var item usecase.Item
	err = json.Unmarshal(body, &item)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, errors.New("error parsing request body"))
		return
	}

	if id != item.Id || item.Id == "" {
		c.JSON(http.StatusUnprocessableEntity, errors.New("id mismatch in URL and request body"))
		return
	}

	updatedItem, ok := ms.dataProvider.UpdateMenuItem(item)
	if ok {
		c.JSON(http.StatusCreated, updatedItem)
		return
	}
	c.JSON(http.StatusInternalServerError, errors.New("id does not exist"))
}
