-- +goose Up
CREATE TABLE IF NOT EXISTS "company_performance"
(
    "id"                      bigserial PRIMARY KEY,
    "company_id"              integer NOT NULL REFERENCES company (id) ON DELETE CASCADE,
    "quarter"                 integer,
    "year"                    integer,
    "income"                  numeric,
    "total_wage_bill"         numeric,
    "executive_wage_bill"     numeric,
    "total_staff"             integer,
    "total_executive"         integer,
    "total_assets"            numeric,
    "operational_expenditure" numeric
);
-- +goose Down
DROP TABLE "company_performance";
