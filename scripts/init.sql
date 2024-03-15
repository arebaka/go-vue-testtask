CREATE TABLE IF NOT EXISTS public."Products" (
	"ID" bigserial NOT NULL PRIMARY KEY,
	"Name" varchar(255) NOT NULL,
	"Description" text NOT NULL,
	"Price" bigint NOT NULL,
	"CreatedAt" timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS products_name_idx
ON public."Products"
USING btree ("Name");
