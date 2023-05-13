resource "aws_dynamodb_table" "table" {
  name           = var.table-name
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "Id"

  stream_enabled = var.event-stream-enabled
  stream_view_type = var.event-stream-enabled ? "NEW_IMAGE" : null

  attribute {
    name = "Id"
    type = "S"
  }
}
