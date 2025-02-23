---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume_backup_policies"
sidebar_current: "docs-oci-datasource-core-volume_backup_policies"
description: |-
  Provides the list of Volume Backup Policies in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_volume_backup_policies
This data source provides the list of Volume Backup Policies in Oracle Cloud Infrastructure Core service.

Lists all volume backup policies available to the caller.

## Example Usage

```hcl
data "oci_core_volume_backup_policies" "test_volume_backup_policies" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `volume_backup_policies` - The list of volume_backup_policies.

### VolumeBackupPolicy Reference

The following attributes are exported:

* `display_name` - A user-friendly name for the volume backup policy. Does not have to be unique and it's changeable. Avoid entering confidential information. 
* `id` - The OCID of the volume backup policy.
* `schedules` - The collection of schedules that this policy will apply.
	* `backup_type` - The type of backup to create.
	* `offset_seconds` - The number of seconds that the backup time should be shifted from the default interval boundaries specified by the period. Backup time = Frequency start time + Offset.
	* `period` - How often the backup should occur.
	* `retention_seconds` - How long, in seconds, backups created by this schedule should be kept until being automatically deleted.
* `time_created` - The date and time the volume backup policy was created. Format defined by RFC3339. 

