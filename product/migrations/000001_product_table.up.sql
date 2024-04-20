CREATE TABLE "product" (
  "id" SERIAL PRIMARY KEY,
  "name" text,
  "image" text,
  "price" float,
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamp
);