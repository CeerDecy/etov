package main

import (
	"github.com/gin-gonic/gin"

	"etov/internal/router"
)

func main() {
	engine := gin.Default()
	router.DefaultRouter(engine)
	_ = engine.Run(":8181")
}
