DROP TABLE IF EXISTS person_thought;

DROP TABLE IF EXISTS thought;

DROP TABLE IF EXISTS person;

CREATE TABLE person (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE thought (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE person_thought (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    person_id INTEGER NOT NULL,
    thought_id INTEGER NOT NULL,
    FOREIGN KEY(person_id) REFERENCES person(id),
    FOREIGN KEY(thought_id) REFERENCES thought(id),
    UNIQUE (person_id, thought_id) ON CONFLICT IGNORE
);

INSERT INTO person
(name)
VALUES
("Alice"),
("Bob");

INSERT INTO thought
(name)
VALUES
("gratitude"),
("indifference");

INSERT INTO person_thought
(person_id, thought_id)
VALUES
(1, 1),
(2, 1),
(2, 2);
