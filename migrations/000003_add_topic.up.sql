ALTER TABLE series ADD COLUMN topic VARCHAR(255);

CREATE INDEX IF NOT EXISTS idx_series_topic ON series (topic);
