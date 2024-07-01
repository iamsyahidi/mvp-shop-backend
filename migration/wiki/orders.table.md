# Table: orders

## `Primary Key`

| `Columns`    |
| ------------ |
| id           |

## `Indexes`
| `Column`         | `Index Name`                                 | `Unique`   | `Access Method`     |
| ---------------- | -------------------------------------------- | ---------- | ------------------- |
| id               | orders_pkey                                  | `Yes`      | btree               |
| amount           | idx_orders_amount                            | `No`       | btree               |
| invoice          | idx_orders_invoice                           | `No`       | btree               |
| payment          | idx_orders_payment                           | `No`       | btree               |
| status           | idx_orders_status                            | `No`       | btree               |
| customer_id      | idx_orders_customer_id                       | `No`       | btree               |



## `Foreign Keys`

## `Columns`

| `Name`         | `Type`                                 | `Nullable` | `Default`           | `Comment`            |
| -------------- | -------------------------------------- | ---------- | ------------------- | -------------------- |
| invoice        | varchar(100)                           | `No`       |                     |                      |
| customer_id    | varchar(36)                            | `No`       |                     |                      |
| amount         | numeric                                | `Yes`      |                     |                      |
| payment        | bool                                   | `No`       | false               |                      |
| status         | varchar(10)                            | `No`       |                     |                      |
| created_at     | timestamptz                            | `No`       | now()               |                      |
| created_by     | varchar(150)                           | `No`       |                     |                      |
| updated_at     | timestamptz                            | `Yes`      | current_timestamp   |                      |
| updated_by     | varchar(150)                           | `Yes`      |                     |                      |