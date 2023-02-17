ALTER TABLE app_user DROP COLUMN salt;
ALTER TABLE app_user ADD COLUMN role INTEGER;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; -- This avails the uuid_generate_v4() function for uuid generation
ALTER TABLE app_user ADD COLUMN user_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY;