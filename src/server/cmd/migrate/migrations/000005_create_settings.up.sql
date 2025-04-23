CREATE TABLE IF NOT EXISTS settings (
  id bigserial PRIMARY KEY,
  user_id bigint NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
  temp_format varchar(10) NOT NULL CHECK (temp_format IN ('celsius', 'fahrenheit')),
  time_format varchar(10) NOT NULL CHECK (time_format IN ('24h', '12h')),
  speed_format varchar(10) NOT NULL CHECK (speed_format IN ('kmph', 'mph'))
);
