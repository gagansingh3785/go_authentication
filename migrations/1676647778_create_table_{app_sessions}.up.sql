CREATE TABLE app_sessions (
    user_id     UUID,
    session_id  VARCHAR,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES app_user(user_id)
);