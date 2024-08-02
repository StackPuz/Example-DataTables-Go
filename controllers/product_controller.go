package controllers

import (
	"app/config"
	"app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
}

func (con *ProductController) Index(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("length", "10"))
	start, _ := strconv.Atoi(c.Query("start"))
	order := "id"
	if c.Query("order[0][column]") != "" {
		order = c.Query("columns[" + c.Query("order[0][column]") + "][data]")
	}
	direction := c.DefaultQuery("order[0][dir]", "asc")
	var products []models.Product
	query := config.DB.Model(&products)
	var recordsTotal, recordsFiltered int64
	query.Count(&recordsTotal)
	search := c.Query("search[value]")
	if search != "" {
		search = "%" + search + "%"
		query.Where("name like ?", search)
	}
	query.Count(&recordsFiltered)
	query.Order(order + " " + direction).
		Offset(start).
		Limit(size).
		Find(&products)
	c.JSON(http.StatusOK, gin.H{"draw": c.Query("draw"), "recordsTotal": recordsTotal, "recordsFiltered": recordsFiltered, "data": products})
}
