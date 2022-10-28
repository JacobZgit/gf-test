//package main
//
//import (
//	_ "gf-test/internal/packed"
//
//	"github.com/gogf/gf/v2/os/gctx"
//
//	"gf-test/internal/cmd"
//)
//
//func main() {
//	cmd.Main.Run(gctx.New())
//}

package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"time"
)

func main() {
	s := g.Server()
	s.SetSessionMaxAge(time.Minute)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/set", func(r *ghttp.Request) {
			r.Session.Set("time", gtime.Timestamp())
			r.Response.Write("ok")
		})
		group.ALL("/get", func(r *ghttp.Request) {
			r.Response.Write(r.Session.Data())
		})
		group.ALL("/del", func(r *ghttp.Request) {
			_ = r.Session.RemoveAll()
			r.Response.Write("ok")
		})
	})
	s.SetPort(8199)
	s.Run()
}
