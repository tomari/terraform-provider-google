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
subcategory: "Cloud Build"
layout: "google"
page_title: "Google: google_cloudbuild_trigger"
sidebar_current: "docs-google-cloudbuild-trigger"
description: |-
  Configuration for an automated build in response to source repository changes.
---

# google\_cloudbuild\_trigger

Configuration for an automated build in response to source repository changes.


To get more information about Trigger, see:

* [API documentation](https://cloud.google.com/cloud-build/docs/api/reference/rest/v1/projects.triggers)
* How-to Guides
    * [Automating builds using build triggers](https://cloud.google.com/cloud-build/docs/running-builds/automate-builds)

~> **Note:** You can retrieve the email of the Cloud Build Service Account used in jobs by using the `google_project_service_identity` resource.

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=cloudbuild_trigger_filename&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Cloudbuild Trigger Filename


```hcl
resource "google_cloudbuild_trigger" "filename-trigger" {
  trigger_template {
    branch_name = "master"
    repo_name   = "my-repo"
  }

  substitutions = {
    _FOO = "bar"
    _BAZ = "qux"
  }

  filename = "cloudbuild.yaml"
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=cloudbuild_trigger_build&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Cloudbuild Trigger Build


```hcl
resource "google_cloudbuild_trigger" "build-trigger" {
  trigger_template {
    branch_name = "master"
    repo_name   = "my-repo"
  }

  build {
    step {
      name = "gcr.io/cloud-builders/gsutil"
      args = ["cp", "gs://mybucket/remotefile.zip", "localfile.zip"]
      timeout = "120s"
      secret_env = ["MY_SECRET"]
    }

    source {
      storage_source {
        bucket = "mybucket"
        object = "source_code.tar.gz"
      }
    }
    tags = ["build", "newFeature"]
    substitutions = {
      _FOO = "bar"
      _BAZ = "qux"
    }
    queue_ttl = "20s"
    logs_bucket = "gs://mybucket/logs"
    secret {
      kms_key_name = "projects/myProject/locations/global/keyRings/keyring-name/cryptoKeys/key-name"
      secret_env = {
        PASSWORD = "ZW5jcnlwdGVkLXBhc3N3b3JkCg=="
      }
    }
    available_secrets {
      secret_manager {
        env          = "MY_SECRET"
        version_name = "projects/myProject/secrets/mySecret/versions/latest"
      }
    }
    artifacts {
      images = ["gcr.io/$PROJECT_ID/$REPO_NAME:$COMMIT_SHA"]
      objects {
        location = "gs://bucket/path/to/somewhere/"
        paths = ["path"]
      }
    }
    options {
      source_provenance_hash = ["MD5"]
      requested_verify_option = "VERIFIED"
      machine_type = "N1_HIGHCPU_8"
      disk_size_gb = 100
      substitution_option = "ALLOW_LOOSE"
      dynamic_substitutions = true
      log_streaming_option = "STREAM_OFF"
      worker_pool = "pool"
      logging = "LEGACY"
      env = ["ekey = evalue"]
      secret_env = ["secretenv = svalue"]
      volumes {
        name = "v1"
        path = "v1"
      }
    }
  }
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=cloudbuild_trigger_service_account&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Cloudbuild Trigger Service Account


```hcl
data "google_project" "project" {}

resource "google_cloudbuild_trigger" "service-account-trigger" {
  trigger_template {
    branch_name = "master"
    repo_name   = "my-repo"
  }

  service_account = google_service_account.cloudbuild_service_account.id
  filename        = "cloudbuild.yaml"
  depends_on = [
    google_project_iam_member.act_as,
    google_project_iam_member.logs_writer
  ]
}

resource "google_service_account" "cloudbuild_service_account" {
  account_id = "my-service-account"
}

resource "google_project_iam_member" "act_as" {
  project = data.google_project.project.project_id
  role    = "roles/iam.serviceAccountUser"
  member  = "serviceAccount:${google_service_account.cloudbuild_service_account.email}"
}

resource "google_project_iam_member" "logs_writer" {
  project = data.google_project.project.project_id
  role    = "roles/logging.logWriter"
  member  = "serviceAccount:${google_service_account.cloudbuild_service_account.email}"
}
```

## Argument Reference

The following arguments are supported:



- - -


* `name` -
  (Optional)
  Name of the trigger. Must be unique within the project.

* `description` -
  (Optional)
  Human-readable description of the trigger.

* `tags` -
  (Optional)
  Tags for annotation of a BuildTrigger

* `disabled` -
  (Optional)
  Whether the trigger is disabled or not. If true, the trigger will never result in a build.

* `substitutions` -
  (Optional)
  Substitutions data for Build resource.

* `service_account` -
  (Optional)
  The service account used for all user-controlled operations including
  triggers.patch, triggers.run, builds.create, and builds.cancel.
  If no service account is set, then the standard Cloud Build service account
  ([PROJECT_NUM]@system.gserviceaccount.com) will be used instead.
  Format: projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT_ID_OR_EMAIL}

* `filename` -
  (Optional)
  Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.

* `ignored_files` -
  (Optional)
  ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
  extended with support for `**`.
  If ignoredFiles and changed files are both empty, then they are not
  used to determine whether or not to trigger a build.
  If ignoredFiles is not empty, then we ignore any files that match any
  of the ignored_file globs. If the change has no files that are outside
  of the ignoredFiles globs, then we do not trigger a build.

* `included_files` -
  (Optional)
  ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
  extended with support for `**`.
  If any of the files altered in the commit pass the ignoredFiles filter
  and includedFiles is empty, then as far as this filter is concerned, we
  should trigger the build.
  If any of the files altered in the commit pass the ignoredFiles filter
  and includedFiles is not empty, then we make sure that at least one of
  those files matches a includedFiles glob. If not, then we do not trigger
  a build.

* `trigger_template` -
  (Optional)
  Template describing the types of source changes to trigger a build.
  Branch and tag names in trigger templates are interpreted as regular
  expressions. Any branch or tag change that matches that regular
  expression will trigger a build.
  One of `trigger_template`, `github`, `pubsub_config` or `webhook_config` must be provided.
  Structure is [documented below](#nested_trigger_template).

* `github` -
  (Optional)
  Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
  One of `trigger_template`, `github`, `pubsub_config` or `webhook_config` must be provided.
  Structure is [documented below](#nested_github).

* `pubsub_config` -
  (Optional)
  PubsubConfig describes the configuration of a trigger that creates 
  a build whenever a Pub/Sub message is published.
  One of `trigger_template`, `github`, `pubsub_config` or `webhook_config` must be provided.
  Structure is [documented below](#nested_pubsub_config).

* `webhook_config` -
  (Optional)
  WebhookConfig describes the configuration of a trigger that creates 
  a build whenever a webhook is sent to a trigger's webhook URL.
  One of `trigger_template`, `github`, `pubsub_config` or `webhook_config` must be provided.
  Structure is [documented below](#nested_webhook_config).

* `build` -
  (Optional)
  Contents of the build template. Either a filename or build template must be provided.
  Structure is [documented below](#nested_build).

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.


<a name="nested_trigger_template"></a>The `trigger_template` block supports:

* `project_id` -
  (Optional)
  ID of the project that owns the Cloud Source Repository. If
  omitted, the project ID requesting the build is assumed.

* `repo_name` -
  (Optional)
  Name of the Cloud Source Repository. If omitted, the name "default" is assumed.

* `dir` -
  (Optional)
  Directory, relative to the source root, in which to run the build.
  This must be a relative path. If a step's dir is specified and
  is an absolute path, this value is ignored for that step's
  execution.

* `invert_regex` -
  (Optional)
  Only trigger a build if the revision regex does NOT match the revision regex.

* `branch_name` -
  (Optional)
  Name of the branch to build. Exactly one a of branch name, tag, or commit SHA must be provided.
  This field is a regular expression.

* `tag_name` -
  (Optional)
  Name of the tag to build. Exactly one of a branch name, tag, or commit SHA must be provided.
  This field is a regular expression.

* `commit_sha` -
  (Optional)
  Explicit commit SHA to build. Exactly one of a branch name, tag, or commit SHA must be provided.

<a name="nested_github"></a>The `github` block supports:

* `owner` -
  (Optional)
  Owner of the repository. For example: The owner for
  https://github.com/googlecloudplatform/cloud-builders is "googlecloudplatform".

* `name` -
  (Optional)
  Name of the repository. For example: The name for
  https://github.com/googlecloudplatform/cloud-builders is "cloud-builders".

* `pull_request` -
  (Optional)
  filter to match changes in pull requests. Specify only one of `pull_request` or `push`.
  Structure is [documented below](#nested_pull_request).

* `push` -
  (Optional)
  filter to match changes in refs, like branches or tags. Specify only one of `pull_request` or `push`.
  Structure is [documented below](#nested_push).


<a name="nested_pull_request"></a>The `pull_request` block supports:

* `branch` -
  (Required)
  Regex of branches to match.

* `comment_control` -
  (Optional)
  Whether to block builds on a "/gcbrun" comment from a repository owner or collaborator.
  Possible values are `COMMENTS_DISABLED`, `COMMENTS_ENABLED`, and `COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY`.

* `invert_regex` -
  (Optional)
  If true, branches that do NOT match the git_ref will trigger a build.

<a name="nested_push"></a>The `push` block supports:

* `invert_regex` -
  (Optional)
  When true, only trigger a build if the revision regex does NOT match the git_ref regex.

* `branch` -
  (Optional)
  Regex of branches to match.  Specify only one of branch or tag.

* `tag` -
  (Optional)
  Regex of tags to match.  Specify only one of branch or tag.

<a name="nested_pubsub_config"></a>The `pubsub_config` block supports:

* `subscription` -
  Output only. Name of the subscription.

* `topic` -
  (Required)
  The name of the topic from which this subscription is receiving messages.

* `service_account_email` -
  (Optional)
  Service account that will make the push request.

* `state` -
  Potential issues with the underlying Pub/Sub subscription configuration.
  Only populated on get requests.

<a name="nested_webhook_config"></a>The `webhook_config` block supports:

* `secret` -
  (Required)
  Resource name for the secret required as a URL parameter.

* `state` -
  Potential issues with the underlying Pub/Sub subscription configuration.
  Only populated on get requests.

<a name="nested_build"></a>The `build` block supports:

* `source` -
  (Optional)
  The location of the source files to build.
  One of `storageSource` or `repoSource` must be provided.
  Structure is [documented below](#nested_source).

* `tags` -
  (Optional)
  Tags for annotation of a Build. These are not docker tags.

* `images` -
  (Optional)
  A list of images to be pushed upon the successful completion of all build steps.
  The images are pushed using the builder service account's credentials.
  The digests of the pushed images will be stored in the Build resource's results field.
  If any of the images fail to be pushed, the build status is marked FAILURE.

* `substitutions` -
  (Optional)
  Substitutions data for Build resource.

* `queue_ttl` -
  (Optional)
  TTL in queue for this build. If provided and the build is enqueued longer than this value, 
  the build will expire and the build status will be EXPIRED.
  The TTL starts ticking from createTime.
  A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".

* `logs_bucket` -
  (Optional)
  Google Cloud Storage bucket where logs should be written. 
  Logs file names will be of the format ${logsBucket}/log-${build_id}.txt.

* `timeout` -
  (Optional)
  Amount of time that this build should be allowed to run, to second granularity.
  If this amount of time elapses, work on the build will cease and the build status will be TIMEOUT.
  This timeout must be equal to or greater than the sum of the timeouts for build steps within the build.
  The expected format is the number of seconds followed by s.
  Default time is ten minutes (600s).

* `secret` -
  (Optional)
  Secrets to decrypt using Cloud Key Management Service.
  Structure is [documented below](#nested_secret).

* `available_secrets` -
  (Optional)
  Secrets and secret environment variables.
  Structure is [documented below](#nested_available_secrets).

* `step` -
  (Required)
  The operations to be performed on the workspace.
  Structure is [documented below](#nested_step).

* `artifacts` -
  (Optional)
  Artifacts produced by the build that should be uploaded upon successful completion of all build steps.
  Structure is [documented below](#nested_artifacts).

* `options` -
  (Optional)
  Special options for this build.
  Structure is [documented below](#nested_options).


<a name="nested_source"></a>The `source` block supports:

* `storage_source` -
  (Optional)
  Location of the source in an archive file in Google Cloud Storage.
  Structure is [documented below](#nested_storage_source).

* `repo_source` -
  (Optional)
  Location of the source in a Google Cloud Source Repository.
  Structure is [documented below](#nested_repo_source).


<a name="nested_storage_source"></a>The `storage_source` block supports:

* `bucket` -
  (Required)
  Google Cloud Storage bucket containing the source.

* `object` -
  (Required)
  Google Cloud Storage object containing the source.
  This object must be a gzipped archive file (.tar.gz) containing source to build.

* `generation` -
  (Optional)
  Google Cloud Storage generation for the object. 
  If the generation is omitted, the latest generation will be used

<a name="nested_repo_source"></a>The `repo_source` block supports:

* `project_id` -
  (Optional)
  ID of the project that owns the Cloud Source Repository. 
  If omitted, the project ID requesting the build is assumed.

* `repo_name` -
  (Required)
  Name of the Cloud Source Repository.

* `dir` -
  (Optional)
  Directory, relative to the source root, in which to run the build.
  This must be a relative path. If a step's dir is specified and is an absolute path, 
  this value is ignored for that step's execution.

* `invert_regex` -
  (Optional)
  Only trigger a build if the revision regex does NOT match the revision regex.

* `substitutions` -
  (Optional)
  Substitutions to use in a triggered build. Should only be used with triggers.run

* `branch_name` -
  (Optional)
  Regex matching branches to build. Exactly one a of branch name, tag, or commit SHA must be provided.
  The syntax of the regular expressions accepted is the syntax accepted by RE2 and 
  described at https://github.com/google/re2/wiki/Syntax

* `tag_name` -
  (Optional)
  Regex matching tags to build. Exactly one a of branch name, tag, or commit SHA must be provided.
  The syntax of the regular expressions accepted is the syntax accepted by RE2 and 
  described at https://github.com/google/re2/wiki/Syntax

* `commit_sha` -
  (Optional)
  Explicit commit SHA to build. Exactly one a of branch name, tag, or commit SHA must be provided.

<a name="nested_secret"></a>The `secret` block supports:

* `kms_key_name` -
  (Required)
  Cloud KMS key name to use to decrypt these envs.

* `secret_env` -
  (Optional)
  Map of environment variable name to its encrypted value.
  Secret environment variables must be unique across all of a build's secrets, 
  and must be used by at least one build step. Values can be at most 64 KB in size. 
  There can be at most 100 secret values across all of a build's secrets.

<a name="nested_available_secrets"></a>The `available_secrets` block supports:

* `secret_manager` -
  (Required)
  Pairs a secret environment variable with a SecretVersion in Secret Manager.
  Structure is [documented below](#nested_secret_manager).


<a name="nested_secret_manager"></a>The `secret_manager` block supports:

* `version_name` -
  (Required)
  Resource name of the SecretVersion. In format: projects/*/secrets/*/versions/*

* `env` -
  (Required)
  Environment variable name to associate with the secret. Secret environment
  variables must be unique across all of a build's secrets, and must be used
  by at least one build step.

<a name="nested_step"></a>The `step` block supports:

* `name` -
  (Required)
  The name of the container image that will run this particular build step.
  If the image is available in the host's Docker daemon's cache, it will be
  run directly. If not, the host will attempt to pull the image first, using
  the builder service account's credentials if necessary.
  The Docker daemon's cache will already have the latest versions of all of
  the officially supported build steps (see https://github.com/GoogleCloudPlatform/cloud-builders 
  for images and examples).
  The Docker daemon will also have cached many of the layers for some popular
  images, like "ubuntu", "debian", but they will be refreshed at the time
  you attempt to use them.
  If you built an image in a previous build step, it will be stored in the
  host's Docker daemon's cache and is available to use as the name for a
  later build step.

* `args` -
  (Optional)
  A list of arguments that will be presented to the step when it is started.
  If the image used to run the step's container has an entrypoint, the args
  are used as arguments to that entrypoint. If the image does not define an
  entrypoint, the first element in args is used as the entrypoint, and the
  remainder will be used as arguments.

* `env` -
  (Optional)
  A list of environment variable definitions to be used when
  running a step.
  The elements are of the form "KEY=VALUE" for the environment variable
  "KEY" being given the value "VALUE".

* `id` -
  (Optional)
  Unique identifier for this build step, used in `wait_for` to
  reference this build step as a dependency.

* `entrypoint` -
  (Optional)
  Entrypoint to be used instead of the build step image's
  default entrypoint.
  If unset, the image's default entrypoint is used

* `dir` -
  (Optional)
  Working directory to use when running this step's container.
  If this value is a relative path, it is relative to the build's working
  directory. If this value is absolute, it may be outside the build's working
  directory, in which case the contents of the path may not be persisted
  across build step executions, unless a `volume` for that path is specified.
  If the build specifies a `RepoSource` with `dir` and a step with a
  `dir`,
  which specifies an absolute path, the `RepoSource` `dir` is ignored
  for the step's execution.

* `secret_env` -
  (Optional)
  A list of environment variables which are encrypted using
  a Cloud Key
  Management Service crypto key. These values must be specified in
  the build's `Secret`.

* `timeout` -
  (Optional)
  Time limit for executing this build step. If not defined,
  the step has no
  time limit and will be allowed to continue to run until either it
  completes or the build itself times out.

* `timing` -
  (Optional)
  Output only. Stores timing information for executing this
  build step.

* `volumes` -
  (Optional)
  List of volumes to mount into the build step.
  Each volume is created as an empty volume prior to execution of the
  build step. Upon completion of the build, volumes and their contents
  are discarded.
  Using a named volume in only one step is not valid as it is
  indicative of a build request with an incorrect configuration.
  Structure is [documented below](#nested_volumes).

* `wait_for` -
  (Optional)
  The ID(s) of the step(s) that this build step depends on.
  This build step will not start until all the build steps in `wait_for`
  have completed successfully. If `wait_for` is empty, this build step
  will start when all previous build steps in the `Build.Steps` list
  have completed successfully.


<a name="nested_volumes"></a>The `volumes` block supports:

* `name` -
  (Required)
  Name of the volume to mount.
  Volume names must be unique per build step and must be valid names for
  Docker volumes. Each named volume must be used by at least two build steps.

* `path` -
  (Required)
  Path at which to mount the volume.
  Paths must be absolute and cannot conflict with other volume paths on
  the same build step or with certain reserved volume paths.

<a name="nested_artifacts"></a>The `artifacts` block supports:

* `images` -
  (Optional)
  A list of images to be pushed upon the successful completion of all build steps.
  The images will be pushed using the builder service account's credentials.
  The digests of the pushed images will be stored in the Build resource's results field.
  If any of the images fail to be pushed, the build is marked FAILURE.

* `objects` -
  (Optional)
  A list of objects to be uploaded to Cloud Storage upon successful completion of all build steps.
  Files in the workspace matching specified paths globs will be uploaded to the
  Cloud Storage location using the builder service account's credentials.
  The location and generation of the uploaded objects will be stored in the Build resource's results field.
  If any objects fail to be pushed, the build is marked FAILURE.
  Structure is [documented below](#nested_objects).


<a name="nested_objects"></a>The `objects` block supports:

* `location` -
  (Optional)
  Cloud Storage bucket and optional object path, in the form "gs://bucket/path/to/somewhere/".
  Files in the workspace matching any path pattern will be uploaded to Cloud Storage with
  this location as a prefix.

* `paths` -
  (Optional)
  Path globs used to match files in the build's workspace.

* `timing` -
  Output only. Stores timing information for pushing all artifact objects.
  Structure is [documented below](#nested_timing).


<a name="nested_timing"></a>The `timing` block contains:

* `start_time` -
  (Optional)
  Start of time span.
  A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
  nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".

* `end_time` -
  (Optional)
  End of time span.
  A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
  nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".

<a name="nested_options"></a>The `options` block supports:

* `source_provenance_hash` -
  (Optional)
  Requested hash for SourceProvenance.
  Each value may be one of `NONE`, `SHA256`, and `MD5`.

* `requested_verify_option` -
  (Optional)
  Requested verifiability options.
  Possible values are `NOT_VERIFIED` and `VERIFIED`.

* `machine_type` -
  (Optional)
  Compute Engine machine type on which to run the build.
  Possible values are `UNSPECIFIED`, `N1_HIGHCPU_8`, `N1_HIGHCPU_32`, `E2_HIGHCPU_8`, and `E2_HIGHCPU_32`.

* `disk_size_gb` -
  (Optional)
  Requested disk size for the VM that runs the build. Note that this is NOT "disk free";
  some of the space will be used by the operating system and build utilities.
  Also note that this is the minimum disk size that will be allocated for the build --
  the build may run with a larger disk than requested. At present, the maximum disk size
  is 1000GB; builds that request more than the maximum are rejected with an error.

* `substitution_option` -
  (Optional)
  Option to specify behavior when there is an error in the substitution checks.
  NOTE this is always set to ALLOW_LOOSE for triggered builds and cannot be overridden
  in the build configuration file.
  Possible values are `MUST_MATCH` and `ALLOW_LOOSE`.

* `dynamic_substitutions` -
  (Optional)
  Option to specify whether or not to apply bash style string operations to the substitutions.
  NOTE this is always enabled for triggered builds and cannot be overridden in the build configuration file.

* `log_streaming_option` -
  (Optional)
  Option to define build log streaming behavior to Google Cloud Storage.
  Possible values are `STREAM_DEFAULT`, `STREAM_ON`, and `STREAM_OFF`.

* `worker_pool` -
  (Optional)
  Option to specify a WorkerPool for the build. Format projects/{project}/workerPools/{workerPool}
  This field is experimental.

* `logging` -
  (Optional)
  Option to specify the logging mode, which determines if and where build logs are stored.
  Possible values are `LOGGING_UNSPECIFIED`, `LEGACY`, `GCS_ONLY`, `STACKDRIVER_ONLY`, and `NONE`.

* `env` -
  (Optional)
  A list of global environment variable definitions that will exist for all build steps
  in this build. If a variable is defined in both globally and in a build step,
  the variable will use the build step value.
  The elements are of the form "KEY=VALUE" for the environment variable "KEY" being given the value "VALUE".

* `secret_env` -
  (Optional)
  A list of global environment variables, which are encrypted using a Cloud Key Management
  Service crypto key. These values must be specified in the build's Secret. These variables
  will be available to all build steps in this build.

* `volumes` -
  (Optional)
  Global list of volumes to mount for ALL build steps
  Each volume is created as an empty volume prior to starting the build process.
  Upon completion of the build, volumes and their contents are discarded. Global
  volume names and paths cannot conflict with the volumes defined a build step.
  Using a global volume in a build with only one step is not valid as it is indicative
  of a build request with an incorrect configuration.
  Structure is [documented below](#nested_volumes).


<a name="nested_volumes"></a>The `volumes` block supports:

* `name` -
  (Optional)
  Name of the volume to mount.
  Volume names must be unique per build step and must be valid names for Docker volumes.
  Each named volume must be used by at least two build steps.

* `path` -
  (Optional)
  Path at which to mount the volume.
  Paths must be absolute and cannot conflict with other volume paths on the same
  build step or with certain reserved volume paths.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/triggers/{{trigger_id}}`

* `trigger_id` -
  The unique identifier for the trigger.

* `create_time` -
  Time when the trigger was created.


## Timeouts

This resource provides the following
[Timeouts](/docs/configuration/resources.html#timeouts) configuration options:

- `create` - Default is 4 minutes.
- `update` - Default is 4 minutes.
- `delete` - Default is 4 minutes.

## Import


Trigger can be imported using any of these accepted formats:

```
$ terraform import google_cloudbuild_trigger.default projects/{{project}}/triggers/{{trigger_id}}
$ terraform import google_cloudbuild_trigger.default {{project}}/{{trigger_id}}
$ terraform import google_cloudbuild_trigger.default {{trigger_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/guides/provider_reference.html#user_project_override).
