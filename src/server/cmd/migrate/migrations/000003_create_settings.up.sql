DO $$ BEGIN
    CREATE TYPE temperature_unit AS ENUM ('Celsius', 'Fahrenheit');
EXCEPTION
    WHEN duplicate_object THEN NULL;
END $$;

DO $$ BEGIN
    CREATE TYPE wind_speed_unit AS ENUM ('Km/h', 'Miles/h');
EXCEPTION
    WHEN duplicate_object THEN NULL;
END $$;

DO $$ BEGIN
    CREATE TYPE time_format AS ENUM ('12h', '24h');
EXCEPTION
    WHEN duplicate_object THEN NULL;
END $$;

CREATE TABLE IF NOT EXISTS settings (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    temperature temperature_unit NOT NULL DEFAULT 'Celsius',
    wind_speed wind_speed_unit NOT NULL DEFAULT 'Km/h',
    time_display time_format NOT NULL DEFAULT '24h'
);
