package interceptor

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"

	"etov/internal/svc"
)

func Recover(next svc.HandlerFunc) svc.HandlerFunc {
	return func(ctx *svc.Context) {
		defer func() {
			if err := recover(); err != nil {
				var msg string
				switch v := err.(type) {
				case error:
					msg = v.Error()
					logrus.Error(printStackTrace(v))
				default:
					msg = fmt.Sprintf("%v", err)
					logrus.Error(msg)
				}
				ctx.ErrorMsg(msg)
			}
		}()
		next(ctx)
	}
}

// 打印堆栈信息
func printStackTrace(err interface{}) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:])
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%v", err))
	for _, pc := range pcs[0:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		builder.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return builder.String()
}
