// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateInstanceConfigurationFromInstanceDetails Create an instance configuration from an existing instance.
type CreateInstanceConfigurationFromInstanceDetails struct {

	// The OCID of the compartment containing the instance configuration.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The ID of the instance that will be used to create instance configuration.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name for the instance configuration
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

//GetCompartmentId returns CompartmentId
func (m CreateInstanceConfigurationFromInstanceDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDefinedTags returns DefinedTags
func (m CreateInstanceConfigurationFromInstanceDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetDisplayName returns DisplayName
func (m CreateInstanceConfigurationFromInstanceDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetFreeformTags returns FreeformTags
func (m CreateInstanceConfigurationFromInstanceDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

func (m CreateInstanceConfigurationFromInstanceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateInstanceConfigurationFromInstanceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateInstanceConfigurationFromInstanceDetails CreateInstanceConfigurationFromInstanceDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateInstanceConfigurationFromInstanceDetails
	}{
		"INSTANCE",
		(MarshalTypeCreateInstanceConfigurationFromInstanceDetails)(m),
	}

	return json.Marshal(&s)
}
