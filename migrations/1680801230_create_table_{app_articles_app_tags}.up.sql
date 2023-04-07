CREATE TABLE app_articles_app_tags (
    article_id  UUID NOT NULL,
    tag_id      UUID NOT NULL,
    CONSTRAINT  fk_article_id
        FOREIGN KEY (article_id) REFERENCES app_articles(article_id),
    CONSTRAINT  fk_tag_id
        FOREIGN KEY (tag_id)    REFERENCES app_tags(tag_id),
    CONSTRAINT  pk_article_id_tag_id
        PRIMARY KEY (article_id, tag_id)
);