provider "terraform-provider-dummy"{}

data "dummy_data" "api" {
  service = "api"
}

resource "dummy_server" "foo" {
  address = "${data.dummy_data.api.name}"
}
