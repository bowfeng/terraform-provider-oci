// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	networkSecurityGroupVnicDataSourceRepresentation = map[string]interface{}{
		"network_security_group_id": Representation{repType: Required, create: `${oci_core_network_security_group.test_network_security_group1.id}`},
	}

	NetworkSecurityGroupVnicResourceConfig = VnicAttachmentResourceConfig
)

func TestCoreNetworkSecurityGroupVnicResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreNetworkSecurityGroupVnicResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_network_security_group_vnics.test_network_security_group_vnics"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + compartmentIdVariableStr + NetworkSecurityGroupVnicResourceConfig,
			},
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_network_security_group_vnics", "test_network_security_group_vnics", Required, Create, networkSecurityGroupVnicDataSourceRepresentation) +
					compartmentIdVariableStr + NetworkSecurityGroupVnicResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "network_security_group_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "network_security_group_vnics.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "network_security_group_vnics.0.resource_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "network_security_group_vnics.0.time_associated"),
					resource.TestCheckResourceAttrSet(datasourceName, "network_security_group_vnics.0.vnic_id"),
				),
			},
		},
	})
}
