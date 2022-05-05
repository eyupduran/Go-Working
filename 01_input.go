package main

import "fmt"

func main() {

	//BUFİO İLE KULLANICI GİRİŞİ ALMA
	// fmt.Print("Lütfen bir metin giriniz")
	// reader := bufio.NewReader(os.Stdin)
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(input)

	//SCANF İLE KULLANICI GİRİŞİ ALMA
	var sayi int
	fmt.Println("Bir sayi giriniz")
	fmt.Scanf("%d", &sayi)

	if sayi%2 == 0 {
		fmt.Println("çift sayı girdiniz")
	} else {
		fmt.Println("Tek sayı girdiniz")
	}

}
