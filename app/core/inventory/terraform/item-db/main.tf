resource "aws_dynamodb_table" "table" {
  name           = var.table-name
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "Id"

  attribute {
    name = "Id"
    type = "S"
  }
}
