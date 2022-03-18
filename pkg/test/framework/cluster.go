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

// TODO golang naming conventions for packages?

// Functions to create and manipulate clusters for end-to-end tests

type Cluster interface {
	Provision() error
	Reset() error
	TearDown() error
	InstallNodeTerminationHandler() error
	DeployTestPod(p *Pod) error
	SimulateITN(n *Node) error
}
type Pod interface{}
type Node interface{}

// TODO golang library for XUnit output
