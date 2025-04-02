CREATE EXTENSION IF NOT EXISTS citext;

-- Create the extension and indexes for full-text search
-- Check article: https://niallburkley.com/blog/index-columns-for-like-in-postgres/
CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE IF NOT EXISTS users(
  id bigserial PRIMARY KEY,
  email citext UNIQUE NOT NULL,
  username varchar(255) UNIQUE NOT NULL,
  password bytea NOT NULL,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_users_username ON users (username);
