resource "aws_dynamodb_table" "table" {
  name           = "inventory-item-repo"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "Id"

  attribute {
    name = "Id"
    type = "S"
  }
}
