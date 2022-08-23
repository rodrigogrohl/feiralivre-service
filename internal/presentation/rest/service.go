package rest

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/rodrigogrohl/feiralivre-service/internal/application"
	"github.com/rodrigogrohl/feiralivre-service/internal/infrastructure/config"
	"github.com/rodrigogrohl/feiralivre-service/pkg/canonical"
	"github.com/sirupsen/logrus"
)

const (
	streetMarket = "/v1/streetmarket/"
)

var (
	_onceRestService    sync.Once
	streetMarketService application.StreetMarketService
)

func InitRestService() {
	_onceRestService.Do(func() {
		streetMarketService = application.InitStreetMarketService()

		r := gin.Default()
		r.GET(streetMarket, GetStreetMarket)
		r.POST(streetMarket, CreateStreetMarket)
		r.DELETE(streetMarket, DeleteStreetMarket)
		r.POST(streetMarket+"query/", QueryStreetMarket)
		r.POST(streetMarket+"update/", UpdateStreetMarket)

		err := r.Run(fmt.Sprintf(":%d", config.RestPort))
		if err != nil {
			logrus.WithError(err).Panic()
		}
	})

}

func GetStreetMarket(c *gin.Context) {
	var kr KeyRequest
	if err := c.ShouldBindJSON(&kr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sm, err := streetMarketService.Get(c.Request.Context(), kr.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ToJSON(sm))
}

func DeleteStreetMarket(c *gin.Context) {
	var kr KeyRequest
	if err := c.ShouldBindJSON(&kr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	streetMarketService.Remove(c.Request.Context(), kr.Id)
}

func CreateStreetMarket(c *gin.Context) {
	var input canonical.StreetMarket
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := streetMarketService.Create(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": input.Id})
}

func UpdateStreetMarket(c *gin.Context) {
	var input canonical.StreetMarket
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := streetMarketService.Update(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func QueryStreetMarket(c *gin.Context) {
	var filter canonical.StreetMarketFilter
	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultList, err := streetMarketService.QueryBy(c.Request.Context(), &filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ListToJSON(resultList))
}
