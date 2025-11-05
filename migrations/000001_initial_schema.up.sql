
-- A series consists of posts
CREATE TABLE IF NOT EXISTS series (
    id SERIAL PRIMARY KEY,
    bg_url VARCHAR(2048)
);

-- A post may or may not belong to one of the series, if it does, it has an index in its series
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    series_id INT,
    idx_in_series INT,

    UNIQUE (series_id, idx_in_series),

    FOREIGN KEY (series_id) 
        REFERENCES series(id) 
        ON DELETE SET NULL
);

-- A series (title...) is available in multiple languages
CREATE TABLE IF NOT EXISTS series_translations (
    series_id INT NOT NULL,
    language_code VARCHAR(2) NOT NULL,
    title VARCHAR(255) NOT NULL,

    PRIMARY KEY (series_id, language_code),

    FOREIGN KEY (series_id) 
        REFERENCES series(id) 
        ON DELETE CASCADE
);

-- A post is available in multiple languages
CREATE TABLE IF NOT EXISTS posts_translations (
    post_id INT NOT NULL,
    language_code VARCHAR(2) NOT NULL,
    title VARCHAR(255) NOT NULL,
    markdown_file_path VARCHAR(2048) NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE,

    PRIMARY KEY (post_id, language_code),

    FOREIGN KEY (post_id) 
        REFERENCES posts(id) 
        ON DELETE CASCADE
);

-- function for updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- trigger for the posts_translations table
CREATE TRIGGER update_post_translations_updated_at
BEFORE UPDATE ON posts_translations
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
