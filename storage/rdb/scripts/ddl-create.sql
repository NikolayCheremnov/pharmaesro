CREATE TABLE medicament (
    uuid UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    dosage FLOAT CHECK (Dosage > 0),
    dosage_type INT NOT NULL
);
