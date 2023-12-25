package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("What you think we are doing here")
	inReader := bufio.NewReader(os.Stdin)
	userInput, err := inReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading:%v", err)
	}
	result := "Yes, we are motherfucket! " + strings.ToUpper(userInput)
	fmt.Println(result)
}
