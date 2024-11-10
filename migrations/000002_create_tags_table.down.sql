-- Drop indexes first
DROP INDEX IF EXISTS idx_tags_name;
DROP INDEX IF EXISTS idx_url_tags_tag_id;
DROP INDEX IF EXISTS idx_url_tags_url_id;

-- Then tables
DROP TABLE IF EXISTS url_tags;
DROP TABLE IF EXISTS tags;
