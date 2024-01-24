package router

import (
	"github.com/gin-gonic/gin"

	"etov/internal/svc"
)

type Router struct {
	addons      *svc.Addons
	engine      *gin.Engine
	basePath    string
	middlewares []svc.MiddleFunc
}

func NewRouter(middle *svc.Addons, engine *gin.Engine) *Router {
	return &Router{addons: middle, engine: engine, basePath: ""}
}

func (r *Router) routerHandler(handler svc.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context := svc.NewContextFromAddon(r.addons)
		context.Context = ctx
		for _, middleFunc := range r.middlewares {
			handler = middleFunc(handler)
		}
		handler(context)
	}
}

func (r *Router) Group(relativePath string, middlewares ...svc.MiddleFunc) *Router {
	middleFunc := append(r.middlewares, middlewares...)
	return &Router{
		addons:      r.addons,
		engine:      r.engine,
		basePath:    r.basePath + relativePath,
		middlewares: middleFunc,
	}
}

func (r *Router) POST(path string, handler svc.HandlerFunc) {
	r.engine.POST(r.basePath+path, r.routerHandler(handler))
}

func (r *Router) GET(path string, handler svc.HandlerFunc) {
	r.engine.GET(r.basePath+path, r.routerHandler(handler))
}

func (r *Router) Static(path string, root string) {
	r.engine.Static(path, root)
}
