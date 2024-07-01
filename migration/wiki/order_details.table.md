# Table: order_details

## `Primary Key`

| `Columns`    |
| ------------ |
| id           |

## `Indexes`
| `Column`         | `Index Name`                                 | `Unique`   | `Access Method`     |
| ---------------- | -------------------------------------------- | ---------- | ------------------- |
| id               | order_details_pkey                           | `Yes`      | btree               |
| amount           | idx_order_details_amount                     | `No`       | btree               |
| invoice          | idx_order_details_invoice                    | `No`       | btree               |
| price            | idx_order_details_price                      | `No`       | btree               |
| product_id       | idx_order_details_product_id                 | `No`       | btree               |
| qty              | idx_order_details_qty                        | `No`       | btree               |
| status           | idx_order_details_status                     | `No`       | btree               |



## `Foreign Keys`

## `Columns`

| `Name`         | `Type`                                 | `Nullable` | `Default`           | `Comment`            |
| -------------- | -------------------------------------- | ---------- | ------------------- | -------------------- |
| invoice        | varchar(100)                           | `No`       |                     |                      |
| product_id     | varchar(36)                            | `No`       |                     |                      |
| qty            | numeric                                | `Yes`      |                     |                      |
| price          | numeric                                | `Yes`      |                     |                      |
| amount         | numeric                                | `Yes`      |                     |                      |
| status         | varchar(10)                            | `No`       |                     |                      |
| created_at     | timestamptz                            | `No`       | now()               |                      |
| created_by     | varchar(150)                           | `No`       |                     |                      |
| updated_at     | timestamptz                            | `Yes`      | current_timestamp   |                      |
| updated_by     | varchar(150)                           | `Yes`      |                     |                      |