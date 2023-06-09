data "archive_file" "main" {
  type        = "zip"
  source_file = var.binary_file_path
  output_path = var.zip_file_path
}

resource "aws_lambda_function" "main" {
  filename      = data.archive_file.main.output_path
  function_name = "go_shopping-main-prod"
  role          = aws_iam_role.main-iam-role.arn
  handler       = "main"

  source_code_hash = data.archive_file.main.output_base64sha256

  runtime = "go1.x"

  environment {
    variables = {
      INVENTORY_ITEM_TABLE_NAME = var.inventory-itemTableName
    }
  }
}

resource "aws_lambda_function_url" "main" {
  function_name      = aws_lambda_function.main.function_name
  authorization_type = "NONE"
}
