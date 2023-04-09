CREATE TABLE app_likes
(
    article_id UUID    NOT NULL,
    username   VARCHAR NOT NULL,
    CONSTRAINT fk_article_id
        FOREIGN KEY (article_id) REFERENCES app_articles (article_id),
    CONSTRAINT fk_username
        FOREIGN KEY (username) REFERENCES app_user (username),
    CONSTRAINT pk_article_id_username
        PRIMARY KEY (article_id, username)
);