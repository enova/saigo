CREATE TABLE people (
    person_id SERIAL PRIMARY KEY,
    name TEXT,
    ssn INTEGER UNIQUE
);
