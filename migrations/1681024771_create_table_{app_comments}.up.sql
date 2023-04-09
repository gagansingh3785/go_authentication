CREATE TABLE app_comments
(
    username      VARCHAR     NOT NULL,
    article_id    UUID        NOT NULL,
    comment_id    UUID        DEFAULT uuid_generate_v4() PRIMARY KEY,
    content       VARCHAR     NOT NULL,
    creation_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT fk_username
        FOREIGN KEY (username) REFERENCES app_user(username),
    CONSTRAINT fk_article_id
        FOREIGN KEY (article_id) REFERENCES app_articles(article_id)
);