CREATE TABLE IF NOT EXISTS student_registration (
    id_student_registration SERIAL PRIMARY KEY,
    student_id INT NOT NULL,
    practicum_id INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE, 
    FOREIGN KEY (practicum_id) REFERENCES practicums(id_practicum) ON DELETE CASCADE
);


CREATE INDEX IF NOT EXISTS idx_student_registration_student_id ON student_registration(student_id);
CREATE INDEX IF NOT EXISTS idx_student_registration_practicum_id ON student_registration(practicum_id);
