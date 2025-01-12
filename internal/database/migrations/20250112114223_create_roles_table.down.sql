-- Drop the user_roles table
DROP TABLE IF EXISTS user_roles;

-- Delete the seed data from roles table
DELETE FROM roles
WHERE name IN ('admin', 'student', 'lecturer', 'laboratory_assistant');

-- Drop the roles table
DROP TABLE IF EXISTS roles;
