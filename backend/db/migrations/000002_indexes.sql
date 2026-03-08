-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_media_key ON media(key);
CREATE INDEX IF NOT EXISTS idx_links_user_id ON links(user_id);
CREATE INDEX IF NOT EXISTS idx_links_created_at ON links(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_click_events_link_id ON click_events(link_id);
CREATE INDEX IF NOT EXISTS idx_click_events_clicked_at ON click_events(clicked_at DESC);
CREATE INDEX IF NOT EXISTS idx_api_keys_user_id ON api_keys(user_id);
CREATE UNIQUE INDEX IF NOT EXISTS uq_api_keys_prefix ON api_keys(key_prefix);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS uq_api_keys_prefix;
DROP INDEX IF EXISTS idx_api_keys_user_id;
DROP INDEX IF EXISTS idx_click_events_clicked_at;
DROP INDEX IF EXISTS idx_click_events_link_id;
DROP INDEX IF EXISTS idx_links_created_at;
DROP INDEX IF EXISTS idx_links_user_id;
DROP INDEX IF EXISTS idx_media_key;
-- +goose StatementEnd
