CREATE TABLE recipes
(
    id          INTEGER PRIMARY KEY,
    name        TEXT    NOT NULL,
    servings    INTEGER NULL,
    minutes     INTEGER NULL,
    description TEXT NULL,
    created_by  INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE units
(
    id   INTEGER PRIMARY KEY,
    code TEXT NULL,
    name TEXT NOT NULL
);

CREATE TABLE ingredients
(
    id        INTEGER PRIMARY KEY,
    recipe_id INTEGER NOT NULL REFERENCES recipes (id) ON DELETE CASCADE,
    unit_id   INTEGER NOT NULL REFERENCES units (id) ON DELETE RESTRICT,
    name      TEXT    NOT NULL,
    amount    REAL    NOT NULL
);

CREATE TABLE users
(
    id            INTEGER PRIMARY KEY,
    email         TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL
);