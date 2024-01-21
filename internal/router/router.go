package router

import (
	"github.com/gin-gonic/gin"

	"etov/internal/svc"
)

type Router struct {
	addons      *svc.Addons
	engine      *gin.Engine
	basePath    string
	middlewares []svc.HandlerFunc
}

func NewRouter(middle *svc.Addons, engine *gin.Engine) *Router {
	return &Router{addons: middle, engine: engine, basePath: ""}
}

func (r *Router) routerHandler(handler svc.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context := svc.NewContextFromAddon(r.addons)
		context.Context = ctx
		for _, middleFunc := range r.middlewares {
			middleFunc(context)
		}
		handler(context)
	}
}

func (r *Router) Group(relativePath string, middlewares ...svc.HandlerFunc) *Router {
	return &Router{
		addons:      r.addons,
		engine:      r.engine,
		basePath:    r.basePath + relativePath,
		middlewares: middlewares,
	}
}

func (r *Router) POST(path string, handler svc.HandlerFunc) {
	r.engine.POST(r.basePath+path, r.routerHandler(handler))
}

func (r *Router) GET(path string, handler svc.HandlerFunc) {
	r.engine.GET(r.basePath+path, r.routerHandler(handler))
}
