CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL    
);

CREATE TABLE pet (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    category INT REFERENCES category(id)
);

INSERT INTO category(name) VALUES ('Dog'), ('Cat'), ('Hamster'), ('Rabbit');

INSERT INTO pet(name, category) VALUES 
    ('Rafale', 1), 
    ('Rio', 1), 
    ('Mia', 2), 
    ('Laverne', 3);