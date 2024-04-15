CREATE TABLE notes (
                       id SERIAL NOT NULL PRIMARY KEY,
                       user_id INT NOT NULL,
                       title VARCHAR NOT NULL,
                       content VARCHAR,
                       date_created TIMESTAMP,
                       date_modified TIMESTAMP
);

-- DROP TABLE notes;