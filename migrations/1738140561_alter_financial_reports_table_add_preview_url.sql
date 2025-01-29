-- +goose Up
ALTER TABLE "financial_reports" ADD COLUMN "preview_url" text;
-- +goose Down
ALTER TABLE "financial_reports" DROP COLUMN "preview_url";

