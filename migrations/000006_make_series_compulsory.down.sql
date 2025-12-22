ALTER TABLE posts DROP CONSTRAINT IF EXISTS fk_posts_series;

ALTER TABLE posts ALTER COLUMN series_id DROP NOT NULL;
ALTER TABLE posts ALTER COLUMN idx_in_series DROP NOT NULL；

ALTER TABLE posts 
    ADD CONSTRAINT fk_posts_series 
    FOREIGN KEY (series_id) 
    REFERENCES series(id) 
    ON DELETE SET NULL;