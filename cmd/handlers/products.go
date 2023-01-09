package handlers

import (
	"errors"
	"example/pkg/response"
	"example/services"
	"example/services/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Get(ctx *gin.Context) {
	// request
	// process
	products := services.Get()

	// response
	ctx.JSON(http.StatusOK, response.Ok("succeed to get products", products))
}

type request struct {
	Id           int     `json:"id" validate:"required"`
	Name         string  `json:"name" validate:"required"`
	Quantity     int     `json:"quantity" validate:"required"`
	Code_value   string  `json:"code_value" validate:"required"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration" validate:"required"`
	Price        float64 `json:"price" validate:"required"`
}

func ProductsPriceGt(ctx *gin.Context) {
	//request
	price, ok := ctx.GetQuery("priceGt")
	if !ok {
		Get(ctx)
		return
	}
	priceFloat, _ := strconv.ParseFloat(price, 64)
	productsPriceGt := services.GetPriceGt(priceFloat)
	if len(productsPriceGt) == 0 {
		message := fmt.Sprintf("No se econtraron productos con precio mayor a %f", priceFloat)
		ctx.JSON(200, response.Response{
			Message: message,
			Data:    nil,
		})
		return
	}
	message := fmt.Sprintf("Se encontraron %d productos.", len(productsPriceGt))
	ctx.JSON(200, response.Response{
		Message: message,
		Data:    productsPriceGt,
	})

}

func ProductByID(ctx *gin.Context) {
	// request
	id, ok := ctx.GetQuery("id")

	// process
	var p models.Product
	if !ok {
		ctx.JSON(http.StatusRequestedRangeNotSatisfiable, response.Err(errors.New("ID not specified")))
		return
	}
	idint, _ := strconv.Atoi(id)
	exists := services.ExistsProductId(idint)
	if !exists {
		message := fmt.Sprintf("No se econtraron productos con el id %d", idint)
		ctx.JSON(200, response.Response{
			Message: message,
			Data:    nil,
		})
		return
	}
	p = services.GetById(idint)
	// response
	message := fmt.Sprintf("Se encontro el siguente elemento con id %d", idint)
	ctx.JSON(200, response.Response{
		Message: message,
		Data:    p,
	})

}

func Pong(ctx *gin.Context) {
	// request
	// process
	// response
	ctx.String(200, "Pong")

}

func Create(ctx *gin.Context) {
	// request
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err(err))
		return
	}

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.Err(err))
		return
	}

	// process
	product, err := services.Create(req.Id, req.Name, req.Quantity, req.Code_value, req.Is_published, req.Expiration, req.Price)
	if err != nil {
		if errors.Is(err, services.ErrAlreadyExist) {
			ctx.JSON(http.StatusConflict, response.Err(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, response.Err(err))
		return
	}

	// response
	ctx.JSON(http.StatusCreated, response.Ok("suceed to create website", product))
}
