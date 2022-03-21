// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License

package framework

import (
	"fmt"
	"os"
	"os/exec"
)

const K8_1_22 string = "kindest/node:v1.22.2@sha256:f638a08c1f68fe2a99e724ace6df233a546eaf6713019a0b310130a4f91ebe7f"
const K8_1_21 string = "kindest/node:v1.21.2@sha256:9d07ff05e4afefbba983fac311807b3c17a5f36e7061f6cb7e2ba756255b2be4"
const K8_1_20 string = "kindest/node:v1.20.70@sha256:cbeaf907fc78ac97ce7b625e4bf0de16e3ea725daf6b04f930bd14c67c671ff9"
const K8_1_19 string = "kindest/node:v1.19.11@sha256:07db187ae84b4b7de440a73886f008cf903fcf5764ba8106a9fd5243d6f32729"
const K8_1_18 string = "kindest/node:v1.18.19@sha256:7af1492e19b3192a79f606e43c35fb741e520d195f96399284515f077b3b622c"
const K8_1_17 string = "kindest/node:v1.17.17@sha256:66f1d0d91a88b8a001811e2f1054af60eef3b669a9a74f9b6db871f2f1eeed00"

func NewLocalKindCluster() (*LocalKindCluster, error) {
	fmt.Println("Created a new local Kind cluster object")
	return &LocalKindCluster{Name: GenerateUniqueName("testcluster")}, nil
}

type LocalKindCluster struct {
	Name             string
	Provisioned      bool
	NTHInstalled     bool
	TestPodInstalled bool
}

func (c *LocalKindCluster) Provision() error {
	fmt.Printf("Provisioning a kind cluster %s...\n", c.Name)
	// TODO implement retry logic?
	cmd := exec.Command("kind", "create", "cluster", "--name", c.Name) //, "--image", K8_1_22)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	c.Provisioned = true
	fmt.Println("Provisioned the cluster!")
	return nil
}

func (c *LocalKindCluster) Reset() error {
	fmt.Println("Resetting the cluster...")
	return nil
}

func (c *LocalKindCluster) TearDown() error {
	if !c.Provisioned {
		return fmt.Errorf("Cluster was not successfully provisioned; cannot tear down")
	}
	fmt.Printf("Tearing down kind cluster %s...\n", c.Name)
	cmd := exec.Command("kind", "delete", "cluster", "--name", c.Name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	c.Provisioned = false
	fmt.Println("Tore down the cluster")
	return nil
}

func (c *LocalKindCluster) InstallNodeTerminationHandler() error {
	c.NTHInstalled = true
	fmt.Println("NTH installed")
	return nil
}

func (c *LocalKindCluster) DeployTestPod(p *Pod) error {
	c.TestPodInstalled = true
	fmt.Println("Test pod deployed")
	return nil
}

func (c *LocalKindCluster) SimulateITN(n *Node) error {
	fmt.Println("Simulating an ITN")
	return nil
}

func (c *LocalKindCluster) KindSpecificOperation() error {
	fmt.Println("This is a Kind-specific operation")
	return nil
}
