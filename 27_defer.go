package main

import (
	"fmt"
)

func d1() {
	for i := 3; i > 0; i-- { //defer sözcüğünün olduğu yerdeki ifade ters sırayla çalışacak
		defer fmt.Print(i, " ") //iş görecek elemanları stacke iter ve stackten çekip o şekilde işlemi tamamlar
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

//İçteki fonksiyonlar anonim fonksiyondur
//Anonim fonksiyonların en büyük özelliği isimsiz olmalarıdır.
// (Zaten adından da belli oluyor ) Yazıldıkları yerde direkt olarak çalışırlar.
//Çalışırken diğer fonksiyonlardaki gibi parametre
// verilemediği için fonksiyonun sonuna parametre eklenerek çalışıtırılırlar.

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i) //Function her ertelendiğinde i nin geçerli değerini alır, bu yüzden bu değerle birlikte bunu işler
	}
}

func main() {
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
}
