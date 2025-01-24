-- Create the roles table
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);

-- Seed predefined roles into the roles table
-- INSERT IGNORE INTO roles (name) VALUES
-- ('admin'),
-- ('student'),
-- ('lecturer'),
-- ('laboratory_assistant');

-- Create the user_roles table
CREATE TABLE IF NOT EXISTS user_roles (
    id_user INT NOT NULL,
    id_role INT NOT NULL,
    PRIMARY KEY (id_user, id_role),
    FOREIGN KEY (id_user) REFERENCES users(id_user) ON DELETE CASCADE,
    FOREIGN KEY (id_role) REFERENCES roles(id) ON DELETE CASCADE
);