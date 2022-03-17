package main

import (
	"fmt"
	"log"
)

func main() {
	cluster, err := NewLocalKindCluster()
	if err != nil {
		log.Fatal(err)
	}
	cluster.Provision()
	fmt.Println("Testing")
}
