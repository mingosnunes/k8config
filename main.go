/*
Copyright © 2022 Domingos Nunes mingosnunes94@gmail.com

*/
package main

import (
	"log"

	"github.com/mingosnunes/k8config/cmd"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatalln(err.Error())
	}
}
