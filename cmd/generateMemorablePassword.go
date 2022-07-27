package cmd

import (
	"fmt"
	"log"

	"github.com/rnemeth90/generator"
)

func main() {
	i, e := generator.GenerateMemorablePassword()
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(i)
}
