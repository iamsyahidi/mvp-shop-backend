# Table: customers

## `Primary Key`

| `Columns`    |
| ------------ |
| id           |

## `Indexes`
| `Column`         | `Index Name`                                 | `Unique`   | `Access Method`     |
| ---------------- | -------------------------------------------- | ---------- | ------------------- |
| id               | customers_pkey                               | `Yes`      | btree               |
| email            | idx_customers_email                          | `Yes`      | btree               |
| name             | idx_customers_name                           | `No`       | btree               |
| password         | idx_customers_password                       | `No`       | btree               |
| status           | idx_customers_status                         | `No`       | btree               |



## `Foreign Keys`

## `Columns`

| `Name`         | `Type`                                 | `Nullable` | `Default`           | `Comment`            |
| -------------- | -------------------------------------- | ---------- | ------------------- | -------------------- |
| id             | varchar(36)                            | `No`       |                     |                      |
| email          | varchar(100)                           | `No`       |                     |                      |
| name           | varchar(250)                           | `No`       |                     |                      |
| password       | varchar(150)                           | `No`       |                     |                      |
| status         | varchar(10)                            | `No`       |                     |                      |
| created_at     | timestamptz                            | `No`       | now()               |                      |
| created_by     | varchar(150)                           | `No`       |                     |                      |
| updated_at     | timestamptz                            | `Yes`      | current_timestamp   |                      |
| updated_by     | varchar(150)                           | `Yes`      |                     |                      |