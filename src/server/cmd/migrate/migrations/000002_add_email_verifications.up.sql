CREATE TABLE IF NOT EXISTS email_verifications (
  token bytea PRIMARY KEY,
  user_id bigint NOT NULL
)

