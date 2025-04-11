-- Remove index
DROP INDEX IF EXISTS idx_practicum_module_content_material_id;

-- Drop unique constraint
ALTER TABLE practicum_module_content
DROP CONSTRAINT IF EXISTS unique_material_id;

-- Drop the column
ALTER TABLE practicum_module_content
DROP COLUMN IF EXISTS material_id;