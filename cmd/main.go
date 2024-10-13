package main

import (
	"log"

	"github.com/bondzai/gogear/rest"
)

func main() {
	log.Println("test libs")
	params := map[string]string{}
	rest.Validate(params)
}
