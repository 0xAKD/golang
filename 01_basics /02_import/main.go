package main

//importing a single library.
//import "fmt"

//importing multiple libraries
import (
	"fmt"
	foo "net/http" //importing library with a custom name
)

func main() {
	fmt.Println("Hey, there i am coding Go.!!")

	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("HTTPS Response Status: ", resp.Status)

}
