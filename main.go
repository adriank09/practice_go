package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
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

func main() {
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
