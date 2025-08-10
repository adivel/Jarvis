package main

import "fmt"

type Jarvis struct{}

func (j *Jarvis) SayHello(name string) string {
    return fmt.Sprintf("Hello, %s! I am Jarvis.", name)
}
