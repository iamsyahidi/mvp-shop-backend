# Table: products

## `Primary Key`

| `Columns`    |
| ------------ |
| id           |

## `Indexes`
| `Column`         | `Index Name`                                 | `Unique`   | `Access Method`     |
| ---------------- | -------------------------------------------- | ---------- | ------------------- |
| id               | products_pkey                                | `Yes`      | btree               |
| name             | idx_products_name                            | `No`       | btree               |
| price            | idx_products_price                           | `No`       | btree               |
| stock            | idx_products_stock                           | `No`       | btree               |
| status           | idx_products_status                          | `No`       | btree               |



## `Foreign Keys`

## `Columns`

| `Name`         | `Type`                                 | `Nullable` | `Default`           | `Comment`            |
| -------------- | -------------------------------------- | ---------- | ------------------- | -------------------- |
| id             | varchar(36)                            | `No`       |                     |                      |
| name           | varchar(250)                           | `No`       |                     |                      |
| price          | numeric                                | `Yes`      |                     |                      |
| stock          | numeric                                | `Yes`      |                     |                      |
| status         | varchar(10)                            | `No`       |                     |                      |
| category_id    | varchar(36)                            | `No`       |                     |                      |
| created_at     | timestamptz                            | `No`       | now()               |                      |
| created_by     | varchar(150)                           | `No`       |                     |                      |
| updated_at     | timestamptz                            | `Yes`      | current_timestamp   |                      |
| updated_by     | varchar(150)                           | `Yes`      |                     |                      |