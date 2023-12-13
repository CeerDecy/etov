package main

import (
	"github.com/gin-gonic/gin"

	"etov/conf"
	"etov/internal/handle"
	"etov/internal/router"
	"etov/internal/svc"
)

func main() {
	engine := gin.Default()
	cfg := conf.EtovCfg
	r := router.NewRouter(svc.NewContext(&cfg), engine)
	handle.RegisterHandler(r)
	_ = engine.Run(":8181")
}
