CREATE TABLE IF NOT EXISTS public.customers (
	id varchar(36) NOT NULL,
	email varchar(100) NOT NULL,
	"name" varchar(250) NOT NULL,
	"password" varchar(150) NULL,
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

INSERT INTO public.customers (id, email, "name", "password", "status", created_at, created_by, updated_at, updated_by) VALUES(gen_random_uuid(), 'ilhamsyahidi66@gmail.com', 'ilham', '$2a$10$2RLJwYzS2wCGTkgr145miOvXe52nI0JRE.8uOgqEmM8ulmmRufWle', 'active', now(), 'ilham', NULL, NULL);

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
