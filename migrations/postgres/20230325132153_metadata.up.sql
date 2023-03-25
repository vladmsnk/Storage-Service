CREATE TABLE IF NOT EXISTS material_info (
    material_id VARCHAR(255) PRIMARY KEY,
    author VARCHAR(255),
    title VARCHAR(255) UNIQUE NOT NULL,
    file_type VARCHAR(255),
    file_link VARCHAR(255)
);