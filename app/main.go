package main

import (
	"fmt"
	"bufio"
	"os"
)

var _ = fmt.Print;

func main(){
	fmt.Print("$ ")
	reader := bufio.NewReader(os.Stdin)
	text,error := reader.ReadString('\n')

	if error!=nil {
		fmt.Print("Error during printing")
	}

	//text = strings.TrimSpace(text);

	fmt.Println(text[:len(text)-1] + ": command not found")
}