CREATE TABLE recipes
(
    id          INTEGER PRIMARY KEY,
    name        TEXT NOT NULL,
    servings    INTEGER NULL,
    minutes     INTEGER NULL,
    description TEXT NULL
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
    recipe_id INTEGER NOT NULL REFERENCES recipes (id),
    unit_id   INTEGER NOT NULL REFERENCES units (id),
    name      TEXT    NOT NULL,
    amount    REAL    NOT NULL
);