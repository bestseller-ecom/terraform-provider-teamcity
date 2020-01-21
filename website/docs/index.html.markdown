---
layout: "teamcity"
page_title: "Provider: TeamCity"
sidebar_current: "docs-teamcity-index"
description: |-
  The TeamCity provider is used to interact with JetBrains TeamCity server. The provider needs to be configured with the proper credentials before it can be used.
---


# TeamCity Provider

The TeamCity provider is used to interact with
[TeamCity Server by JetBrains](https://www.jetbrains.com/teamcity/). The provider needs
to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.


## Example Usage

```hcl
// Configure the TeamCity provider
provider "teamcity" {
  // url         = "${var.teamcity_url}"
  // api_version = "${var.teamcity_api_version}"
  // username    = "${var.teamcity_username}"
  // password    = "${var.teamcity_password}"
}
```


## Requirements

Compatibility is defined by [teamcity-go-sdk](https://github.com/bestseller-ecom/teamcity-go-sdk) which bundled with this provider.

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)
-	[TeamCity](https://www.jetbrains.com/teamcity/)) 9.x - latest


## Argument Reference

The following arguments are supported in the `provider` block:

* `url` - (optional) This is the TeamCity Server URL e.g. https://teamcity.domain.com:8111.
  It must be provided but it can also be sourced from the `TEAMCITY_URL` environment variable.
  Defaults to `http://localhost:8111`

* `api_version` - (optional) This is the TeamCity Server REST API Version e.g. `latest`.
  It must be provided but it can also be sourced from the `TEAMCITY_API_VERSION` environment variable.
  Defaults to `10.0`

* `username` - (Optional) This is the TeamCity username. It must be provided, but
  it can also be sourced from the `TEAMCITY_USERNAME` environment variable.

* `password` - (Optional) This is the TeamCity Password. It must be provided, but
  it can also be sourced from the `TEAMCITY_PASSWORD` environment variable.

