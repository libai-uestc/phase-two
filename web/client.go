package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"libai/go/phase-two/web/idl"
	"log/slog"
	"net/http"
	"net/url"
	"os"

	"go.yaml.in/yaml/v3"
	"google.golang.org/protobuf/proto"
)

// go run ./web

type Student struct {
	Name    string `form:"username" uri:"user" json:"name" xml:"user"`
	Address string `form:"addr" uri:"addr" json:"addr" xml:"addr"`
}

type User struct {
	Name       string `form:"name"`
	Score      int    `form:"score"`
	Enrollment string `form:"enrollment"`
	Graduation string `form:"graduation"`
}

func processResponse(resp *http.Response) {
	defer resp.Body.Close()
	fmt.Println("响应头: ")
	for k, v := range resp.Header {
		fmt.Printf("%s=%s\n", k, v[0])
	}
	fmt.Print("响应体: ")
	io.Copy(os.Stdout, resp.Body)
	os.Stdout.WriteString("\n")
	if resp.StatusCode != http.StatusOK {
		slog.Error("异常状态码", "http response code", resp.StatusCode)
	}
	os.Stdout.WriteString("\n")
}

func Head(path string) {
	fmt.Println("HEAD " + path)
	resp, err := http.Head("http://127.0.0.1:5678" + path)
	if err != nil {
		slog.Error("http get failed", "error", err)
		return
	}
	processResponse(resp)
}

func Get(path string) {
	fmt.Println("GET " + path)
	resp, err := http.Get("http://127.0.0.1:5678" + path)
	if err != nil {
		slog.Error("http get failed", "error", err)
		return
	}
	processResponse(resp)
}

func GetStudentPB(path string) {
	fmt.Println("GET " + path)
	resp, err := http.Get("http://127.0.0.1:5678" + path)
	if err != nil {
		slog.Error("http get failed", "error", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("响应头：")
	for k, v := range resp.Header {
		fmt.Printf("%s=%s\n", k, v[0])
	}
	if resp.StatusCode != http.StatusOK {
		slog.Error("异常状态码", "http response code", resp.StatusCode)
		io.Copy(os.Stdout, resp.Body)
	} else {
		bs, _ := io.ReadAll(resp.Body)
		var stu idl.Student
		if err := proto.Unmarshal(bs, &stu); err == nil {
			fmt.Printf("pb反序列化成功, name=%s,addr=%s\n", stu.Name, stu.Address)
		}
	}
	os.Stdout.WriteString("\n\n")
}

func PostForm(path string,stu Student){
	fmt.Println("post form " + path)
	if resp,err := http.PostForm("http://127.0.0.1:5678"+path,url.Values
	{"username":{stu.Name},"addr":{stu.Address}})
}

func PostJson(path string,stu Student){
	fmt.Println("post json "+ path)
	if bs,err := json.Marshal(stu);err == nil {
		if resp,err := http.Post("http://127.0.0.1:5678"+path,"application/xml",bytes.NewReader(bs));err != nil{
			panic(err)
		} else {
			processResponse(resp)
		}
	} else {
		slog.Error("xml marchal failed","error",err)
	}
}

func PostYaml(path string,stu Student){
	fmt.Println("post yaml " + path)
	if bs,err := yaml.Marshal(stu);err == nil {
		if resp,err := http.Post("http://127.0.0.1:5678"+path, "application/json",bytes.NewReader(bs)); err != nil {
			panic(err)
		} else {
			processResponse(resp)
		}
	} else {
		slog.Error("json marchal failed","error",err)
	}
}


func main() {
	Get("/home")
	// Get("/home")
}
