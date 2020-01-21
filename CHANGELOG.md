# 0.1.11 (Unreleased)

BACKWARDS INCOMPATIBILITIES / NOTES:

FEATURES:

IMPROVEMENTS:

BUG FIXES:


# 0.1.10 (June 23, 2019)

IMPROVEMENTS:
 - Upgrade to Current version of Master for teamcity-go-sdk
 - Upgrade to Terraform v0.12.5 as refrence

BUG FIXES:
 - teamcity-go-sdk: Add 2 second API timeout

# 0.1.9 (May 24, 2019)
IMPROVEMENTS:
- Migrate from `govendor` to `go mod`
- Add Support for Terraform `v0.12.0`

## 0.1.8 (March 12, 2019)

BUG FIXES:
- Renaming of Teamcity Project without ForceNew  ([#25](https://github.com/bestseller-ecom/terraform-provider-teamcity/issues/25))

## 0.1.7 (September 4, 2018)

BUG FIXES:
- Suppress diff due to trailing line returns ([#22](https://github.com/bestseller-ecom/terraform-provider-teamcity/issues/22))

## 0.1.6

IMPROVEMENTS:
- Renaming of Teamcity Build Configuration without ForceNew

BUG FIXES:
- Build Configuration no longer forces a new resource on Name change.

## 0.1.5 (January 16, 2018)

BUG FIXES:
- underlying go provider had a typo for buildTypes creation ([Seen Here](https://github.com/bestseller-ecom/teamcity-sdk-go/commit/ce1da1a5348c3e788e980cea37b9b588a68c2036))

## 0.1.4 (December 21, 2017)

BACKWARDS INCOMPATIBILITIES / NOTES:

- `api_version` is set to `10.0` if your using a version different then `10.0` or below please set the `api_version` in the provider to your supported version.
- Currently Multiple Templates attached to a single Build Configuration is not supported.

FEATURES:

BUG FIXES / IMPROVEMENTS:
- API has been pinned to `10.0` because of build template api changes ([#16](https://github.com/bestseller-ecom/terraform-provider-teamcity/issues/16))

## 0.1.3 (December 6, 2017)

IMPROVEMENTS:
- **Resource:**  `parameter` Now Supports display types ([#12](https://github.com/bestseller-ecom/terraform-provider-teamcity/pull/12]))

## 0.1.2 (November 20, 2017)

BACKWARDS INCOMPATIBILITIES / NOTES:

- `agent_pool_project_attachemnt` `pool` now uses ID of the agent pool not Name. Please use the new Datasource to lookup the agent id.

FEATURES:

- **Datasource:**  `agent_pool` ([#11](https://github.com/bestseller-ecom/terraform-provider-teamcity/pull/11]))

BUG FIXES / IMPROVEMENTS:

- **Resource:**  `agent_pool_project_attachment` ([#10](https://github.com/bestseller-ecom/terraform-provider-teamcity/pull/10]))
  - Now uses id for `pool` not name


## 0.1.1 (November 17, 2017)

BACKWARDS INCOMPATIBILITIES / NOTES:

FEATURES:

- **Resource:**  `agent_pool_project_attachment` ([#10](https://github.com/bestseller-ecom/terraform-provider-teamcity/pull/10]))

IMPROVEMENTS:

- **Resource:**  `build_configuration` and `build_template` Now support Settings ([#6](https://github.com/bestseller-ecom/terraform-provider-teamcity/pull/6]))
- **Resource**: `vcs_root` Now support Propeties on Read ([#5](https://github.com/bestseller-ecom/terraform-provider-teamcity/pull/5]))
- `parameter` Now supports `read_only = true/false` ([#7](https://github.com/bestseller-ecom/terraform-provider-teamcity/pull/7]))
- `Step` Now ignores the constant `teamcity.step.mode` on read ([#8](https://github.com/bestseller-ecom/terraform-provider-teamcity/pull/8/))

BUG FIXES:

- `parameter` and `parameter_value`replaces single element not all elements ([#9](https://github.com/bestseller-ecom/terraform-provider-teamcity/pull/9/))
- Project Root is not defined in TF breaking compare ([#4](https://github.com/bestseller-ecom/terraform-provider-teamcity/pull/4]))

## 0.1.0 (October 19, 2017)

BACKWARDS INCOMPATIBILITIES / NOTES:

FEATURES:

* **New Resource**: `vcs_root` ([#1](https://github.com/bestseller-ecom/terraform-provider-teamcity/pull/1]))

IMPROVEMENTS:

* Support for Importing States on All Objects
* Added Website Docs

BUG FIXES:

- Vendored Go to keep changes at a minimum
