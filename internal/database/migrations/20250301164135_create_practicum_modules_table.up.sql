CREATE TABLE IF NOT EXISTS practicum_modules (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    practicum_id INT NOT NULL,
    FOREIGN KEY (practicum_id) REFERENCES practicums(id_practicum) ON DELETE CASCADE 
);

CREATE INDEX IF NOT EXISTS idx_practicum_modules_practicum_id ON practicum_modules(practicum_id);
