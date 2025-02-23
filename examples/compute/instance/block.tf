// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

resource "oci_core_volume" "TFBlock" {
  count               = "${var.NumInstances * var.NumIscsiVolumesPerInstance}"
  availability_domain = "${data.oci_identity_availability_domain.ad.name}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "TFBlock${count.index}"
  size_in_gbs         = "${var.DBSize}"
}

resource "oci_core_volume_attachment" "TFBlockAttach" {
  count           = "${var.NumInstances * var.NumIscsiVolumesPerInstance}"
  attachment_type = "iscsi"
  instance_id     = "${oci_core_instance.TFInstance.*.id[floor(count.index / var.NumIscsiVolumesPerInstance)]}"
  volume_id       = "${oci_core_volume.TFBlock.*.id[count.index]}"
  device          = "${count.index == 0 ? var.volume_attachment_device : ""}"

  # Set this to enable CHAP authentication for an ISCSI volume attachment. The oci_core_volume_attachment resource will
  # contain the CHAP authentication details via the "chap_secret" and "chap_username" attributes.
  use_chap = true

  # Set this to attach the volume as read-only.
  #is_read_only = true
}

resource "oci_core_volume" "TFBlockParavirtualized" {
  count               = "${var.NumInstances * var.NumParavirtualizedVolumesPerInstance}"
  availability_domain = "${data.oci_identity_availability_domain.ad.name}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "TFBlockParavirtualized${count.index}"
  size_in_gbs         = "${var.DBSize}"
}

resource "oci_core_volume_attachment" "TFBlockAttachParavirtualized" {
  count           = "${var.NumInstances * var.NumParavirtualizedVolumesPerInstance}"
  attachment_type = "paravirtualized"
  instance_id     = "${oci_core_instance.TFInstance.*.id[floor(count.index / var.NumParavirtualizedVolumesPerInstance)]}"
  volume_id       = "${oci_core_volume.TFBlockParavirtualized.*.id[count.index]}"

  # Set this to attach the volume as read-only.
  #is_read_only = true
}

// Backup Policy assignment
resource "oci_core_volume_backup_policy_assignment" "policy" {
  count     = 2
  asset_id  = "${oci_core_instance.TFInstance.*.boot_volume_id[count.index]}"
  policy_id = "${data.oci_core_volume_backup_policies.test_predefined_volume_backup_policies.volume_backup_policies.0.id}"
}
