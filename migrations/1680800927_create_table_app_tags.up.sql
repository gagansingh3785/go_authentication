CREATE TABLE app_tags (
    tag_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

INSERT INTO app_tags  (name) VALUES ('sports');
INSERT INTO app_tags (name) VALUES ('opinion');
INSERT INTO app_tags  (name) VALUES ('politics');
INSERT INTO app_tags (name) VALUES ('news');
INSERT INTO app_tags  (name) VALUES ('entertainment');
INSERT INTO app_tags (name) VALUES ('nature');
INSERT INTO app_tags (name) VALUES ('sex');
INSERT INTO app_tags (name) VALUES ('tech');
INSERT INTO app_tags (name) VALUES ('space');
INSERT INTO app_tags (name) VALUES ('other');