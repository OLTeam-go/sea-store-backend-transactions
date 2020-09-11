CREATE TABLE transactions (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	customer_id uuid NOT NULL,
	cart_id uuid NOT NULL,
	bank_id uuid NOT NULL,
	bank_account_number varchar NOT NULL,
	status int4 NOT NULL DEFAULT 1,
	"cost" numeric NOT NULL DEFAULT 0,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NOT NULL DEFAULT now(),
	CONSTRAINT transactions_pk PRIMARY KEY (id),
	CONSTRAINT transactions_unique UNIQUE (customer_id, cart_id)
);

ALTER TABLE transactions ADD CONSTRAINT transactions_fk FOREIGN KEY (bank_id) REFERENCES banks(id);
ALTER TABLE transactions ADD CONSTRAINT transactions_fk_1 FOREIGN KEY (cart_id) REFERENCES carts(id);