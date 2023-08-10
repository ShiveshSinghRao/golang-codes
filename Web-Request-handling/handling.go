package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
const url = "https://leetcode.com/shiveshsinghrao/"

func main(){
	fmt.Println("leetcode web request ")
	response,err:=http.Get(url)
	if err!=nil{
		panic(err)
	}
	fmt.Println("response of type : %T\n", response)
	defer response.Body.Close() // caller's responsibilty to close the connection
    
	databytes,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}
	fmt.Println("the databyte is ", string(databytes))
}