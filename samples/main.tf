provider "segment" {
  access_token = "XUnxSMBgMjJH4br4XcCrY0lk665OnpadBT6AUPFE4eQ.e2Lm6YK-zLTfpbqc0-ap7FEiKeKDacvLNMLdDaiPZqo"
  workspace = "ipsy-production"
}
locals {
  isDev = true
}
module "fanout" {
  source = "./components/fanout"
  function_name = "yaowei_segment_io"
  events = ["hahah","asfdasdf"]
  name = "yaowei-test"
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
      list = ["${module.fanout.fanout_write-key}"]
    },
  ]
}
