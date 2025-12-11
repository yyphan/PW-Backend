ALTER TABLE posts ADD COLUMN post_slug VARCHAR(255);
ALTER TABLE series ADD COLUMN series_slug VARCHAR(255) UNIQUE;

ALTER TABLE posts ADD CONSTRAINT uq_series_slug UNIQUE (series_id, post_slug);
