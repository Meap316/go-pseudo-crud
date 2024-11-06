-- Modify "users" table
ALTER TABLE "public"."users" DROP COLUMN "is_deleted";
ALTER TABLE "public"."users" ADD COLUMN "is_deleted" boolean NULL;

