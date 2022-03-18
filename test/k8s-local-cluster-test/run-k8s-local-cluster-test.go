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

package main

import (
	"fmt"
	"log"

	f "github.com/aws/aws-node-termination-handler/pkg/test/framework"
)

func main() {
	cluster, err := f.NewLocalKindCluster()
	if err != nil {
		log.Fatal(err)
	}
	var theCluster f.Cluster = cluster
	if err := theCluster.Provision(); err != nil {
		log.Fatal(err)
	}
	if err := theCluster.InstallNodeTerminationHandler(); err != nil {
		log.Fatal(err)
	}
	if err := theCluster.TearDown(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done with the test")
}
