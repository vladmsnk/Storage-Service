CREATE TABLE IF NOT EXISTS material_info (
    material_id uuid PRIMARY KEY,
    author      VARCHAR(255),
    title       VARCHAR(255),
    file_type   VARCHAR(255),
    file_link   VARCHAR(2048)
);