-- +goose Up
ALTER TABLE "user" ALTER COLUMN "company_id" DROP NOT NULL;
-- +goose Down
ALTER TABLE "user" ALTER COLUMN "company_id" SET NOT NULL;
