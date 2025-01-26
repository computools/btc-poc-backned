-- +goose Up
CREATE TABLE IF NOT EXISTS "company"
(
    "id"               bigserial PRIMARY KEY,
    "public_id"        text UNIQUE NOT NULL,
    "name"             text,
    "short_name"       text,
    "physical_address" text,
    "physical_city"    text,
    "physical_zip"     text,
    "postal_address"   text,
    "postal_city"      text,
    "postal_zip"       text
);
-- +goose Down
DROP TABLE "company";
