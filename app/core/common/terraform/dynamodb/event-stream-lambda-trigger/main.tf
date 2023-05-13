resource "aws_lambda_event_source_mapping" "example" {
  event_source_arn  = var.stream-arn
  function_name     = var.lambda-arn
  starting_position = "LATEST"
}
