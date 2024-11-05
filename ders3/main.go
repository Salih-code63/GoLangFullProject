package main

import (

	"fmt"
	"os"
	"bufio"

	"string"
)
      



func main() {

	filepath := "otomatik_metn.txt"

file, err := os.openFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644 )
if err != nil {
	fmt.println("file does not open ", err)
	return
}
defer file.close() 

scanner := bufio.newscanner(os.stdin)
fmt.println("metin yazin")}


	


for {

	var fName string
    fmt.println("Enter first name")
	fmt.scanln(&fName)
	input := scanner.text()
    
	var lName string
	fmt.println("Enter Last Name")
	fmt.scanln(&lName)
	input := scanner.text()

	fmt.print("Full name is: " + fName + " " + lName)
	input := scanner.text()

	if strings.ToLOwer(input) == "q" {
		break
	}

	err := file.writingstring(input + "\n" )
	if err != nil {
		fmt.println("dosyayi yazma hatasi", err)
		return
	}



}()




