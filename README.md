# tf-provider
Learning to build a TF provider

Notes
- You must build against the api with the same version of Terraform as the executable you will use to run plans.
- Naming of exe is important, `go build -o terraform-provider-dummy`
- Naming of the resources *MUST* start with the provider name used in the executable. If they don't start with dummy they won't be found.
