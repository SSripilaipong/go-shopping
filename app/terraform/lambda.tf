module "main-lambda" {
  source = "./main-lambda"

  binary_file_path = "../../build/main"
  zip_file_path = "../../build/main.zip"

  inventory-itemTableName = module.inventory-item-db.table-name
}
