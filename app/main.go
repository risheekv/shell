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
		if command == "exit\n" {
			break
		}
		fmt.Println(command[:len(command)-1] + ": command not found")
	}
}