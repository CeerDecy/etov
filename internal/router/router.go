package router

import (
	"github.com/gin-gonic/gin"

	"etov/internal/svc"
)

type Router struct {
	middle *svc.Addons
	engine *gin.Engine
}

func NewRouter(middle *svc.Addons, engine *gin.Engine) *Router {
	return &Router{middle: middle, engine: engine}
}

func (r *Router) routerHandler(handler svc.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context := svc.NewContextFromMiddleWare(r.middle)
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
