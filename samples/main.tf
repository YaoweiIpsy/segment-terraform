provider "segment" {
  access_token = "${var.token}"
  workspace = "ipsy-production"
}
locals {
  isDev = true
}
module "yaowei-fanout" {
  source = "./components/fanout"
  function_name = "yaowei_segment_io"
  events = [
    "hahah",
    "asfdasdf"]
  name = "yaowei-test"
}

module "elaine-fanout" {
  source = "./components/fanout"
  function_name = "elaine_segment_io"
  events = [
    "hahah",
    "asfdasdf"]
  name = "elaine-test"
}

resource "segment_source" "event" {
  source_name = "yaowei-source"
  display_name = "<${local.isDev ? "staging" : "prod"}> - Yaowei , test terraform"
  catalog_name = "catalog/sources/javascript"
}

resource "segment_destination" "test_destination" {
  source_name = "${segment_source.event.source_name}"
  destination_name = "repeater"
  configs = [
    {
      name = "repeatKeys"
      type = "list"
      list = [
        "${module.yaowei-fanout.fanout_write-key}",
        "${module.elaine-fanout.fanout_write-key}"
      ]
    },
  ]
}
