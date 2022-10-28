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
)

func main() {
	s := g.Server()
	s.BindHandler("/cookie", func(r *ghttp.Request) {
		datetime := r.Cookie.Get("datetime")
		//r.Cookie.Set("datetime", gtime.Datetime())
		r.Response.Write("datetime:", datetime)
	})
	s.SetPort(8199)
	s.Run()
}
