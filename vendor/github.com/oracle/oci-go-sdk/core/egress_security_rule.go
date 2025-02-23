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
	"github.com/oracle/oci-go-sdk/common"
)

// EgressSecurityRule A rule for allowing outbound IP packets.
type EgressSecurityRule struct {

	// Conceptually, this is the range of IP addresses that a packet originating from the instance
	// can go to.
	// Allowed values:
	//   * IP address range in CIDR notation. For example: `192.168.1.0/24` or `2001:0db8:0123:45::/56`
	//   * The `cidrBlock` value for a Service, if you're
	//     setting up a security list rule for traffic destined for a particular `Service` through
	//     a service gateway. For example: `oci-phx-objectstorage`.
	Destination *string `mandatory:"true" json:"destination"`

	// The transport protocol. Specify either `all` or an IPv4 protocol number as
	// defined in
	// Protocol Numbers (http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).
	// Options are supported only for ICMP ("1"), TCP ("6"), UDP ("17"), and ICMPv6 ("58").
	Protocol *string `mandatory:"true" json:"protocol"`

	// Type of destination for the rule. The default is `CIDR_BLOCK`.
	// Allowed values:
	//   * `CIDR_BLOCK`: If the rule's `destination` is an IP address range in CIDR notation.
	//   * `SERVICE_CIDR_BLOCK`: If the rule's `destination` is the `cidrBlock` value for a
	//     Service (the rule is for traffic destined for a
	//     particular `Service` through a service gateway).
	DestinationType EgressSecurityRuleDestinationTypeEnum `mandatory:"false" json:"destinationType,omitempty"`

	// Optional and valid only for ICMP and ICMPv6. Use to specify a particular ICMP type and code
	// as defined in:
	// * ICMP Parameters (http://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml)
	// * ICMPv6 Parameters (https://www.iana.org/assignments/icmpv6-parameters/icmpv6-parameters.xhtml)
	// If you specify ICMP or ICMPv6 as the protocol but omit this object, then all ICMP types and
	// codes are allowed. If you do provide this object, the type is required and the code is optional.
	// To enable MTU negotiation for ingress internet traffic via IPv4, make sure to allow type 3 ("Destination
	// Unreachable") code 4 ("Fragmentation Needed and Don't Fragment was Set"). If you need to specify
	// multiple codes for a single type, create a separate security list rule for each.
	IcmpOptions *IcmpOptions `mandatory:"false" json:"icmpOptions"`

	// A stateless rule allows traffic in one direction. Remember to add a corresponding
	// stateless rule in the other direction if you need to support bidirectional traffic. For
	// example, if egress traffic allows TCP destination port 80, there should be an ingress
	// rule to allow TCP source port 80. Defaults to false, which means the rule is stateful
	// and a corresponding rule is not necessary for bidirectional traffic.
	IsStateless *bool `mandatory:"false" json:"isStateless"`

	// Optional and valid only for TCP. Use to specify particular destination ports for TCP rules.
	// If you specify TCP as the protocol but omit this object, then all destination ports are allowed.
	TcpOptions *TcpOptions `mandatory:"false" json:"tcpOptions"`

	// Optional and valid only for UDP. Use to specify particular destination ports for UDP rules.
	// If you specify UDP as the protocol but omit this object, then all destination ports are allowed.
	UdpOptions *UdpOptions `mandatory:"false" json:"udpOptions"`
}

func (m EgressSecurityRule) String() string {
	return common.PointerString(m)
}

// EgressSecurityRuleDestinationTypeEnum Enum with underlying type: string
type EgressSecurityRuleDestinationTypeEnum string

// Set of constants representing the allowable values for EgressSecurityRuleDestinationTypeEnum
const (
	EgressSecurityRuleDestinationTypeCidrBlock        EgressSecurityRuleDestinationTypeEnum = "CIDR_BLOCK"
	EgressSecurityRuleDestinationTypeServiceCidrBlock EgressSecurityRuleDestinationTypeEnum = "SERVICE_CIDR_BLOCK"
)

var mappingEgressSecurityRuleDestinationType = map[string]EgressSecurityRuleDestinationTypeEnum{
	"CIDR_BLOCK":         EgressSecurityRuleDestinationTypeCidrBlock,
	"SERVICE_CIDR_BLOCK": EgressSecurityRuleDestinationTypeServiceCidrBlock,
}

// GetEgressSecurityRuleDestinationTypeEnumValues Enumerates the set of values for EgressSecurityRuleDestinationTypeEnum
func GetEgressSecurityRuleDestinationTypeEnumValues() []EgressSecurityRuleDestinationTypeEnum {
	values := make([]EgressSecurityRuleDestinationTypeEnum, 0)
	for _, v := range mappingEgressSecurityRuleDestinationType {
		values = append(values, v)
	}
	return values
}
