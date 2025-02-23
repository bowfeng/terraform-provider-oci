// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

/* This example demonstrates object store object management. It uses Terraforms built-in `file` function to upload a file.
 * It also demonstrates the use of `local-exec` to download the object content using a preauthenticated request.
 *
 * WARNING: This should only be used with small files. The file helper does stringification so large files
 * may cause terraform to slow, become unresponsive or exceed allowed memory usage.
 */

resource "oci_objectstorage_object" "object1" {
  namespace        = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket           = "${oci_objectstorage_bucket.bucket1.name}"
  object           = "index.html"
  content_language = "en-US"
  content_type     = "text/html"
  content          = "${file("index.html")}"
}

resource "oci_objectstorage_object" "source_object" {
  namespace        = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket           = "${oci_objectstorage_bucket.bucket1.name}"
  object           = "same_index.html"
  content_language = "en-US"
  content_type     = "text/html"
  source           = "index.html"
}

resource "oci_objectstorage_object" "source_uri_object" {
  namespace        = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket           = "${oci_objectstorage_bucket.bucket1.name}"
  object           = "copy_index.html"
  content_language = "en-US"
  content_type     = "text/html"

  source_uri_details {
    region    = "${local.source_region}"
    namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
    bucket    = "${oci_objectstorage_bucket.bucket1.name}"
    object    = "index.html"
  }
}

data "oci_objectstorage_object_head" "object_head1" {
  namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket    = "${oci_objectstorage_bucket.bucket1.name}"
  object    = "${oci_objectstorage_object.object1.object}"
}

data "oci_objectstorage_object_head" "source_object_head" {
  namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket    = "${oci_objectstorage_bucket.bucket1.name}"
  object    = "${oci_objectstorage_object.source_object.object}"
}

data "oci_objectstorage_objects" "objects1" {
  namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket    = "${oci_objectstorage_bucket.bucket1.name}"
}

data "oci_objectstorage_object" "object" {
  namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket    = "${oci_objectstorage_bucket.bucket1.name}"
  object    = "index.html"
}

output object_data {
  value = <<EOF

  content = ${data.oci_objectstorage_object.object.content}
  content-length = ${data.oci_objectstorage_object.object.content_length}
  content-type = ${data.oci_objectstorage_object.object.content_type}
EOF
}

output object-head-data {
  value = <<EOF

  object = ${data.oci_objectstorage_object_head.object_head1.object}
  content-length = ${data.oci_objectstorage_object_head.object_head1.content_length}
  content-type = ${data.oci_objectstorage_object_head.object_head1.content_type}
EOF
}

output object-source-head-data {
  value = <<EOF

  object = ${data.oci_objectstorage_object_head.source_object_head.object}
  content-length = ${data.oci_objectstorage_object_head.source_object_head.content-length}
  content-type = ${data.oci_objectstorage_object_head.source_object_head.content-type}
EOF
}

output objects {
  value = "${data.oci_objectstorage_objects.objects1.objects}"
}

# example to download object content to local file using object PAR with local-exec

resource "null_resource" "download_object_content" {
  # using command
  provisioner "local-exec" {
    command = "curl -o object-content https://objectstorage.${var.region}.oraclecloud.com${oci_objectstorage_preauthrequest.object_par.access_uri}"
  }

  # using script
  provisioner "local-exec" {
    command = "sh get-content.sh https://objectstorage.${var.region}.oraclecloud.com${oci_objectstorage_preauthrequest.object_par.access_uri} object-content-using-script"
  }

  provisioner "local-exec" {
    when    = "destroy"
    command = "rm -rf object-content content/"
  }
}
