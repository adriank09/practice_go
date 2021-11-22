package main

import (
	"errors"
	"fmt"
	"image"
	"io"
	"log"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"

	"github.com/google/go-cmp/cmp"
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

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}

func main_types_and_string() {
	// v := Vertex{3, 4}
	// fmt.Println(v.Abs())
	// fmt.Println(v.Output())

	var i I = T{"Hello world"}
	i.M()

	invoke(i)

	var i2 I
	invoke(i2)

	// type assertion
	var i3 interface{} = "hello"
	r := i3.(string)
	// r := i3.(float64) // panic will occur - i3 is of type stirng, not float64
	fmt.Println(r)

	invoke2(3)
	invoke2("foo")
	invoke2(Vertex{1, 2})

	// exercise to make IPAddr String()able
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	if err := err_run(); err != nil {
		fmt.Println(err)
	}
}

func main_functions() {
	fmt.Println(Calculate(1, 2, Minus))
	fmt.Println(Calculate(4, 5, Multiply))
	fmt.Println(Calculate(2, 2, Add))

	// local function
	divide := func(n1 int, n2 int) int {
		return n1 / n2
	}

	fmt.Println(Calculate(10, 2, divide))
}

func main_map() {
	// map of values
	// method 1
	// var m map[string]int = make(map[string]int, 2)
	// m["Adam"] = 42
	// m["Brian"] = 26

	// method 2
	var m = map[string]int{
		"Adam": 42, "Brian": 26,
	}
	fmt.Println(m)

	// mutating a map
	// deleting an element from map
	delete(m, "Adam")
	fmt.Println(m)

	// checking the existence of an element in a map
	var elem, ok = m["Brian"]
	fmt.Println(elem, ok)

	elem, ok = m["Charlie"]
	fmt.Println(elem, ok)
}

func main_range() {
	var num = []int{2, 3, 4, 5, 6, 7, 8}
	// 1
	for i, v := range num {
		fmt.Printf("index %v, value: %v\n", i, v)
	}

	for _, v := range num {
		fmt.Println("Value is:", v)
	}
}

func main_array_slices() {
	// Arrays & slices
	// Array - fixed size
	// Slice - dynamically sized

	// 1
	var a [2]string
	fmt.Println(a)

	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 2
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 3
	s := []struct {
		A int
		B bool
	}{
		{1, true},
		{2, false},
		{3, true},
	}
	fmt.Println(s)

	// 4
	var s2 []int
	fmt.Println(len(s2), cap(s2))
	// Creates and allocates an array of the specified size
	s2 = make([]int, 3)
	fmt.Println(len(s2), cap(s2))

	// append
	var s3 []int
	fmt.Println(s3, len(s3), cap(s3))
	s3 = append(s3, 1)
	fmt.Println(s3, len(s3), cap(s3))
}

func main_struct() {
	v := Vertex{1, 2}
	fmt.Println(Vertex{1, 2})
	fmt.Println(v.X)
	fmt.Println(v.Y)

	v2 := Vertex{1, 2}
	p := &v2
	p.Y = 3
	fmt.Println(v2)
	fmt.Println(p)
}

func main_pointer() {
	// pointers
	i, j := 42, 7201

	p := &i
	fmt.Println(p)
	*p = 21
	fmt.Println(i)

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func main_4() {
	// Defer calls are executed in LIFO order
	// Defer 1
	defer fmt.Println("World")
	fmt.Println("Hello")

	// Defer 2
	fmt.Println("Counting...")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}

func main_3() {
	fmt.Println("The time is ", time.Now(), ". OK!")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("The sqrt for 64 is", math.Sqrt(64))
	fmt.Println(math.Pi)
	fmt.Println(swap("Hello", "World"))
	fmt.Println("The length of the string 'foo' is", string_length("foo"))

	// var i int
	// var foo bool
	var i, foo = 42, true
	// k := 3 is a shorthand for var k = 3
	k := 3
	fmt.Println("Value of i is", i)
	fmt.Println("Value of foo is", foo)
	fmt.Println("Value of k is", k)

	// for loop 1
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// // for loop 2
	// sum2 := 0
	// for i := 0; i < 10; {
	// 	sum2 += i
	// }
	// fmt.Println(sum2)

	// for loop 3
	sum3 := 0
	for sum3 < 10 {
		sum3++
	}
	fmt.Println(sum3)

	// for loop 4
	sum4 := 0
	// essentially is while loop
	for {
		if sum4 >= 20 {
			break
		}
		fmt.Println(sum4)
		sum4++
	}

	// runtime switch
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s. ¥n", os)
	}

	// switch
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today - 1:
		fmt.Println("Yesterday")
	case today + 2:
		fmt.Println("In two days.")
	}

	// switch 2
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() > 12 && t.Hour() < 18:
		fmt.Println("Good afternoon")
	case t.Hour() > 18:
		fmt.Println("Good evening")
	}
}

func pow(x, n, lim float64) float64 {
	// if with short statement
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func swap(x string, y string) (string, string) {
	return y, x
}

// Gets the length of a given string
// Named return demo
func string_length(s string) (length int) {
	length = len(s)
	return
}

func main_1() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// // Request greeting for Adrian
	// result, err := Hello("Adrian")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result)

	// A slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting for a slice of names
	result, err := Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func Add(n1 int, n2 int) int {
	return n1 + n2
}

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}
	// return a greeting that embeds the name in a message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages
	// Initialize
	messages := make(map[string]string)

	// Loop through the received slice of names, calling the
	// Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message
		// with the name.
		messages[name] = message
	}

	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
