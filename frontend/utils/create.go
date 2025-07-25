// Copyright 2025 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/Azure/ARO-HCP/internal/api"
	"github.com/Azure/ARO-HCP/internal/api/arm"
)

const (
	flagName = "type"
	cluster  = "cluster"
	nodePool = "node_pool"
)

func main() {
	example := "go run frontend/utils/create.go -type cluster"
	usage := fmt.Sprintf("type of object you want to create: %v or %v.\nExample: %v\n", cluster, nodePool, example)
	objectType := flag.String(flagName, cluster, usage)
	flag.Parse()

	if *objectType == cluster {
		err := CreateJSONFile()
		if err != nil {
			panic(err)
		}
		return
	}

	if *objectType == nodePool {
		err := CreateNodePool()
		if err != nil {
			panic(err)
		}
		return
	}

	help := "go run frontend/utils/create.go -type"
	panic(fmt.Sprintf("invalid object type, run: '%v'", help))
}

// CreateJSONFile creates a base cluster JSON file for use with testing frontend to create clusters
func CreateJSONFile() error {
	cluster := api.HCPOpenShiftCluster{
		Properties: api.HCPOpenShiftClusterProperties{
			Version: api.VersionProfile{
				ChannelGroup: "stable",
			},
			DNS: api.DNSProfile{},
			Network: api.NetworkProfile{
				NetworkType: api.NetworkTypeOVNKubernetes,
				PodCIDR:     "10.128.0.0/14",
				ServiceCIDR: "172.30.0.0/16",
				MachineCIDR: "10.0.0.0/16",
				HostPrefix:  23,
			},
			Console: api.ConsoleProfile{},
			API: api.APIProfile{
				Visibility: api.Visibility("Public"),
			},
			Platform: api.PlatformProfile{
				ManagedResourceGroup:   "dev-test-mrg",
				NetworkSecurityGroupID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/dev-test-rg/providers/Microsoft.Network/networkSecurityGroups/xyz",
				SubnetID:               "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/dev-test-rg/providers/Microsoft.Network/virtualNetworks/xyz/subnets/xyz",
				OutboundType:           api.OutboundType("LoadBalancer"),
				IssuerURL:              "",
			},
		},
	}

	data, err := json.MarshalIndent(cluster, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("cluster.json", data, 0643)
	if err != nil {
		return err
	}

	return nil
}

func CreateNodePool() error {
	nodePool := api.HCPOpenShiftClusterNodePool{
		Properties: api.HCPOpenShiftClusterNodePoolProperties{
			ProvisioningState: arm.ProvisioningState(""),
			Version: api.NodePoolVersionProfile{
				ChannelGroup: "stable",
			},
			Platform: api.NodePoolPlatformProfile{
				SubnetID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/dev-test-rg/providers/Microsoft.Network/virtualNetworks/xyz/subnets/xyz",
				OSDisk: api.OSDiskProfile{
					SizeGiB:                30,
					DiskStorageAccountType: "StandardSSD_LRS",
				},

				// VMSize should match configs/cloud-resources/instance-types.yaml
				// and configs/cloud-resource-constraints/instance-type-constraints.yaml
				// in CS config files.
				VMSize: "Standard_D8s_v3",
			},
			Replicas: 2,
		},
	}

	data, err := json.MarshalIndent(nodePool, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("node_pool.json", data, 0643)
	if err != nil {
		return err
	}

	return nil
}
