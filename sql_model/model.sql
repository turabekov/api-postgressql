
CREATE TABLE "book" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "price" NUMERIC NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "author_id" UUID NOT NULL REFERENCES author(id)
);

CREATE TABLE "author" (
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


ALTER TABLE "book" ADD COLUMN author_id UUID NOT NULL REFERENCES author(id);