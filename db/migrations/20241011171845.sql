-- Create "recipes" table
CREATE TABLE `recipes` (`id` integer NULL, `name` text NOT NULL, `servings` integer NULL, `minutes` integer NULL, `description` text NULL, PRIMARY KEY (`id`));
-- Create "units" table
CREATE TABLE `units` (`id` integer NULL, `code` text NULL, `name` text NOT NULL, PRIMARY KEY (`id`));
-- Create "ingredients" table
CREATE TABLE `ingredients` (`id` integer NULL, `recipe_id` integer NOT NULL, `unit_id` integer NOT NULL, `name` text NOT NULL, `amount` real NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `0` FOREIGN KEY (`unit_id`) REFERENCES `units` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT `1` FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION);
