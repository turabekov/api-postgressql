CREATE TABLE "users" (
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR NOT NULL,
    "balance" NUMERIC CHECK(balance >= 0) DEFAULT 0,
    "updated_at" TIMESTAMP,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);