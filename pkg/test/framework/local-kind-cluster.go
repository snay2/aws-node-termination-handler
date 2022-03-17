package framework

import "fmt"

func NewLocalKindCluster() (*Cluster, error) {
	fmt.Println("Created a new local Kind cluster object")
	return &Cluster{}, nil
}
