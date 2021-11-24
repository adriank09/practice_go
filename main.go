package main

import (
	"fmt"
	"io"
	"math/cmplx"
	"math/rand"
	"net/http"
	"os"
	"time"
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

func (s Student) IsAdult() bool {
	return s.Age >= 18
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

	// Demo student
	// var s Student
	// s = Student{"Brian", 14}
	// fmt.Println(s)
	// if s.IsAdult() {
	// 	fmt.Println("You may buy alcohol")
	// } else {
	// 	fmt.Println("You may not buy alcohol")
	// }

	// env := os.Environ()
	// for _, val := range os.Environ() {
	// 	if
	// }

	// // map
	// students := map[string]int{
	// 	"Ann": 13,
	// 	"Joe": 14,
	// }
	// person := "Kate"
	// if age, ok := students[person]; ok { // will be false if person is not in the map
	// 	fmt.Println(person, "is", age, "years old")
	// } else {
	// 	fmt.Println(person, "is not registered.")
	// }
	// person = "Ann"
	// if age, ok := students[person]; ok { // will be false if person is not in the map
	// 	fmt.Println(person, "is", age, "years old")
	// }
	// delete(students, "Ann")
	// fmt.Println(students)

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error while executing HTTP GET: ", err)
		return
	}
	//defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}

func main_test_array() {
	// array
	arr1 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr1)
	// array's values are copied in its entirety, so, to emulate C's copy by reference,
	// use &.
	arr2 := &arr1
	fmt.Println(arr2)
	arr1[0] = 9
	fmt.Println(arr1)
	fmt.Println(arr2)
}

func main_test_file() {
	file, err := os.Open("/Users/adriank09/Desktop/Untitled.rtf")
	if err != nil {
		fmt.Println(err)
		return
	}
	// make sure the opened file is closed later
	defer file.Close()

	var result []byte
	b := make([]byte, 10240)
	for {
		n, err := file.Read(b[0:])
		result = append(result, b[0:n]...) // append is discussed later
		if err != nil {
			if err == io.EOF {
				break
			}
			return // f will be closed if we return here.
		}
	}

	fmt.Println(string(result))
}
