-- 2025-03-21 15:39

ALTER TABLE pismo."accounts" ADD COLUMN "available_limit" NUMERIC(10, 2);
ALTER TABLE pismo."accounts" ADD COLUMN "current_limit" NUMERIC(10, 2);
