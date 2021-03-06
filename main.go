package main

import (
	"github.com/Neulhan/king10li-brokers/src/database"
	"github.com/Neulhan/king10li-brokers/src/handler"
	"github.com/Neulhan/king10li-brokers/src/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.DBLoader()
	h := handler.NewHandler(db)

	router := gin.Default()
	router.Use(middleware.CORS("*"))

	router.POST("/broker/create", h.CreateBroker)
	router.GET("/brokers/", h.GetBrokerList)
	router.GET("/broker/:id", h.GetOneBroker)
	router.DELETE("/broker/:id", h.DeleteOneBroker)

	router.Run()
}
