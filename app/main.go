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

		if strings.HasPrefix(command,"type ") {
			typeCommand := strings.TrimSpace(command[5:])
			if(typeCommand == "echo" || typeCommand == "exit" || typeCommand == "type") {
				fmt.Println(typeCommand + " is a shell builtin")
			} else{
				fmt.Println(typeCommand + ": not found")
			}
		} else if command == "exit\n" {
			break
		} else if strings.HasPrefix(command,"echo ") {
			command := strings.TrimPrefix(command,"echo ")
			fmt.Println(command[:len(command)-1])
		} else {
			fmt.Println(command[:len(command)-1] + ": command not found")
		}
	}
}