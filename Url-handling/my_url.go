package main

import (
	"fmt"
	"net/url"
)

const my_url = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main(){
	fmt.Println("Welcome to handing URls in golang")
	fmt.Println(my_url)
	// parsing of url as it has many parameters
	result,_ :=url.Parse(my_url)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println("result Raw quesry is : ",result.RawQuery)
	
	q_params := result.Query()
	fmt.Printf("the type of query params are : %T\n",q_params)
	fmt.Println(q_params["coursename"])

	for _,val:=range q_params{
		fmt.Println("Param is : ",val)
	}
    partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path:	"/tutcss",
		RawPath: "user=hitesh",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}