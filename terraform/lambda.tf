module "main-lambda" {
  source = "./modules/main-lambda"

  inventory-itemTableName = module.inventory-item-db.table-name
}
