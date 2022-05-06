package main

import (
	"fmt"
)

func a() {
	fmt.Println("Inside a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()!", c) //Recovery=> panicten gelen hatayı tutuyor
		} //panic=>hatanın kendisi (b den geliyor) - Panic ekrana yazdırmak istediğimiz hata mesajını tutar
	}() //deferde panic oluşması halinde function normal akışını
	//kesiyor ve bir sonraki kod satırı yani mainden devam ediyor

	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited!")
	fmt.Println("Exiting a().")
}

func b() {
	fmt.Println("Inside b()")
	panic("Panic in b()!")
	fmt.Println("Exiting b()")
}

func main() {
	a()
	fmt.Println("main() ended!")
}
