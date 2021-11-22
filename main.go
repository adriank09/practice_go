package main

import (
	"fmt"
	"io"
	"math/cmplx"
	"math/rand"
	"time"
	//"rsc.io/quote"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

type I interface {
	M()
}
type T struct {
	S string
}
type Int struct {
	I int
}

// type T implements interface I
func (t T) M() {
	fmt.Println(t.S)
}

type MyError struct {
	When time.Time
	What string
}

func (err *MyError) String() string {
	return fmt.Sprintf("Occured at %v, reason: %v", err.When, err.What)
}

func err_run() *MyError {
	return &MyError{time.Now(), "Foo does not contain bar."}
}

func invoke(i I) {
	fmt.Printf("Type: %T\n", i)
	//fmt.Println(i.(type))
	if i == nil {
		fmt.Println("Object is nil")
		return
	}
	i.M()
}

func invoke2(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Value of integer:", v)
	case string:
		fmt.Println("Value of string:", v)
	case Vertex:
		fmt.Println("Value of Vertex:", v.X, v.Y)
	}
}

type IPAddr [4]byte

// exercise to make IPAddr String()able
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	return r.Read(b)
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("Name: %v, age: %v", s.Name, s.Age)
}

func main() {
	// r := strings.NewReader("Hello world")
	// b := make([]byte, 8)

	// for {
	// 	n, err := r.Read(b)
	// 	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	// 	fmt.Printf("b[:n] = %q\n", b[:n])
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }

	// m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	// fmt.Println(m.Bounds())
	// fmt.Println(m.At(0, 0).RGBA())
	// fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	var s Student
	s = Student{"Brian", 14}
	fmt.Println(s)
}
