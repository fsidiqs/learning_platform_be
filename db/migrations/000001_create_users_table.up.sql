CREATE TABLE IF NOT EXISTS `users` (
  `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) UNIQUE NOT NULL,
  `role_id` TINYINT UNSIGNED NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP NULL
);

INSERT INTO
  `users` (`name`, `email`, `password`, `role_id`)
VALUES
  (
    "admin",
    "admin@gmail.com",
    "$2y$10$50KlerDENnD3OUu7JDMMMOHUD2Dbc.1JkJuyR5igCTvzGOBeTDDRe",
    1 
  ),
  (
    "user",
    "user@gmail.com",
    "$2y$10$50KlerDENnD3OUu7JDMMMOHUD2Dbc.1JkJuyR5igCTvzGOBeTDDRe",
    2 
  );