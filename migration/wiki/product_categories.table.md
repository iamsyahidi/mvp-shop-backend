# Table: product_categories

## `Primary Key`

| `Columns`    |
| ------------ |
| id           |

## `Indexes`
| `Column`         | `Index Name`                                 | `Unique`   | `Access Method`     |
| ---------------- | -------------------------------------------- | ---------- | ------------------- |
| id               | product_categories_pkey                      | `Yes`      | btree               |
| name             | idx_product_categories_name                  | `No`       | btree               |
| status           | idx_product_categories_status                | `No`       | btree               |



## `Foreign Keys`

## `Columns`

| `Name`         | `Type`                                 | `Nullable` | `Default`           | `Comment`            |
| -------------- | -------------------------------------- | ---------- | ------------------- | -------------------- |
| id             | varchar(36)                            | `No`       |                     |                      |
| name           | varchar(250)                           | `No`       |                     |                      |
| status         | varchar(10)                            | `No`       |                     |                      |
| created_at     | timestamptz                            | `No`       | now()               |                      |
| created_by     | varchar(150)                           | `No`       |                     |                      |
| updated_at     | timestamptz                            | `Yes`      | current_timestamp   |                      |
| updated_by     | varchar(150)                           | `Yes`      |                     |                      |