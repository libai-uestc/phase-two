package main

import (
	"encoding/json"
	"fmt"
	"io"
	"libai/go/phase-two/web/idl"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func url(engine *gin.Engine) {
	engine.GET("/student", func(ctx *gin.Context) {
		a := ctx.Query("name")
		b := ctx.DefaultQuery("addr", "China")
		ctx.String(http.StatusOK, a+" live in "+b)
	})
}

func restful(engine *gin.Engine) {
	engine.GET("/student/:name/*addr", func(ctx *gin.Context) {
		name := ctx.Param("name")
		addr := ctx.Param("addr")
		ctx.String(http.StatusOK, name+" live in "+addr)
	})
}

func postForm(engine *gin.Engine) {
	engine.POST("/student/form", func(ctx *gin.Context) {
		name := ctx.PostForm("username")
		addr := ctx.DefaultPostForm("addr", "China")
		ctx.String(http.StatusOK, name+" live in "+addr)
	})
}

func postJson(engine *gin.Engine) {
	engine.POST("/student/json", func(ctx *gin.Context) {
		var stu Student
		bs, _ := io.ReadAll(ctx.Request.Body)
		if err := json.Unmarshal(bs, &stu); err == nil {
			name := stu.Name
			addr := stu.Addr
			ctx.String(http.StatusOK, name+" live in "+addr)
		}
	})
}

func ShouldBindBodyWith(ctx *gin.Context) error {
	var body []byte
	if v, exists := ctx.Get("rbody"); exists {
		if bs, ok := v.([]byte); ok {
			body = bs
		}
	}
	if body == nil {
		body, _ = io.ReadAll(ctx.Request.Body)
		ctx.Set("rbody", body)
	}
	var stu Student
	return json.Unmarshal(body, &stu)
}

func Handler(ctx *gin.Context) {
	var stu Student
	ctx.BindJSON(&stu)
	ctx.String(200, stu.Name+stu.Addr)
}

func upload_file(engine *gin.Engine) {
	engine.MaxMultipartMemory = 8 << 20
	engine.POST("/upload", func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err != nil {
			fmt.Printf("get file error %v\n", err)
			ctx.String(http.StatusInternalServerError, "upload file failed")
		} else {
			if err = ctx.SaveUploadedFile(file, "./data/"+file.Filename); err == nil {
				ctx.String(http.StatusOK, file.Filename)
			} else {
				fmt.Printf("save file to %s failed: %v\n", "./data/"+file.Filename, err)
			}
		}
	})
}

func upload_multi_file(engine *gin.Engine) {
	engine.POST("/upload_files", func(ctx *gin.Context) {
		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			files := form.File["files"]
			for _, file := range files {
				ctx.SaveUploadedFile(file, "./data/"+file.Filename)
			}
			ctx.String(http.StatusOK, "upload"+strconv.Itoa(len(files))+" files")
		}
	})
}

type Student struct {
	Name     string   `form:"username" json:"name" uri:"user" xml:"user" yaml:"user" binding:"required"`
	Addr     string   `form:"addr" json:"addr" uri:"addr" xml:"addr" yaml:"addr" binding:"required"`
	Keywords []string `form:"keywords"`
}

func formBind(engine *gin.Engine) {
	engine.POST("/stu/form", func(ctx *gin.Context) {
		var stu Student
		if err := ctx.ShouldBind(&stu); err != nil {
			fmt.Println(err)
			ctx.String(http.StatusBadRequest, "parse paramter failed")
		} else {
			ctx.String(http.StatusOK, stu.Name+" live in "+stu.Addr)
		}
	})
}

func jsonBind(engine *gin.Engine) {
	engine.POST("/stu/json", func(ctx *gin.Context) {
		var stu Student
		if err := ctx.ShouldBindJSON(&stu); err != nil {
			fmt.Println(err)
			ctx.String(http.StatusBadRequest, "parse paramter failed")
		} else {
			ctx.String(http.StatusOK, stu.Name+" live in "+stu.Addr)
		}
	})
}

func uriBind(engine *gin.Engine) {
	engine.GET("/stu/uri/:user/*addr", func(ctx *gin.Context) {
		var stu Student
		if err := ctx.ShouldBindUri(&stu); err != nil {
			fmt.Println(err)
			ctx.String(http.StatusBadRequest, "parse paramter failed")
		} else {
			ctx.String(http.StatusOK, stu.Name+" live in "+stu.Addr)
		}
	})
}

func xmlBind(engine *gin.Engine) {
	engine.POST("/stu/xml", func(ctx *gin.Context) {
		var stu Student
		if err := ctx.ShouldBindXML(&stu); err != nil {
			fmt.Println(err)
			ctx.String(http.StatusBadRequest, "parse paramter failed")
		} else {
			ctx.String(http.StatusOK, stu.Name+" live in "+stu.Addr)
		}
	})
}

func yamlBind(engine *gin.Engine) {
	engine.POST("/stu/yaml", func(ctx *gin.Context) {
		var stu Student
		if err := ctx.ShouldBindYAML(&stu); err != nil {
			fmt.Println(err)
			ctx.String(http.StatusBadRequest, "parse paramter failed")
		} else {
			ctx.String(http.StatusOK, stu.Name+" live in "+stu.Addr)
		}
	})
}

// func multiBind(engine *gin.Engine) {
// 	engine.POST("/stu/multi_type", func(ctx *gin.Context) {
// 		var stu Student
// 		var stu2 idl.Student
// 		if err := ctx.ShouldBindBodyWith(&stu, binding.JSON); err == nil {
// 			ctx.String(http.StatusOK, stu.Name+" live in "+stu.Addr)
// 		} else if err := ctx.ShouldBindBodyWith(&stu, binding.XML); err == nil {
// 			ctx.String(http.StatusOK, stu.Name+" live in "+stu.Addr)
// 		} else if err := ctx.ShouldBindBodyWith(&stu, binding.YAML); err == nil {
// 			ctx.String(http.StatusOK, stu.Name+" live in "+stu.Addr)
// 		} else if err := ctx.ShouldBindBodyWith(&stu2, binding.ProtoBuf); err == nil {
// 			ctx.String(http.StatusOK, stu2.Name+" live in "+stu2.Address)
// 		} else {
// 			ctx.String(http.StatusBadRequest, "不支持的参数类型")
// 		}
// 	})
// }

func multiBind(engine *gin.Engine) {
	engine.POST("/stu/multi_type", func(ctx *gin.Context) {
		var stu Student
		var stu2 idl.Student

		// Form 一般直接用 ShouldBind 即可
		if err := ctx.ShouldBindBodyWith(&stu, binding.JSON); err == nil {
			ctx.String(http.StatusOK, stu.Name+" live in "+stu.Addr)
		} else if err := ctx.ShouldBindBodyWith(&stu, binding.XML); err == nil {
			ctx.String(http.StatusOK, stu.Name+" live in "+stu.Addr)
		} else if err := ctx.ShouldBindBodyWith(&stu, binding.YAML); err == nil {
			ctx.String(http.StatusOK, stu.Name+" live in "+stu.Addr)
		} else if err := ctx.ShouldBind(&stu); err == nil { // 增加对 Form 的兼容处理
			ctx.String(http.StatusOK, stu.Name+" live in "+stu.Addr)
		} else if err := ctx.ShouldBindBodyWith(&stu2, binding.ProtoBuf); err == nil {
			ctx.String(http.StatusOK, stu2.Name+" live in "+stu2.Address)
		} else {
			ctx.String(http.StatusBadRequest, "不支持的参数类型")
		}
	})
}

func main5() {
	engine := gin.Default()
	url(engine)
	restful(engine)
	postForm(engine)
	postJson(engine)
	upload_file(engine)       //用postman模拟一个post请求，注意body类型选择form-data，Key名称为file，类型为File，在Value里选择本地文件
	upload_multi_file(engine) //用postman模拟一个post请求，注意body类型选择form-data，Key名称为files，类型为File，在Value里选择多个本地文件

	formBind(engine)
	jsonBind(engine)
	uriBind(engine)
	xmlBind(engine)
	yamlBind(engine)

	multiBind(engine)

	engine.Run("127.0.0.1:5678")
}
