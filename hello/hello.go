package main

import "fmt"
import "rsc.io/quote"

func Hello(name string)  string  {
	message := fmt.Sprintf("Hi %v , welcome!",name)
	return message
}
func main()  {
	fmt.Println(Hello("Tehillah"))
	fmt.Println(quote.Go())
}