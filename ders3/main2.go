 package main

import (
	"fmt"
	"os"
	"bufio"
	"string"
)


func main () {

		
	filepath := "yasi uygun olanlar"

	file, err := os.openFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644 )
	if err != nil {
		fmt.println("file does not open ", err)
		return
	}
	defer file.close() 
	
	scanner := bufio.newscanner(os.stdin)
	fmt.println("metin yazin")

	customerChannel := make(chan string)
	go writeToFile(customerChannel)}



func main() {

	filepath := "yasi uygun olmayanlar"

file, err := os.openFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644 )
if err != nil {
	fmt.println("file does not open ", err)
	return
}
defer file.close() 

scanner := bufio.newscanner(os.stdin)
fmt.println("metin yazin")

    CUstomerChannel2 := make(chan string)
	go writeToFile(customerChannel)




for {
	
	var age int
	fmt.println("enter your age")
	fm.scanln(&age)
	input := scanner.text()

	if strings.ToLOwer(input) == "q" {
		break
	}

	err := file.writingstring(input + "\n" )
	if err != nil {
		fmt.println("dosyayi yazma hatasi", err)
		return
	}	

}}


func main() {

	var nargileKafeYas = 18
	var CustomerAge int
	fmt.Println("Lütfen yaşınızı giriniz")
	fmt.Scan(&CustomerAge)

	if CustomerAge >= nargileKafeYas {
		fmt.Println("Hoşgeldiniz")
		customerChannel <- "18 ve 18 den buyuk kullanicilar"

	} else {
		fmt.Println("velinizle geliniz.")
		customerChanne2 <- "18 den kucuk olan kullanicilar"
	}
}
	q