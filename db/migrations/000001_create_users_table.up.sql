CREATE TABLE IF NOT EXISTS `users` (
  `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) UNIQUE NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP NULL
);

INSERT INTO
  `users` (`name`, `email`, `password`)
VALUES
  (
    "fajarkidut",
    "aspisetan@gmail.com",
    "adit1234"
  );