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

// LaunchOptions Options for tuning compatibility and performance of VM shapes.
type LaunchOptions struct {

	// Emulation type for volume.
	// * `ISCSI` - ISCSI attached block storage device. This is the default for Boot Volumes and Remote Block
	// Storage volumes on Oracle provided images.
	// * `SCSI` - Emulated SCSI disk.
	// * `IDE` - Emulated IDE disk.
	// * `VFIO` - Direct attached Virtual Function storage.  This is the default option for Local data
	// volumes on Oracle provided images.
	// * `PARAVIRTUALIZED` - Paravirtualized disk.
	BootVolumeType LaunchOptionsBootVolumeTypeEnum `mandatory:"true" json:"bootVolumeType"`

	// Firmware used to boot VM.  Select the option that matches your operating system.
	// * `BIOS` - Boot VM using BIOS style firmware.  This is compatible with both 32 bit and 64 bit operating
	// systems that boot using MBR style bootloaders.
	// * `UEFI_64` - Boot VM using UEFI style firmware compatible with 64 bit operating systems.  This is the
	// default for Oracle provided images.
	Firmware LaunchOptionsFirmwareEnum `mandatory:"true" json:"firmware"`

	// Emulation type for NIC.
	// * `E1000` - Emulated Gigabit ethernet controller.  Compatible with Linux e1000 network driver.
	// * `VFIO` - Direct attached Virtual Function network controller.  Default for Oracle provided images.
	// * `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using virtio drivers.
	NetworkType LaunchOptionsNetworkTypeEnum `mandatory:"true" json:"networkType"`

	// Emulation type for volume.
	// * `ISCSI` - ISCSI attached block storage device. This is the default for Boot Volumes and Remote Block
	// Storage volumes on Oracle provided images.
	// * `SCSI` - Emulated SCSI disk.
	// * `IDE` - Emulated IDE disk.
	// * `VFIO` - Direct attached Virtual Function storage.  This is the default option for Local data
	// volumes on Oracle provided images.
	// * `PARAVIRTUALIZED` - Paravirtualized disk.
	RemoteDataVolumeType LaunchOptionsRemoteDataVolumeTypeEnum `mandatory:"true" json:"remoteDataVolumeType"`

	// Whether to enable in-transit encryption for the boot volume's paravirtualized attachment. The default value is false.
	IsPvEncryptionInTransitEnabled *bool `mandatory:"false" json:"isPvEncryptionInTransitEnabled"`

	// Whether to enable consistent volume naming feature. Defaults to false.
	IsConsistentVolumeNamingEnabled *bool `mandatory:"false" json:"isConsistentVolumeNamingEnabled"`
}

func (m LaunchOptions) String() string {
	return common.PointerString(m)
}

// LaunchOptionsBootVolumeTypeEnum Enum with underlying type: string
type LaunchOptionsBootVolumeTypeEnum string

// Set of constants representing the allowable values for LaunchOptionsBootVolumeTypeEnum
const (
	LaunchOptionsBootVolumeTypeIscsi           LaunchOptionsBootVolumeTypeEnum = "ISCSI"
	LaunchOptionsBootVolumeTypeScsi            LaunchOptionsBootVolumeTypeEnum = "SCSI"
	LaunchOptionsBootVolumeTypeIde             LaunchOptionsBootVolumeTypeEnum = "IDE"
	LaunchOptionsBootVolumeTypeVfio            LaunchOptionsBootVolumeTypeEnum = "VFIO"
	LaunchOptionsBootVolumeTypeParavirtualized LaunchOptionsBootVolumeTypeEnum = "PARAVIRTUALIZED"
)

var mappingLaunchOptionsBootVolumeType = map[string]LaunchOptionsBootVolumeTypeEnum{
	"ISCSI":           LaunchOptionsBootVolumeTypeIscsi,
	"SCSI":            LaunchOptionsBootVolumeTypeScsi,
	"IDE":             LaunchOptionsBootVolumeTypeIde,
	"VFIO":            LaunchOptionsBootVolumeTypeVfio,
	"PARAVIRTUALIZED": LaunchOptionsBootVolumeTypeParavirtualized,
}

// GetLaunchOptionsBootVolumeTypeEnumValues Enumerates the set of values for LaunchOptionsBootVolumeTypeEnum
func GetLaunchOptionsBootVolumeTypeEnumValues() []LaunchOptionsBootVolumeTypeEnum {
	values := make([]LaunchOptionsBootVolumeTypeEnum, 0)
	for _, v := range mappingLaunchOptionsBootVolumeType {
		values = append(values, v)
	}
	return values
}

// LaunchOptionsFirmwareEnum Enum with underlying type: string
type LaunchOptionsFirmwareEnum string

// Set of constants representing the allowable values for LaunchOptionsFirmwareEnum
const (
	LaunchOptionsFirmwareBios   LaunchOptionsFirmwareEnum = "BIOS"
	LaunchOptionsFirmwareUefi64 LaunchOptionsFirmwareEnum = "UEFI_64"
)

var mappingLaunchOptionsFirmware = map[string]LaunchOptionsFirmwareEnum{
	"BIOS":    LaunchOptionsFirmwareBios,
	"UEFI_64": LaunchOptionsFirmwareUefi64,
}

// GetLaunchOptionsFirmwareEnumValues Enumerates the set of values for LaunchOptionsFirmwareEnum
func GetLaunchOptionsFirmwareEnumValues() []LaunchOptionsFirmwareEnum {
	values := make([]LaunchOptionsFirmwareEnum, 0)
	for _, v := range mappingLaunchOptionsFirmware {
		values = append(values, v)
	}
	return values
}

// LaunchOptionsNetworkTypeEnum Enum with underlying type: string
type LaunchOptionsNetworkTypeEnum string

// Set of constants representing the allowable values for LaunchOptionsNetworkTypeEnum
const (
	LaunchOptionsNetworkTypeE1000           LaunchOptionsNetworkTypeEnum = "E1000"
	LaunchOptionsNetworkTypeVfio            LaunchOptionsNetworkTypeEnum = "VFIO"
	LaunchOptionsNetworkTypeParavirtualized LaunchOptionsNetworkTypeEnum = "PARAVIRTUALIZED"
)

var mappingLaunchOptionsNetworkType = map[string]LaunchOptionsNetworkTypeEnum{
	"E1000":           LaunchOptionsNetworkTypeE1000,
	"VFIO":            LaunchOptionsNetworkTypeVfio,
	"PARAVIRTUALIZED": LaunchOptionsNetworkTypeParavirtualized,
}

// GetLaunchOptionsNetworkTypeEnumValues Enumerates the set of values for LaunchOptionsNetworkTypeEnum
func GetLaunchOptionsNetworkTypeEnumValues() []LaunchOptionsNetworkTypeEnum {
	values := make([]LaunchOptionsNetworkTypeEnum, 0)
	for _, v := range mappingLaunchOptionsNetworkType {
		values = append(values, v)
	}
	return values
}

// LaunchOptionsRemoteDataVolumeTypeEnum Enum with underlying type: string
type LaunchOptionsRemoteDataVolumeTypeEnum string

// Set of constants representing the allowable values for LaunchOptionsRemoteDataVolumeTypeEnum
const (
	LaunchOptionsRemoteDataVolumeTypeIscsi           LaunchOptionsRemoteDataVolumeTypeEnum = "ISCSI"
	LaunchOptionsRemoteDataVolumeTypeScsi            LaunchOptionsRemoteDataVolumeTypeEnum = "SCSI"
	LaunchOptionsRemoteDataVolumeTypeIde             LaunchOptionsRemoteDataVolumeTypeEnum = "IDE"
	LaunchOptionsRemoteDataVolumeTypeVfio            LaunchOptionsRemoteDataVolumeTypeEnum = "VFIO"
	LaunchOptionsRemoteDataVolumeTypeParavirtualized LaunchOptionsRemoteDataVolumeTypeEnum = "PARAVIRTUALIZED"
)

var mappingLaunchOptionsRemoteDataVolumeType = map[string]LaunchOptionsRemoteDataVolumeTypeEnum{
	"ISCSI":           LaunchOptionsRemoteDataVolumeTypeIscsi,
	"SCSI":            LaunchOptionsRemoteDataVolumeTypeScsi,
	"IDE":             LaunchOptionsRemoteDataVolumeTypeIde,
	"VFIO":            LaunchOptionsRemoteDataVolumeTypeVfio,
	"PARAVIRTUALIZED": LaunchOptionsRemoteDataVolumeTypeParavirtualized,
}

// GetLaunchOptionsRemoteDataVolumeTypeEnumValues Enumerates the set of values for LaunchOptionsRemoteDataVolumeTypeEnum
func GetLaunchOptionsRemoteDataVolumeTypeEnumValues() []LaunchOptionsRemoteDataVolumeTypeEnum {
	values := make([]LaunchOptionsRemoteDataVolumeTypeEnum, 0)
	for _, v := range mappingLaunchOptionsRemoteDataVolumeType {
		values = append(values, v)
	}
	return values
}
