CREATE TABLE IF NOT EXISTS practicum_module_content (
    id_content SERIAL PRIMARY KEY,
    id_module INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL, -- Use TEXT for potentially large content
    sequence INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_module) REFERENCES practicum_modules(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_practicum_module_content_id_module ON practicum_module_content(id_module);
