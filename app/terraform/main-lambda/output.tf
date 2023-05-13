output "role-id" {
  value = aws_iam_role.main-iam-role.id
}

output "lambda-arn" {
  value = aws_lambda_function.main.arn
}