package controllers

import (
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/DubrovskijRD/budget_assistant_go/application/services"
	"github.com/DubrovskijRD/budget_assistant_go/domain/interfaces/repositories"
	"github.com/DubrovskijRD/budget_assistant_go/entrypoints/http/schemas"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReceiptController struct {
	service services.ReceiptService
	logger  *slog.Logger
}

func NewReceiptController(reseiptService services.ReceiptService, logger *slog.Logger) *ReceiptController {
	return &ReceiptController{
		service: reseiptService,
		logger:  logger,
	}
}

func (ctr ReceiptController) AddReceipt(c *gin.Context) {
	bid := c.Param("id")
	_, err := uuid.Parse(bid)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	receipt := schemas.Receipt{}

	if err := c.ShouldBindJSON(&receipt); err != nil {
		ctr.logger.Error("ShouldBindJSON", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	items := make([]repositories.ReceiptItemAdd, 0, len(receipt.Items))
	for _, i := range receipt.Items {
		items = append(items, repositories.ReceiptItemAdd{
			Name:   i.Name,
			Amount: i.Amount,
			Qty:    i.Qty,
		})
	}
	in := repositories.ReceiptAdd{
		BudgetId:    bid,
		Amount:      receipt.Amount,
		Items:       items,
		Description: receipt.Description,
		Labels:      receipt.Labels,
		Date:        receipt.Date,
	}
	result, err := ctr.service.AddReceipt(in)
	if err != nil {
		ctr.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// todo result to schema
	out := schemas.RespFromReceipt(result)
	c.JSON(200, gin.H{"result": out})
}

// https://github.com/gin-gonic/gin/issues/1516#issuecomment-1269846541
type CSStringList string

func (c CSStringList) Values() []string {
	str := string(c)
	if len(str) > 0 {
		values := strings.Split(str, ",")
		return values
	}
	return nil
}

type FilterQuery struct {
	Labels   CSStringList `form:"labels"`
	DateFrom time.Time    `form:"date_from"`
	DateTo   time.Time    `form:"date_to"`
}

func (ctr ReceiptController) GetReceipts(c *gin.Context) {
	bid := c.Param("id")
	_, err := uuid.Parse(bid)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	var query FilterQuery
	err = c.BindQuery(&query)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	receipts, err := ctr.service.GetReceipts(bid, query.Labels.Values(), query.DateFrom, query.DateTo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	out := schemas.RespFromReceiptList(receipts)
	c.JSON(200, gin.H{"result": out})
}

func (ctr ReceiptController) GetLabels(c *gin.Context) {
	bid := c.Param("id")
	_, err := uuid.Parse(bid)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	result, err := ctr.service.GetLabels(bid)
	if err != nil {
		ctr.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, gin.H{"result": result})
}
