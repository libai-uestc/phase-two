package io_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
)

type User struct {
	Id       int
	Name     string
	Weight   float32
	BodyTall float32
	Age      int
	Sex      byte `json:"gender"`
	Ignore   int  `json:"-" gorm:"-"`
}

type Book struct {
	ISBN     string `json:"isbn"`
	Name     string
	Price    float32  `json:"price"`
	Author   *User    `json:"author"`
	Keywords []string `json:"kws"`
	Local    map[int]bool
}

var (
	user = User{
		Name:   "钱钟书",
		Age:    57,
		Sex:    1,
		Ignore: 1,
	}
	book = Book{
		ISBN:     "1234567890",
		Name:     "围城",
		Price:    34.8,
		Author:   &user,
		Keywords: []string{"婚姻", "爱情", "民国"},
		Local:    map[int]bool{2: true, 3: false},
	}
)

func Init() {
	bs, _ := json.Marshal(book)
	fmt.Printf("序列化之后的字节数%d\n", len(bs))
}

func BenchmarkJsonStd(b *testing.B) {
	b.ResetTimer()
	var inst Book
	for b.Loop() {
		bs, _ := json.Marshal(book)
		json.Unmarshal(bs, &inst)
	}
}

func BenchmarkJsonSonic(b *testing.B) {
	var inst Book
	for b.Loop() {
		bs, _ := sonic.Marshal(book)
		sonic.Unmarshal(bs, &inst)
	}
}

func BenchmarkJsoniter(b *testing.B) {
	var inst Book
	for b.Loop() {
		bs, _ := jsoniter.Marshal(book)
		jsoniter.Unmarshal(bs, &inst)
	}
}

func BenchmarkJsonGo(b *testing.B) {
	var inst Book
	for b.Loop() {
		bs, _ := gojson.Marshal(book)
		gojson.Unmarshal(bs, &inst)
	}
}

// go test ./io -bench=^BenchmarkJson -run=^$
// goos: windows
// goarch: amd64
// pkg: libai/go/phase-two/io
// cpu: Intel(R) Core(TM) i9-14900HX
// BenchmarkJsonStd-32               298156              4309 ns/op
// BenchmarkJsonSonic-32            1131415              1049 ns/op
// BenchmarkJsoniter-32              616136              1912 ns/op
// BenchmarkJsonGo-32                898599              1328 ns/op
// PASS
// ok      libai/go/phase-two/io   5.762s
// PS C:\Users\18101\Desktop\第二阶段\phase-two>
