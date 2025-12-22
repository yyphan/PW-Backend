ALTER TABLE posts ALTER COLUMN series_id SET NOT NULL;
ALTER TABLE posts ALTER COLUMN idx_in_series SET NOT NULL;

ALTER TABLE posts DROP CONSTRAINT IF EXISTS posts_series_id_fkey;

ALTER TABLE posts 
    ADD CONSTRAINT posts_series_id_fkey 
    FOREIGN KEY (series_id) 
    REFERENCES series(id) 
    ON DELETE CASCADE;
