data "aws_iam_policy_document" "main-assume-role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "main-iam-role" {
  name               = "lambda-main-iam"
  assume_role_policy = data.aws_iam_policy_document.main-assume-role.json
}
