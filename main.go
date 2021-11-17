package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
	//"rsc.io/quote"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
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
