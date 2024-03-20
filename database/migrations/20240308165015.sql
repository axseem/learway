-- Create "deck" table
CREATE TABLE `deck` (`id` integer NULL, `title` text NOT NULL, `cards` text NOT NULL, `created_at` datetime NOT NULL DEFAULT (CURRENT_TIMESTAMP), `updated_at` datetime NOT NULL DEFAULT (CURRENT_TIMESTAMP), PRIMARY KEY (`id`));
