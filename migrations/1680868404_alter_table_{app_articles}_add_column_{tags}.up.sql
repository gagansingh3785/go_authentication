ALTER TABLE app_articles ADD COLUMN IF NOT EXISTS tags Text[] NOT NULL DEFAULT '{}'::text[];