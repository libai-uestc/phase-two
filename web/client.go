package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
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

func main() {
	Get("/home")
	// Get("/home")
}
