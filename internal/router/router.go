package router

import (
	"github.com/gin-gonic/gin"

	"etov/internal/svc"
)

type Router struct {
	ctx    *svc.Context
	engine *gin.Engine
}

func NewRouter(ctx *svc.Context, engine *gin.Engine) *Router {
	return &Router{ctx: ctx, engine: engine}
}

func (r *Router) routerHandler(handler svc.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r.ctx.Context = ctx
		handler(r.ctx)
	}
}

func (r *Router) POST(path string, handler svc.HandlerFunc) {
	r.engine.POST(path, r.routerHandler(handler))
}

func (r *Router) GET(path string, handler svc.HandlerFunc) {
	r.engine.GET(path, r.routerHandler(handler))
}
