CREATE TABLE IF NOT EXISTS public.customers (
	id varchar(36) NOT NULL,
	email varchar(100) NOT NULL,
	"name" varchar(250) NOT NULL,
	"password" varchar(150) NOT NULL,
	"status" varchar(10) NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL,
	CONSTRAINT customers_pkey PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS idx_customers_email ON public.customers USING btree (email);
CREATE INDEX IF NOT EXISTS idx_customers_id ON public.customers USING btree (id);
CREATE INDEX IF NOT EXISTS idx_customers_name ON public.customers USING btree ("name");
CREATE INDEX IF NOT EXISTS idx_customers_password ON public.customers USING btree ("password");
CREATE INDEX IF NOT EXISTS idx_customers_status ON public.customers USING btree ("status");

ALTER TABLE IF EXISTS "customers" ADD CONSTRAINT "uni_customers_email" UNIQUE ("email");

CREATE TABLE IF NOT EXISTS public.product_categories (
	id varchar(36) NOT NULL,
	"name" varchar(250) NOT NULL,
	"status" varchar(10) NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL,
	CONSTRAINT product_categories_pkey PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS idx_product_categories_id ON public.product_categories USING btree (id);
CREATE INDEX IF NOT EXISTS idx_product_categories_name ON public.product_categories USING btree ("name");
CREATE INDEX IF NOT EXISTS idx_product_categories_status ON public.product_categories USING btree ("status");

CREATE TABLE IF NOT EXISTS public.products (
	id varchar(36) NOT NULL,
	"name" varchar(250) NOT NULL,
	price numeric NULL,
	stock numeric NULL,
	"status" varchar(10) NOT NULL,
	"category_id" varchar(36) NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL,
	CONSTRAINT products_pkey PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS idx_products_id ON public.products USING btree (id);
CREATE INDEX IF NOT EXISTS idx_products_name ON public.products USING btree ("name");
CREATE INDEX IF NOT EXISTS idx_products_price ON public.products USING btree (price);
CREATE INDEX IF NOT EXISTS idx_products_stock ON public.products USING btree (stock);
CREATE INDEX IF NOT EXISTS idx_products_status ON public.products USING btree ("status");

CREATE TABLE IF NOT EXISTS public.carts (
	id varchar(36) NOT NULL,
    customer_id varchar(36) NOT NULL,
    product_id varchar(36) NOT NULL,
    qty integer NULL,
    price numeric NULL,
	amount numeric NULL,
	"status" varchar(10) NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL,
   	CONSTRAINT carts_pkey PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_carts_id ON public.carts USING btree (id);
CREATE INDEX IF NOT EXISTS idx_carts_customer_id ON public.carts USING btree (customer_id);
CREATE INDEX IF NOT EXISTS idx_carts_product_id ON public.carts USING btree (product_id);

CREATE TABLE IF NOT EXISTS public."orders" (
	invoice varchar(100) NOT NULL,
	customer_id varchar(36) NOT NULL,
	amount numeric NULL,
	payment bool DEFAULT false NOT NULL,
	"status" varchar(10) NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL,
	CONSTRAINT orders_pkey PRIMARY KEY (invoice)
);
CREATE INDEX IF NOT EXISTS idx_orders_amount ON public."orders" USING btree (amount);
CREATE INDEX IF NOT EXISTS idx_orders_invoice ON public."orders" USING btree (invoice);
CREATE INDEX IF NOT EXISTS idx_orders_payment ON public."orders" USING btree (payment);
CREATE INDEX IF NOT EXISTS idx_orders_status ON public."orders" USING btree ("status");
CREATE INDEX IF NOT EXISTS idx_orders_customer_id ON public."orders" USING btree (customer_id);


CREATE TABLE IF NOT EXISTS public.order_details (
	invoice varchar(100) NOT NULL,
	product_id varchar(36) NOT NULL,
	qty numeric NULL,
	price numeric NULL,
	amount numeric NULL,
	"status" varchar(10) NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL
);
CREATE INDEX IF NOT EXISTS idx_order_details_amount ON public.order_details USING btree (amount);
CREATE INDEX IF NOT EXISTS idx_order_details_invoice ON public.order_details USING btree (invoice);
CREATE INDEX IF NOT EXISTS idx_order_details_price ON public.order_details USING btree (price);
CREATE INDEX IF NOT EXISTS idx_order_details_product_id ON public.order_details USING btree (product_id);
CREATE INDEX IF NOT EXISTS idx_order_details_qty ON public.order_details USING btree (qty);
CREATE INDEX IF NOT EXISTS idx_order_details_status ON public.order_details USING btree ("status");




