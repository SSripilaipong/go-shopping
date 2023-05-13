resource "aws_iam_role_policy_attachment" "attachment" {
  role = var.lambda-role-id
  policy_arn = aws_iam_policy.full-access.arn
}
