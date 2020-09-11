CREATE TABLE snapshot_cart_items (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	item_id uuid NOT NULL,
  cart_id uuid NOT NULL,
	transaction_id uuid NOT NULL,
	merchant_id uuid NOT NULL,
	"name" varchar NOT NULL,
	category varchar NOT NULL,
	description varchar NULL,
	price numeric NOT NULL,
	quantity int NOT NULL DEFAULT 1,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NOT NULL DEFAULT now(),
	CONSTRAINT snapshot_cart_items_pk PRIMARY KEY (id),
	CONSTRAINT snapshot_cart_items_fk_1 FOREIGN KEY (transaction_id) REFERENCES transactions(id),
  CONSTRAINT snapshot_cart_items_fk_2 FOREIGN KEY (cart_id) REFERENCES carts(id)
);
