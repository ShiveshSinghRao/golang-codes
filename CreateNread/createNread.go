package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
func readFile(fileName string){
	databyte,err:=ioutil.ReadFile(fileName)
	if err!=nil{
		panic(err)
	}
	fmt.Println("text inside the file is : ", databyte)
	fmt.Println("text always is read in byte format in golang so we need to convert into text by using string : ",string(databyte))
}
func main() {
	fmt.Println("Welcome bhai aaja ")
	stuff:= "file k andar ka maal h bhai "

	file,err := os.Create("./myfile.txt")
	if err!=nil{
		panic(err)
	}

	len,err:=io.WriteString(file,stuff)
	if err!=nil{
		panic(err)
	}
	fmt.Println("length is : ",len)
	
	defer file.Close()
	readFile("./myfile.txt");
}
// OUTPUT:
// A file create wiht content in it:= file k andar ka maal h bhai 
// Terminal outpur is :
// Welcome bhai aaja 
// length is :  28
// text inside the file is :  [102 105 108 101 32 107 32 97 110 100 97 114 32 107 97 32 109 97 97 108 32 104 32 98 104 97 105 32]
// text always is read in byte format in golang so we need to convert into text by using string :  file k andar ka maal h bhai 
