-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_recipes" table
CREATE TABLE `new_recipes` (`id` integer NULL, `name` text NOT NULL, `servings` integer NULL, `minutes` integer NULL, `description` text NULL, `created_by` integer NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `0` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Copy rows from old table "recipes" to new temporary table "new_recipes"
INSERT INTO `new_recipes` (`id`, `name`, `servings`, `minutes`, `description`) SELECT `id`, `name`, `servings`, `minutes`, `description` FROM `recipes`;
-- Drop "recipes" table after copying rows
DROP TABLE `recipes`;
-- Rename temporary table "new_recipes" to "recipes"
ALTER TABLE `new_recipes` RENAME TO `recipes`;
-- Create "new_ingredients" table
CREATE TABLE `new_ingredients` (`id` integer NULL, `recipe_id` integer NOT NULL, `unit_id` integer NOT NULL, `name` text NOT NULL, `amount` real NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `0` FOREIGN KEY (`unit_id`) REFERENCES `units` (`id`) ON UPDATE NO ACTION ON DELETE RESTRICT, CONSTRAINT `1` FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Copy rows from old table "ingredients" to new temporary table "new_ingredients"
INSERT INTO `new_ingredients` (`id`, `recipe_id`, `unit_id`, `name`, `amount`) SELECT `id`, `recipe_id`, `unit_id`, `name`, `amount` FROM `ingredients`;
-- Drop "ingredients" table after copying rows
DROP TABLE `ingredients`;
-- Rename temporary table "new_ingredients" to "ingredients"
ALTER TABLE `new_ingredients` RENAME TO `ingredients`;
-- Create "users" table
CREATE TABLE `users` (`id` integer NULL, `email` text NOT NULL, `password_hash` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "users_email" to table: "users"
CREATE UNIQUE INDEX `users_email` ON `users` (`email`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
