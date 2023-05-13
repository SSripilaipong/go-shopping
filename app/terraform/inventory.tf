module "inventory-item-db" {
  source = "../core/inventory/terraform/item-db"
  table-name = "inventory-item-repo"
}

module "inventory-item-db-access" {
  source = "../core/common/terraform/dynamodb/lambda-full-access"
  policy-name     = "${module.inventory-item-db.table-name}-lambda-access"

  lambda-role-id = module.main-lambda.role-id
  table-arn       = module.inventory-item-db.table-arn
}
