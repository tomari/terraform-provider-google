---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This file is automatically generated by Magic Modules and manual
#     changes will be clobbered when the file is regenerated.
#
#     Please read more about how to change this file in
#     .github/CONTRIBUTING.md.
#
# ----------------------------------------------------------------------------
subcategory: "App Engine"
layout: "google"
page_title: "Google: google_app_engine_service_network_settings"
sidebar_current: "docs-google-app-engine-service-network-settings"
description: |-
  A NetworkSettings resource is a container for ingress settings for a version or service.
---

# google\_app\_engine\_service\_network\_settings

A NetworkSettings resource is a container for ingress settings for a version or service.


To get more information about ServiceNetworkSettings, see:

* [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=app_engine_service_network_settings&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - App Engine Service Network Settings


```hcl
resource "google_storage_bucket" "bucket" {
	name     = "appengine-static-content"
  location = "US"
}

resource "google_storage_bucket_object" "object" {
	name   = "hello-world.zip"
	bucket = google_storage_bucket.bucket.name
	source = "./test-fixtures/appengine/hello-world.zip"
}

resource "google_app_engine_standard_app_version" "liveapp_v1" {
  version_id = "v1"
  service = "liveapp"
  delete_service_on_destroy = true

  runtime = "nodejs10"
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.name}"
    }  
  }
  env_variables = {
    port = "8080"
  }
}

resource "google_app_engine_service_network_settings" "liveapp" {
  service = google_app_engine_standard_app_version.liveapp_v1.service
  network_settings {
    ingress_traffic_allowed = "INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY"
  }
}
```

## Argument Reference

The following arguments are supported:


* `service` -
  (Required)
  The name of the service these settings apply to.

* `network_settings` -
  (Required)
  Ingress settings for this service. Will apply to all versions.
  Structure is [documented below](#nested_network_settings).


<a name="nested_network_settings"></a>The `network_settings` block supports:

* `ingress_traffic_allowed` -
  (Optional)
  The ingress settings for version or service.
  Default value is `INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED`.
  Possible values are `INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED`, `INGRESS_TRAFFIC_ALLOWED_ALL`, `INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY`, and `INGRESS_TRAFFIC_ALLOWED_INTERNAL_AND_LB`.

- - -


* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `apps/{{project}}/services/{{service}}`


## Timeouts

This resource provides the following
[Timeouts](/docs/configuration/resources.html#timeouts) configuration options:

- `create` - Default is 4 minutes.
- `update` - Default is 4 minutes.
- `delete` - Default is 4 minutes.

## Import


ServiceNetworkSettings can be imported using any of these accepted formats:

```
$ terraform import google_app_engine_service_network_settings.default apps/{{project}}/services/{{service}}
$ terraform import google_app_engine_service_network_settings.default {{project}}/{{service}}
$ terraform import google_app_engine_service_network_settings.default {{service}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/guides/provider_reference.html#user_project_override).
