data aws_iam_policy_document "full-access" {
  version = "2012-10-17"
  statement {
    effect = "Allow"
    actions = [
      "dynamodb:BatchGet*",
      "dynamodb:DescribeStream",
      "dynamodb:DescribeTable",
      "dynamodb:Get*",
      "dynamodb:Query",
      "dynamodb:Scan",
      "dynamodb:BatchWrite*",
      "dynamodb:CreateTable",
      "dynamodb:Delete*",
      "dynamodb:Update*",
      "dynamodb:PutItem",
    ]
    resources = [var.table-arn]
  }
}

resource "aws_iam_policy" "full-access" {
  name = var.policy-name
  policy = data.aws_iam_policy_document.full-access.json
}
