-- 1. Drop the default sequence associated with the SERIAL type
ALTER TABLE operation_types ALTER COLUMN id DROP DEFAULT;

-- 2. Change the type from SERIAL to INT
ALTER TABLE operation_types ALTER COLUMN id TYPE INT USING id::INT;
