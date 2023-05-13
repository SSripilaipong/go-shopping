module "inventory-item-db" {
  source = "../core/inventory/terraform/item-db"
  table-name = "inventory-item-repo"

  event-stream-enabled = true
}

module "inventory-item-db-access" {
  source = "../core/common/terraform/dynamodb/lambda-full-access"
  policy-name     = "${module.inventory-item-db.table-name}-lambda-access"

  lambda-role-id = module.main-lambda.role-id
  table-arn       = module.inventory-item-db.table-arn
}

module "inventory-item-db-event-outbox" {
  source = "../core/common/terraform/dynamodb/event-stream-lambda-trigger"
  lambda-arn = module.main-lambda.lambda-arn
  lambda-role-id = module.main-lambda.role-id

  stream-arn = module.inventory-item-db.event-stream-arn
}
