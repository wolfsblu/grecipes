CREATE TABLE recipes
(
    id          INTEGER PRIMARY KEY,
    name        TEXT      NOT NULL,
    servings    INTEGER   NOT NULL,
    minutes     INTEGER   NOT NULL,
    description TEXT      NOT NULL,
    created_by  INTEGER   NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
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
    email         TEXT      NOT NULL UNIQUE,
    password_hash TEXT      NOT NULL,
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE password_resets
(
    user_id    INTEGER PRIMARY KEY REFERENCES users (id) ON DELETE CASCADE,
    token      TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_awards
(
    user_id    INTEGER   NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    recipe_id  INTEGER   NOT NULL REFERENCES recipes (id) ON DELETE CASCADE,
    award_id   INTEGER   NOT NULL REFERENCES awards (id) ON DELETE CASCADE,
    awarded_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_reputation
(
    user_id    INTEGER   NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    recipe_id  INTEGER   NOT NULL REFERENCES recipes (id) ON DELETE CASCADE,
    action_id  INTEGER   NOT NULL REFERENCES actions (id) ON DELETE CASCADE,
    awarded_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE recipe_votes
(
    recipe_id INTEGER   NOT NULL REFERENCES recipes (id) ON DELETE CASCADE,
    user_id   INTEGER   NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    vote      INTEGER   NOT NULL DEFAULT 1,
    voted_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE actions
(
    id     INTEGER PRIMARY KEY,
    name   TEXT    NOT NULL,
    points INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE awards
(
    id   INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);