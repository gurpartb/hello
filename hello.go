package main

import (
    "fmt"
    "github.com/gurpartb/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gurpartap")
    fmt.Println(message)
}