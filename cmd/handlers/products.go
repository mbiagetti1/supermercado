package handlers

import (
	"errors"
	"example/pkg/response"
	"example/services"
	"net/http"

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