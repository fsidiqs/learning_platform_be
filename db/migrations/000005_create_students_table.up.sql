CREATE TABLE IF NOT EXISTS `students` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT students_user_id FOREIGN KEY(user_id) references users(id)
);

INSERT INTO
    `students` (`user_id`)
VALUES
    (2);