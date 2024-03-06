package gin

import (
	"net/http"

	"github.com/ArielFarias/poc-go-api/book"
	"github.com/gin-gonic/gin"
)

func Handlers(s book.UseCase) *gin.Engine {
	r := gin.Default()
	r.GET("/health", healthHandler)
	r.Handle("GET", "/", getBooks(s))
	return r
}

func getBooks(s book.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		books, err := s.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		ctx.JSON(http.StatusOK, books)

	}
}

func healthHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "App is healthy")
}
