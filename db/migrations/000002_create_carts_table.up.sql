CREATE TABLE carts (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	customer_id uuid NOT NULL,
	active boolean NOT NULL DEFAULT true,
	created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NOT NULL DEFAULT now(),
	CONSTRAINT carts_pk PRIMARY KEY (id)
);
