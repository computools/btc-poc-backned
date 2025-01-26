-- +goose Up
ALTER TABLE "user" ADD COLUMN "first_name" text;
ALTER TABLE "user" ADD COLUMN "last_name" text;
ALTER TABLE "user" DROP COLUMN "full_name";
-- +goose Down
ALTER TABLE "user" DROP COLUMN "first_name";
ALTER TABLE "user" DROP COLUMN "last_name";
ALTER TABLE "user" ADD COLUMN "full_name" text;
