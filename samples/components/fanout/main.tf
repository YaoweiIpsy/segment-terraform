resource "segment_source" "fanout_source" {
  source_name = "fanout-${var.isDev ? "staging": "prod"}-${var.name}"
  display_name = "<${var.isDev ? "staging": "prod"}>[fanout] - ${var.name}"
  catalog_name = "catalog/sources/java"
}

resource "segment_destination" "fanout_lambda" {
  source_name = "${segment_source.fanout_source.source_name}"
  destination_name = "amazon-lambda"
  configs = [
    {
      name = "region"
      type = "string"
      value = "us-east-1"
    },
    {
      name = "roleAddress"
      type = "string"
      value = "arn:aws:iam::${var.isDev ?"450096215204" : "769100407790"}:role/${var.role}"
    },
    {
      name = "function"
      type = "string"
      value = "arn:aws:lambda:us-east-1:${var.isDev ?"450096215204" : "769100407790"}:function:${var.function_name}"
    }
  ]
}

resource "segment_destination_filter" "fanout_filter" {
  source_name = "${segment_destination.fanout_lambda.source_name}"
  destination_name = "${segment_destination.fanout_lambda.destination_name}"
  condition = "event != \"${join("\" and event != \"", var.events)}\""
  //"event != \"usms\" and event != \"usms_test\""
  actions = [
    "{\"type\": \"drop_event\"}"
  ]
  title = "<${var.isDev ? "staging": "prod"}>[${var.name}] - allow [${join(",", var.events)}]"
  count = "${length(var.events) > 0 ? 1 : 0}"
  enabled = true
}