package main

import "fmt"

type RolePolicy struct {
	PolicyDocument string
	PolicyName     string
}

func main() {
	rp := RolePolicy{"policydoc", "policyname"}
	fmt.Println(rp)
}
