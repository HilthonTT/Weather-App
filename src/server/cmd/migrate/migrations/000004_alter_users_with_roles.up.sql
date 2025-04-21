-- Add the role_id column with a default value and foreign key constraint
ALTER TABLE IF EXISTS users
ADD COLUMN role_id INT REFERENCES roles(id) DEFAULT 1;

-- Set role_id for existing users to the ID of the 'user' role
UPDATE users
SET role_id = (
  SELECT id
  FROM roles
  WHERE name = 'user'
);

-- Remove the default value for role_id
ALTER TABLE users
ALTER COLUMN role_id DROP DEFAULT;

-- Make role_id a NOT NULL column
ALTER TABLE users
ALTER COLUMN role_id SET NOT NULL;
