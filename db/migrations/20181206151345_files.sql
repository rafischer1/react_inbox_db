-- +goose Up
CREATE TABLE files (
 id serial PRIMARY KEY,
 coolNotCool boolean,
 body text UNIQUE NOT NULL
);

INSERT INTO files
  (id, coolNotCool, body)
VALUES
  (1, true, 'I am super cool'), (2, false, 'I am not cool'), (3, true, 'Go is a hard language to learn');

-- +goose Down
DROP TABLE files;