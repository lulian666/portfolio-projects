package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"portfolio/pkg"
)

type CreateProject struct {
	Name        string `json:"name" binding:"required,min=2"`
	Description string `json:"description" binding:"required"`
	SourceCode  string `json:"sourceCode" binding:"required"`
	Link        string `json:"link" binding:"required"`
}

type UpdateProject struct {
	Name string `json:"name" binding:"required"`
}

func validate[T any](c *gin.Context, v T) {
	valid, errs := pkg.BindAndValid(c, v)
	if !valid {
		response := pkg.NewResponse(c)
		response.ToErrorResponse(pkg.NewError(pkg.InvalidParams, "invalid parameters").WithDetails(
			errs.Errors()...,
		))
		response.Ctx.Abort()
	}

}

func Validator(v interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		validate(c, v)
		fmt.Println(v)
		c.Next()
	}
}
