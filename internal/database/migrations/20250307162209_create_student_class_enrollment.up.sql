CREATE TABLE IF NOT EXISTS student_class_enrollment (
    id SERIAL PRIMARY KEY,
    class_id INT NOT NULL,
    student_id INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (class_id) REFERENCES practicum_class (id_practicum_class) ON DELETE CASCADE,
    FOREIGN KEY (student_id) REFERENCES students (id) ON DELETE CASCADE
);
