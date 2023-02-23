ALTER TABLE app_user ADD CONSTRAINT unique_email UNIQUE(email);
ALTER TABLE app_user ADD CONSTRAINT unique_username UNIQUE(username);