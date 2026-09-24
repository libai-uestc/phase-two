package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v2"
	"io"
	"libai/go/phase-two/web/idl"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// go run ./web

type Student struct {
	Name    string `form:"username" uri:"user" json:"name" xml:"user" yaml:"user" binding:"required"`
	Address string `form:"addr" uri:"addr" json:"addr" xml:"addr" yaml:"addr" binding:"required"`
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

func PostForm(path string, stu Student) {
	fmt.Println("post form " + path)
	if resp, err := http.PostForm("http://127.0.0.1:5678"+path, url.Values{"username": {stu.Name}, "addr": {stu.Address}}); err != nil {
		panic(err)
	} else {
		processResponse(resp)
	}
}

func PostJson(path string, stu Student) {
	fmt.Println("post json " + path)
	if bs, err := json.Marshal(stu); err == nil {
		if resp, err := http.Post("http://127.0.0.1:5678"+path, "application/xml", bytes.NewReader(bs)); err != nil {
			panic(err)
		} else {
			processResponse(resp)
		}
	} else {
		slog.Error("xml marchal failed", "error", err)
	}
}

func PostXml(path string, stu Student) {
	fmt.Println("post xml " + path)
	if bs, err := xml.Marshal(stu); err == nil {
		if resp, err := http.Post("http://127.0.0.1:5678"+path, "application/xml", bytes.NewReader(bs)); err != nil {
			panic(err)
		} else {
			processResponse(resp)
		}
	} else {
		slog.Error("xml marchal failed", "error", err)
	}
}

func PostYaml(path string, stu Student) {
	fmt.Println("post yaml " + path)
	if bs, err := yaml.Marshal(stu); err == nil {
		if resp, err := http.Post("http://127.0.0.1:5678"+path, "application/json", bytes.NewReader(bs)); err != nil {
			panic(err)
		} else {
			processResponse(resp)
		}
	} else {
		slog.Error("json marchal failed", "error", err)
	}
}

func PostPb(path string, stu Student) {
	fmt.Println("post pb" + path)
	inst := idl.Student{Name: stu.Name, Address: stu.Address}
	if bs, err := proto.Marshal(&inst); err == nil {
		if resp, err := http.Post("http://127.0.0.1:5678"+path, "application/x-protobuf", bytes.NewReader(bs)); err != nil {
			panic(err)
		} else {
			processResponse(resp)
		}
	} else {
		slog.Error("yaml marchal failed", "error", err)
	}
}

func PostAll(path string, stu Student) {
	PostForm(path, stu)
	PostJson(path, stu)
	PostXml(path, stu)
	PostYaml(path, stu)
	PostPb(path, stu)
}

func Request(path, method string, body []byte) {
	request, err := http.NewRequest(method, "http://127.0.0.1:5678"+path, bytes.NewReader(body))
	if err != nil {
		panic(err)
	}
	request.AddCookie(
		&http.Cookie{
			Name:  "token",
			Value: "ye38ry4928---",
		},
	)
	client := &http.Client{
		Timeout: 500 * time.Millisecond,
	}
	if resp, err := client.Do(request); err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			io.Copy(os.Stdout, resp.Body)
			return
		}
		fmt.Println("response header:")
		if values, exists := resp.Header["Set-Cookie"]; exists {
			fmt.Println(values[0])
			cookie, _ := http.ParseSetCookie(values[0])
			fmt.Println("Name:", cookie.Name)
			fmt.Println("Value:", cookie.Value)
			fmt.Println("Domain:", cookie.Domain)
			fmt.Println("MaxAge:", cookie.MaxAge)
			fmt.Println(strings.Repeat("-", 50))
		}
		os.Stdout.WriteString("\n\n")
	}
}

func main() {
	student := Student{Name: "李白", Address: "江浙沪"}
	// Get("/home")
	// Get("/home")
	PostAll("/stu/multi_type", student)
}
