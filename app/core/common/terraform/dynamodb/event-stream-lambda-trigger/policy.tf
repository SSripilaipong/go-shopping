data aws_iam_policy_document "stream-access" {
  version = "2012-10-17"
  statement {
    effect = "Allow"
    actions = [
      "dynamodb:DescribeStream",
      "dynamodb:GetRecords",
      "dynamodb:GetShardIterator",
      "dynamodb:ListStreams"
    ]
    resources = [var.stream-arn]
  }
}

resource "aws_iam_policy" "stream-access" {
  policy = data.aws_iam_policy_document.stream-access.json
}
