DROP TRIGGER IF EXISTS update_post_translations_updated_at ON posts_translations;

DROP FUNCTION IF EXISTS update_updated_at_column();

DROP TABLE IF EXISTS posts_translations CASCADE;
DROP TABLE IF EXISTS series_translations CASCADE;
DROP TABLE IF EXISTS posts CASCADE;
DROP TABLE IF EXISTS posts CASCADE;
