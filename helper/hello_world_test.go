package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Zuhry")

	if result != "Hello Zuhry" {
		t.Fatal("Result is not 'Zuhry' ")
	}
}

func TestHelloWorldCulli(t *testing.T) {
	result := HelloWorld("Culli")

	if result != "Hello Culli" {
		t.Fatal("Result is not 'Culli' ")
	}

	fmt.Println("TestHelloWorld Done")
}
