CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL
);

CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id),
    name VARCHAR(255) NOT NULL,
    checked BOOLEAN NOT NULL
);

INSERT INTO users (username) VALUES ('testuser');
INSERT INTO users (username) VALUES ('admin');

INSERT INTO todos (user_id, name, checked) VALUES (1, 'Do the homework', false);
INSERT INTO todos (user_id, name, checked) VALUES (1, 'Test Todo 1', true);
INSERT INTO todos (user_id, name, checked) VALUES (2, 'Test Todo 2', true);
INSERT INTO todos (user_id, name, checked) VALUES (2, 'Test Todo 3', false);