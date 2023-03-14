CREATE TABLE app_articles (
    article_id    UUID   DEFAULT uuid_generate_v4() PRIMARY KEY,
    author  VARCHAR NOT NULL,
    title   VARCHAR NOT NULL,
    article_text VARCHAR NOT NULL,
    reads   INTEGER DEFAULT 0,
    CONSTRAINT fk_author
               FOREIGN KEY (author) REFERENCES app_user(username)
);


