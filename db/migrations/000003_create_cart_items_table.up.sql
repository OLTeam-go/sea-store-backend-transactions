CREATE TABLE cart_items (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	item_id uuid NOT NULL,
	cart_id uuid NOT NULL,
	created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NOT NULL DEFAULT now(),
	quantity int4 NOT NULL DEFAULT 1,
	deleted_at timestamp NULL,
	CONSTRAINT cart_items_pk PRIMARY KEY (id),
	CONSTRAINT cart_items_unique_item UNIQUE (item_id, cart_id, deleted_at)
);

ALTER TABLE cart_items ADD CONSTRAINT cart_items_fk FOREIGN KEY (cart_id) REFERENCES carts(id);