CREATE TABLE
    IF NOT EXISTS user_practicum_progress (
        id SERIAL PRIMARY KEY,
        id_user INT NOT NULL,
        id_practicum INT NOT NULL,
        progress NUMERIC(5, 2) CHECK (
            progress >= 0
            AND progress <= 100
        ) DEFAULT 0,
        completed_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT NULL,
            last_updated_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (id_user) REFERENCES users (id_user) ON DELETE CASCADE,
            FOREIGN KEY (id_practicum) REFERENCES practicums (id_practicum) ON DELETE CASCADE
    );

CREATE INDEX IF NOT EXISTS idx_user_practicum_progress_user ON user_practicum_progress (id_user);

CREATE INDEX IF NOT EXISTS idx_user_practicum_progress_practicum ON user_practicum_progress (id_practicum);