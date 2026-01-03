package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		} else if strings.HasPrefix(command,"echo ") {
			command := strings.TrimPrefix(command,"echo ")
			fmt.Println(command[:len(command)-1])
		} else {
			fmt.Println(command[:len(command)-1] + ": command not found")
		}
	}
}