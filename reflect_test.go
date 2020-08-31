package yidoc

import (
	"fmt"
	"testing"
)

type AS string
type Arr []string
type MAp map[string]struct{}
type Fun func()
type ArrayMap []map[string]struct{}

/**
{
	type: string
	name: "",
}
*/

type A struct {
	Int                  int
	Str                  string
	IntArray             [3]int
	IntString            [3]string
	IntArrayArray        [][][][][][]int
	ExampleArray         [][][][][][]ExampleObject
	ExamplePtrArrayArray [][][][][][]*ExampleObject
	ObjectArray          [3]ObjectTest
	ObjectSlice          []*ExampleObject
	Object               *ExampleObject
	Bool                 bool
	Interface            interface{}
	Map                  map[string]interface{}
	ExampleObject
	MAp
	Arr
	ObjectOne ObjectTest
	AS
	Fun
}

type ObjectTest struct {
	Int         int
	Str         string
	IntArray    [3]int
	IntString   [3]string
	Slice       []string
	ObjectSlice []*ExampleObject
	Object      *ExampleObject
	Bool        bool
	Interface   interface{}
	Map         map[string]interface{}
}

type ExampleObject struct {
	UserName string
	Password string
}

func TestParseObject(t *testing.T) {
	a := new(A)
	obj := parseObject(a)

	_ = obj

	parseObject(Arr{})
	parseObject(1)
	parseObject("a")
	parseObject(true)
	parseObject(-1.1)
	parseObject([]int{111})
	parseObject(MAp{})
	parseObject(ArrayMap{})
	parseObject(Fun(func() {}))
}

type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Response
}

type Response struct {
	Age     int      `json:"age"`
	Uge     int      `json:"uge"`
	Arr     []string `json:"arr"`
	Boolean bool     `json:"boolean"`
}

func TestSwagger(t *testing.T) {
	doc := NewDoc(DefaultOptions()...)
	doc.Model(Request{})
	data, err := doc.Build()
	fmt.Println(string(data), err)
}
