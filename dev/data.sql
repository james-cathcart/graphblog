CREATE DATABASE IF NOT EXISTS blogsite;
USE blogsite;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    display_name VARCHAR(100) UNIQUE NOT NULL,
);

CREATE TABLE IF NOT EXISTS articles (
    id SERIAL PRIMARY KEY,
    title VARCHAR(250) NOT NULL,
    content TEXT NOT NULL,
    status VARCHAR(40) NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES (users.id)
);

INSERT INTO users (display_name) VALUES ('default_user');
INSERT INTO articles (title, content, status, user_id) VALUES ('Default Title', 'Default content.', 'published', 1);
