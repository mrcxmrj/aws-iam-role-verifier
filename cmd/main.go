package main

import (
	"fmt"
	"github.com/mrcxmrj/aws-iam-role-verifier/roleverifier"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Wrong number of arguments!")
		os.Exit(1)
	}
	path := args[1]
	result, err := roleverifier.Verify(path)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
