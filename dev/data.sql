CREATE TABLE IF NOT EXISTS actors (
    id SERIAL PRIMARY KEY,
    display_name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS articles (
    id SERIAL PRIMARY KEY,
    title VARCHAR(250) NOT NULL,
    content TEXT NOT NULL,
    status VARCHAR(40) NOT NULL,
    actor_id INT NOT NULL,
    FOREIGN KEY (actor_id) REFERENCES actors (id)
);

INSERT INTO actors (display_name) VALUES ('default_user');
INSERT INTO articles (title, content, status, actor_id) VALUES ('Default Title', 'Default content.', 'published', 1);
