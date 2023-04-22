package main

import (
	"bytes"
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"
)

func main() {
	app := setupRouter()
	app.SetTrustedProxies(nil)
	app.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	// index.html
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK,
			"index.html",
			gin.H{},
		)
	})
	// turns text to funnytext
	r.GET("/funnytext", func(ctx *gin.Context) {
		data := ctx.DefaultQuery("text", "funny text")

		if data == "" {
			data = "funny text"
		}

		ctx.String(http.StatusOK, makeFunny(data))
	})

	return r
}

func makeFunny(text string) string {
	var output bytes.Buffer
	for pos, char := range text {
		if pos%2 == 0 {
			output.WriteRune(unicode.ToLower(char))
		} else {
			output.WriteRune(unicode.ToUpper(char))
		}
	}
	return output.String()
}
