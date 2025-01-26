-- +goose Up
CREATE TABLE IF NOT EXISTS "financial_reports"
(
    "id"         bigserial PRIMARY KEY,
    "name"       text    NOT NULL,
    "year"       integer,
    "company_id" integer NOT NULL REFERENCES company (id) ON DELETE CASCADE,
    "url"        text
);
-- +goose Down
DROP TABLE "financial_reports";
