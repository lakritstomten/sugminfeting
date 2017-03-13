package main

import (
	"fmt"
	"time"
)

//Commenting the main function.
func main() {
	fmt.Println("Hello world. I'm alive!!!!")
	<-time.After(5 * time.Second)
	fmt.Println("But now I'm about to die... :(")
}
