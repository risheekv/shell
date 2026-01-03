package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	for {
		fmt.Print("$ ")
		reader := bufio.NewReader(os.Stdin)
		command,error := reader.ReadString('\n')

		if error != nil {
			fmt.Println("Error reported")
		}

		fmt.Println(command[:len(command)-1] + ": invalid command")
	}
}