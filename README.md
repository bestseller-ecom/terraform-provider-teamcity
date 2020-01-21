# terraform-provider-teamcity
Terraform provider for TeamCity

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
  -[Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)
  -[teamcity-sdk-go](https://github.com/bestseller-ecom/teamcity-sdk-go) (SDK dependency for building the provider)


Documentation
-------------

- [TeamCity Provider](website/docs/index.html.markdown)
  - Datasources
    - [agent_pool](website/docs/d/agent_pool.html.markdown)
  - Resources
    - [agent_pool_project_attachment](website/docs/r/agent_pool_project_attachment.html.markdown)
    - [project](website/docs/r/project.html.markdown)
    - [vcs_root](website/docs/r/vcs_root.html.markdown)
    - [build_configuration](website/docs/r/build_configuration.html.markdown)
    - [build_template](website/docs/r/build_template.html.markdown)


## Installing the Provider

```bash
#Checkout the current release tag
make install
```



'Developing the Provider
-----------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.9+ is *required*).
You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-teamcity
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources and require the Docker container to be running.
  Currently we use [teamcity-dsk-go](https://github.com/bestseller-ecom/teamcity-sdk-go) as the repository to start the TeamCity container

```sh
$ make testacc
```

### Development Resources on the Web

- [Perl5 TeamCity API](http://eilara.github.io/perl5-teamcity-api/)
- [Swagger](https://dploeger.github.io/teamcity-rest-api)
- [JetBrains REST-API (SVN)](http://svn.jetbrains.org/teamcity/plugins/rest-api/branches/)
- [JetBrains Tickets](https://youtrack.jetbrains.com/issues/TW)
- [JetBrains Support](https://teamcity-support.jetbrains.com)
- [Java Docs for TeamCity](http://javadoc.jetbrains.net/teamcity/openapi/current/)
