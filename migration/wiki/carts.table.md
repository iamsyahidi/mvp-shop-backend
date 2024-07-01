# Table: carts

## `Primary Key`

| `Columns`    |
| ------------ |
| id           |

## `Indexes`
| `Column`         | `Index Name`                                 | `Unique`   | `Access Method`     |
| ---------------- | -------------------------------------------- | ---------- | ------------------- |
| id               | carts_pkey                                   | `Yes`      | btree               |
| customer_id      | idx_carts_customer_id                        | `No`       | btree               |
| product_id       | idx_carts_product_id                         | `No`       | btree               |



## `Foreign Keys`

## `Columns`

| `Name`         | `Type`                                 | `Nullable` | `Default`           | `Comment`            |
| -------------- | -------------------------------------- | ---------- | ------------------- | -------------------- |
| id             | varchar(36)                            | `No`       |                     |                      |
| customer_id    | varchar(36)                            | `No`       |                     |                      |
| product_       | varchar(36)                            | `No`       |                     |                      |
| qty            | integer                                | `Yes`      |                     |                      |
| price          | numeric                                | `Yes`      |                     |                      |
| amount         | numeric                                | `Yes`      |                     |                      |
| status         | varchar(10)                            | `No`       |                     |                      |
| created_at     | timestamptz                            | `No`       | now()               |                      |
| created_by     | varchar(150)                           | `No`       |                     |                      |
| updated_at     | timestamptz                            | `Yes`      | current_timestamp   |                      |
| updated_by     | varchar(150)                           | `Yes`      |                     |                      |