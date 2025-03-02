CREATE TABLE
    IF NOT EXISTS practicum_class (
        id_practicum_class SERIAL PRIMARY KEY,
        practicum_id INT NOT NULL,
        name VARCHAR(255) NOT NULL,
        quota INT NOT NULL,
        day VARCHAR(255) NOT NULL,
        time TIME NOT NULL,
        created_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (practicum_id) REFERENCES practicums (id_practicum) ON DELETE CASCADE
    );

CREATE INDEX IF NOT EXISTS idx_practicum_class_practicum_id ON practicum_class (practicum_id);