ALTER TABLE practicum_module_content
ALTER COLUMN content TYPE TEXT
USING content::TEXT;
