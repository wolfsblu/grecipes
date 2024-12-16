-- Create "recipes" table
CREATE TABLE `recipes` (`id` integer NULL, `name` text NOT NULL, `servings` integer NOT NULL, `minutes` integer NOT NULL, `description` text NOT NULL, `created_by` integer NOT NULL, `created_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP), PRIMARY KEY (`id`), CONSTRAINT `0` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "units" table
CREATE TABLE `units` (`id` integer NULL, `code` text NULL, `name` text NOT NULL, PRIMARY KEY (`id`));
-- Create "ingredients" table
CREATE TABLE `ingredients` (`id` integer NULL, `recipe_id` integer NOT NULL, `unit_id` integer NOT NULL, `name` text NOT NULL, `amount` real NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `0` FOREIGN KEY (`unit_id`) REFERENCES `units` (`id`) ON UPDATE NO ACTION ON DELETE RESTRICT, CONSTRAINT `1` FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "users" table
CREATE TABLE `users` (`id` integer NULL, `email` text NOT NULL, `password_hash` text NOT NULL, `created_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP), PRIMARY KEY (`id`));
-- Create index "users_email" to table: "users"
CREATE UNIQUE INDEX `users_email` ON `users` (`email`);
-- Create "user_awards" table
CREATE TABLE `user_awards` (`user_id` integer NOT NULL, `recipe_id` integer NOT NULL, `award_id` integer NOT NULL, `awarded_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP), CONSTRAINT `0` FOREIGN KEY (`award_id`) REFERENCES `awards` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT `1` FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT `2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "user_reputation" table
CREATE TABLE `user_reputation` (`user_id` integer NOT NULL, `recipe_id` integer NOT NULL, `action_id` integer NOT NULL, `awarded_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP), CONSTRAINT `0` FOREIGN KEY (`action_id`) REFERENCES `actions` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT `1` FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT `2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "recipe_votes" table
CREATE TABLE `recipe_votes` (`recipe_id` integer NOT NULL, `user_id` integer NOT NULL, `vote` integer NOT NULL DEFAULT 1, `voted_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP), CONSTRAINT `0` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT `1` FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "actions" table
CREATE TABLE `actions` (`id` integer NULL, `name` text NOT NULL, `points` integer NOT NULL DEFAULT 0, PRIMARY KEY (`id`));
-- Create "awards" table
CREATE TABLE `awards` (`id` integer NULL, `name` text NOT NULL, PRIMARY KEY (`id`));
