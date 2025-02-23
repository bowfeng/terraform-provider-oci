// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

resource "oci_core_instance" "TFInstance" {
  availability_domain = "${data.oci_identity_availability_domain.ad.name}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "TFInstanceForInstancePool"
  shape               = "${var.instance_shape}"

  create_vnic_details {
    subnet_id        = "${oci_core_subnet.ExampleSubnet.id}"
    display_name     = "primaryvnic"
    assign_public_ip = true
    hostname_label   = "tfexampleinstance"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.region]}"
  }

  timeouts {
    create = "60m"
  }
}

resource "oci_core_image" "TFCustomImage" {
  compartment_id = "${var.compartment_ocid}"
  instance_id    = "${oci_core_instance.TFInstance.id}"
  launch_mode    = "NATIVE"

  timeouts {
    create = "30m"
  }
}

resource "oci_core_instance_configuration" "TFInstanceConfiguration" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFExampleInstanceConfiguration"

  instance_details {
    instance_type = "compute"

    launch_details {
      compartment_id = "${var.compartment_ocid}"
      ipxe_script    = "ipxeScript"
      shape          = "${var.instance_shape}"
      display_name   = "TFExampleInstanceConfigurationLaunchDetails"

      create_vnic_details {
        assign_public_ip       = true
        display_name           = "TFExampleInstanceConfigurationVNIC"
        skip_source_dest_check = false
      }

      extended_metadata = {
        some_string   = "stringA"
        nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
      }

      source_details {
        source_type = "image"
        image_id    = "${oci_core_image.TFCustomImage.id}"
      }
    }
  }
}

resource "oci_core_instance_pool" "TFInstancePool" {
  compartment_id            = "${var.compartment_ocid}"
  instance_configuration_id = "${oci_core_instance_configuration.TFInstanceConfiguration.id}"
  size                      = 2
  state                     = "RUNNING"
  display_name              = "TFInstancePool"

  placement_configurations {
    availability_domain = "${data.oci_identity_availability_domain.ad.name}"
    primary_subnet_id   = "${oci_core_subnet.ExampleSubnet.id}"
  }

  load_balancers = {
    backend_set_name = "${oci_load_balancer_backend_set.test_backend_set.name}"
    load_balancer_id = "${oci_load_balancer.test_load_balancer.id}"
    port             = 80
    vnic_selection   = "primaryvnic"
  }
}
