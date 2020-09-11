CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE banks (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	"name" varchar NOT NULL UNIQUE,
	active bool NOT NULL DEFAULT true,
	CONSTRAINT banks_pk PRIMARY KEY (id)
);