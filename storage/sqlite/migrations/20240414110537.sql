-- Create "deck" table
CREATE TABLE `deck` (`id` text NULL, `user_id` text NOT NULL, `title` text NOT NULL, `cards` blob NOT NULL, `created_at` datetime NOT NULL DEFAULT (CURRENT_TIMESTAMP), `updated_at` datetime NOT NULL DEFAULT (CURRENT_TIMESTAMP), PRIMARY KEY (`id`), CONSTRAINT `0` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "session" table
CREATE TABLE `session` (`id` text NULL, `user_id` text NOT NULL, `fingerprint` blob NOT NULL, `ip` blob NOT NULL, `expires_at` datetime NOT NULL, `created_at` datetime NOT NULL DEFAULT (CURRENT_TIMESTAMP), PRIMARY KEY (`id`), CONSTRAINT `0` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "user" table
CREATE TABLE `user` (`id` text NULL, `username` text NOT NULL, `email` text NOT NULL, `password` blob NOT NULL, `created_at` datetime NOT NULL DEFAULT (CURRENT_TIMESTAMP), PRIMARY KEY (`id`));
-- Create index "user_username" to table: "user"
CREATE UNIQUE INDEX `user_username` ON `user` (`username`);
-- Create index "user_email" to table: "user"
CREATE UNIQUE INDEX `user_email` ON `user` (`email`);
