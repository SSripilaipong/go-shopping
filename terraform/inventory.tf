module "inventory-item-db" {
  source = "./modules/inventory/item-db"
  table-name = "inventory-item-repo"
}

module "inventory-item-db-access" {
  source = "./modules/dynamodb/lambda-full-access"
  policy-name     = "${module.inventory-item-db.table-name}-lambda-access"

  lambda-role-id = module.main-lambda.role-id
  table-arn       = module.inventory-item-db.table-arn
}
