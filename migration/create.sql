CREATE TABLE notes (
                       id SERIAL PRIMARY KEY,
                       user_id INT,
                       header VARCHAR,
                       content VARCHAR,
                       date_created TIMESTAMP,
                       date_modified TIMESTAMP
);
