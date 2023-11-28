package api

import (
	"strconv"

	"github.com/YogeLiu/medical/model"
	"github.com/YogeLiu/medical/service"
	"github.com/gin-gonic/gin"
)

type DetailBackend struct {
	svc *service.MedicineService
}

func NewDetailBackend() *DetailBackend {
	return &DetailBackend{svc: service.NewDetailService()}
}

type AddDetailReq struct {
	Name   string  `json:"name" binding:"required"`
	Amount int     `json:"amount"`
	Price  float32 `json:"price" binding:"required"`
}

func (b *DetailBackend) AddDetail(c *gin.Context) {
	var req AddDetailReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, gin.H{"code": 400, "message": err.Error()})
		return
	}

	detail := &model.Medicine{
		Name:   req.Name,
		Price:  req.Price,
		Amount: req.Amount,
	}
	if err := b.svc.AddDetail(detail); err != nil {
		c.JSON(200, gin.H{"code": 500, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"detail": detail})
}

func (b *DetailBackend) ListDetail(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}
	details, err := b.svc.ListDetail(page, pageSize)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": 0, "message": "", "object": details})
}

type UpdateDetailReq struct {
	OldName string  `json:"old_name" binding:"required"`
	NewName string  `json:"name"`
	Price   float32 `json:"price"`
	Amount  int     `json:"amount"`
}

func (b *DetailBackend) UpdateDetail(c *gin.Context) {
	req := UpdateDetailReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, gin.H{"code": 400, "message": err.Error()})
		return
	}
	err := b.svc.UpdateDetail(&model.Medicine{Name: req.NewName, Price: req.Price, Amount: req.Amount}, req.OldName)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 0})
}

type DeleteMedicineReq struct {
	Name string `json:"name" binding:"required"`
}

func (b *DetailBackend) DeleteDetail(c *gin.Context) {
	req := DeleteMedicineReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(200, gin.H{"code": 400, "message": err.Error()})
		return
	}
	err = b.svc.DeleteDetail(req.Name)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 0})
}
