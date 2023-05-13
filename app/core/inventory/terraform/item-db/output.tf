output "table-name" {
  value = aws_dynamodb_table.table.name
}

output "table-arn" {
  value = aws_dynamodb_table.table.arn
}

output "event-stream-arn" {
  value = aws_dynamodb_table.table.stream_arn
}
