-- 1. Revert the `id` column back to SERIAL by creating a sequence and setting it as the default
ALTER TABLE operation_types ALTER COLUMN id TYPE SERIAL USING id::SERIAL;
