-- Step 1: Add column with a default value (to avoid NOT NULL issues on existing rows)
ALTER TABLE IF EXISTS users
ADD COLUMN role_id INT REFERENCES roles(id) DEFAULT 1;

-- Step 2: Update existing rows with the actual 'user' role ID
UPDATE users
SET role_id = (
    SELECT id
    FROM roles
    WHERE name = 'user'
    LIMIT 1 -- just in case multiple exist
);

-- Step 3: Remove the default now that the column has been populated
ALTER TABLE users
ALTER COLUMN role_id DROP DEFAULT;

-- Step 4: Make the column NOT NULL
ALTER TABLE users
ALTER COLUMN role_id SET NOT NULL;
