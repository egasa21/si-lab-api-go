CREATE EXTENSION IF NOT EXISTS "pgcrypto";

ALTER TABLE practicum_module_content
ADD COLUMN material_id UUID DEFAULT gen_random_uuid ();

UPDATE practicum_module_content
SET
    material_id = gen_random_uuid ()
WHERE
    material_id IS NULL;

ALTER TABLE practicum_module_content
ALTER COLUMN material_id
SET
    NOT NULL;

ALTER TABLE practicum_module_content ADD CONSTRAINT unique_material_id UNIQUE (material_id);

CREATE INDEX IF NOT EXISTS idx_practicum_module_content_material_id ON practicum_module_content (material_id);