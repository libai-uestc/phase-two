package main

import (
	"io"
	"libai/go/phase-two/web/idl"
	"net/http"
	"os"

	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
)

func text(engine *gin.Engine) {
	engine.GET("/user/text", func(c *gin.Context) {
		c.String(http.StatusOK, "hi libai")
	})
}

func json0(engine *gin.Engine) {
	engine.GET("/user/json0", func(c *gin.Context) {
		var stu struct {
			Name    string `json:"name"`
			Address string `json:"addr"`
		}
		stu.Name = "libai"
		stu.Address = "jzh"
		s, _ := sonic.MarshalString(stu)
		c.Request.Header.Add("Content-Type", "application/json")
		c.String(http.StatusOK, s)
	})
}

func json1(engine *gin.Engine) {
	engine.GET("/user/json1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"name": "libai", "addr": "jzh"})
	})
}

func json2(engine *gin.Engine) {
	var stu struct {
		Name    string `json:"name"`
		Address string `json:"addr"`
	}
	stu.Name = "libai"
	stu.Address = "jzh"
	engine.GET("/user/json2", func(c *gin.Context) {
		c.JSON(http.StatusOK, stu)
	})
}

func jsonp(engine *gin.Engine) {
	var stu struct {
		Name    string `json:"name"`
		Address string `json:"addr"`
	}
	stu.Name = "libai"
	stu.Address = "jzh"
	engine.GET("/user/jsonp", func(ctx *gin.Context) {
		ctx.JSONP(http.StatusOK, stu)
	})
}

func xml(engine *gin.Engine) {
	type Stu struct {
		Name    string `xml:"name"`
		Address string `xml:"addr"`
	}
	var stu Stu
	stu.Name = "libai"
	stu.Address = "jzh"
	engine.GET("/user/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, stu)
	})
}

func yaml(engine *gin.Engine) {
	var stu struct {
		Name string
		Addr string
	}
	stu.Name = "libai"
	stu.Addr = "jzh"
	engine.GET("/user/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, stu)
	})
}

func protoBuf(engine *gin.Engine) {
	stu := &idl.Student{
		Name:    "libai",
		Address: "jzh",
	}
	engine.GET("/user/pb", func(ctx *gin.Context) {
		ctx.ProtoBuf(http.StatusOK, stu)
	})
}

func html(engine *gin.Engine) {
	engine.LoadHTMLFiles("web/static/template.html", "web/static/student.html")
	engine.GET("/user/html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template.html", gin.H{"title": "用户信息", "name": "libai", "addr": "jzh"})
	})
}

func redirect(engine *gin.Engine) {
	engine.GET("/user/old_page", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/user/html")
	})
}

func main9() {
	fout, _ := os.OpenFile("log/gin.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, fout)
	engine := gin.Default()
	engine.Static("/css", "web/static/css")
	engine.StaticFile("/favicon.ico", "web/static/img/迈克尔乔丹.png")
	text(engine)
	json0(engine)
	json1(engine)
	json2(engine)
	jsonp(engine)
	xml(engine)
	yaml(engine)
	protoBuf(engine)
	html(engine)
	redirect(engine)
	engine.Run("127.0.0.1:5678")
}
