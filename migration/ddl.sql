CREATE TABLE IF NOT EXISTS public.customer (
	id varchar(36) NOT NULL,
	email varchar(100) NOT NULL,
	"name" varchar(250) NOT NULL,
	"password" varchar(150) NULL,
	"status" varchar(10) NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL,
	CONSTRAINT customer_pkey PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS idx_customer_email ON public.customer USING btree (email);
CREATE INDEX IF NOT EXISTS idx_customer_id ON public.customer USING btree (id);
CREATE INDEX IF NOT EXISTS idx_customer_name ON public.customer USING btree ("name");
CREATE INDEX IF NOT EXISTS idx_customer_password ON public.customer USING btree ("password");
CREATE INDEX IF NOT EXISTS idx_customer_status ON public.customer USING btree ("status");

INSERT INTO public.customer (id, email, "name", "password", "status", created_at, created_by, updated_at, updated_by) VALUES(gen_random_uuid(), 'ilhamsyahidi66@gmail.com', 'ilham', '$2a$10$2RLJwYzS2wCGTkgr145miOvXe52nI0JRE.8uOgqEmM8ulmmRufWle', 'active', now(), 'ilham', NULL, NULL);