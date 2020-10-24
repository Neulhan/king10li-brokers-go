package handler

import (
	"net/http"

	"github.com/Neulhan/king10li-brokers/src/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func (h Handler) GetBrokerList(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func (h Handler) GetOneBroker(c *gin.Context) {
	pk := c.Param("id")

	var broker database.Broker
	h.DB.First(&broker, pk)

	c.JSON(200, gin.H{
		"result": "success",
		"broker": broker,
	})
}

func (h Handler) CreateBroker(c *gin.Context) {
	type createBrokerProps struct {
		Code           string `json:"code" binding:"required"`           // 등록번호
		Name           string `json:"name" binding:"required"`           // 상호
		Address        string `json:"address" binding:"required"`        // 소재지
		Representative string `json:"representative" binding:"required"` // 대표자
		PhoneNumber    string `json:"phoneNumber" binding:"required"`    // 전화번호
		Status         string `json:"status" binding:"required"`         // 상태
	}

	var props createBrokerProps
	if err := c.ShouldBindJSON(&props); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	broker := &database.Broker{
		Code:           props.Code,
		Name:           props.Name,
		Address:        props.Address,
		Representative: props.Representative,
		PhoneNumber:    props.PhoneNumber,
		Status:         props.Status,
	}
	h.DB.Create(broker)
	c.JSON(200, gin.H{
		"result": "success",
		"broker": broker,
	})
}

func (h Handler) DeleteOneBroker(c *gin.Context) {
	pk := c.Param("id")

	var broker database.Broker
	h.DB.Delete(&broker, pk)

	c.JSON(200, gin.H{
		"result": "success",
		"broker": broker,
	})
	c.JSON(200, gin.H{})
}

func (h Handler) GetReview(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func (h Handler) DeleteReview(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func (h Handler) CreateReview(c *gin.Context) {
	c.JSON(200, gin.H{})
}
