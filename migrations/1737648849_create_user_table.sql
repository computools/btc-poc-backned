-- +goose Up
CREATE TABLE IF NOT EXISTS "user"
(
    "id"          bigserial PRIMARY KEY,
    "full_name"   text    NOT NULL,
    "position"    text,
    "company_id"  integer NOT NULL REFERENCES company (id) ON DELETE CASCADE,
    "image_url"   text,
    "keycloak_id" text
);
-- +goose Down
DROP TABLE "user";
