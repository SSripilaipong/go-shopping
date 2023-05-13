data "aws_iam_policy_document" "logging-policy" {
  version = "2012-10-17"
  statement {
    effect = "Allow"

    resources = ["arn:aws:logs:*:*:*"]

    actions = [
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]
  }
}

resource "aws_iam_policy" "logging-policy" {
  name   = "lambda-main-logging-policy"
  policy = data.aws_iam_policy_document.logging-policy.json
}

resource "aws_iam_role_policy_attachment" "logging-policy-attachment" {
  role = aws_iam_role.main-iam-role.id
  policy_arn = aws_iam_policy.logging-policy.arn
}

resource "aws_cloudwatch_log_group" "main-logging" {
  name              = "/aws/lambda/${aws_lambda_function.main.function_name}"
  retention_in_days = 1
  lifecycle {
    prevent_destroy = false
  }
}
