package router

import (
	"github.com/gin-gonic/gin"

	"etov/internal/svc"
)

type Router struct {
	addons *svc.Addons
	engine *gin.Engine
}

func NewRouter(middle *svc.Addons, engine *gin.Engine) *Router {
	return &Router{addons: middle, engine: engine}
}

func (r *Router) routerHandler(handler svc.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context := svc.NewContextFromAddon(r.addons)
		context.Context = ctx
		handler(context)
	}
}

func (r *Router) POST(path string, handler svc.HandlerFunc) {
	r.engine.POST(path, r.routerHandler(handler))
}

func (r *Router) GET(path string, handler svc.HandlerFunc) {
	r.engine.GET(path, r.routerHandler(handler))
}
