ALTER TABLE practicum_module_content
ALTER COLUMN content TYPE JSONB
USING content::JSONB;
