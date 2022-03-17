package framework

import "fmt"

// TODO golang naming conventions for packages?

// Functions to create and manipulate clusters for end-to-end tests

type Cluster struct {
	Provisioned      bool
	NTHInstalled     bool
	TestPodInstalled bool
}

type Pod struct{}

type Node struct{}

// TODO golang naming conventions for constructors
func NewCluster() *Cluster {
	return &Cluster{}
}

func (c *Cluster) Provision() error {
	c.Provisioned = true
	fmt.Println("Provisioned the cluster!")
	return nil
}

// TODO polymorphism/inheritance--does it make sense to do?
// TODO Declare an interface that has these kinds of operations on it
func (c *Cluster) TearDown() error {
	c.Provisioned = false
	fmt.Println("Tore down the cluster")
	return nil
}

func (c *Cluster) InstallNodeTerminationHandler() (*Pod, *Node, error) {
	c.NTHInstalled = true
	fmt.Println("NTH installed")
	return &Pod{}, &Node{}, nil
}

func (c *Cluster) DeployTestPod(p *Pod) (*Node, error) {
	c.TestPodInstalled = true
	fmt.Println("Test pod deployed")
	return &Node{}, nil
}

func (c *Cluster) SimulateITN(n *Node) error {
	fmt.Println("Simulating an ITN")
	return nil
}

// TODO golang library for XUnit output
