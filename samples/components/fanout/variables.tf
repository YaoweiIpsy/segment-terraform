variable "name" {}
variable "isDev" {
  default = true
}
variable "role" {
  default = "yaowei-segment"
}
variable "function_name" {
}

variable "events" {
  type = "list"
  default = []
}
