package main

import (
	"fmt"
	"example.com/greetings"
	"log"
)

func main(){
	messages, err:=greetings.Hello("Tehillah")
	if err!= nil{
		log.Fatal(err)
	}
	fmt.Println(messages)
}