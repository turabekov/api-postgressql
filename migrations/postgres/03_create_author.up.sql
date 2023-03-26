CREATE TABLE "book" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "count" INT CHECK(count >= 0) DEFAULT 0,
    "income_price" NUMERIC CHECK(income_price >= 0) DEFAULT 0,
    "profit_status" VARCHAR CHECK(profit_status = 'fixed' OR profit_status = 'precent') DEFAULT 'fixed',
    "profit_price" NUMERIC CHECK(profit_price >= 0) DEFAULT 0,
    "sell_price"  NUMERIC CHECK(sell_price >= 0) DEFAULT 0,
    "author_id" UUID NOT NULL REFERENCES author(id),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);