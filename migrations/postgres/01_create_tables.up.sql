CREATE TABLE "author" (
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);