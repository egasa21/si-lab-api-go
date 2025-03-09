CREATE TABLE IF NOT EXISTS user_practicum_checkpoint (
    id SERIAL PRIMARY KEY,
    id_user INT NOT NULL,
    id_practicum INT NOT NULL,
    id_module INT NOT NULL,
    id_content INT NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_user) REFERENCES users(id_user) ON DELETE CASCADE,
    FOREIGN KEY (id_practicum) REFERENCES practicums(id_practicum) ON DELETE CASCADE,
    FOREIGN KEY (id_module) REFERENCES practicum_modules(id) ON DELETE CASCADE,
    FOREIGN KEY (id_content) REFERENCES practicum_module_content(id_content) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_user_practicum_checkpoint_user ON user_practicum_checkpoint(id_user);
CREATE INDEX IF NOT EXISTS idx_user_practicum_checkpoint_practicum ON user_practicum_checkpoint(id_practicum);
CREATE INDEX IF NOT EXISTS idx_user_practicum_checkpoint_module ON user_practicum_checkpoint(id_module);
CREATE INDEX IF NOT EXISTS idx_user_practicum_checkpoint_content ON user_practicum_checkpoint(id_content);
