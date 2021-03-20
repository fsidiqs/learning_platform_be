CREATE TABLE IF NOT EXISTS `authors` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT authors_user_id FOREIGN KEY(user_id) references users(id)
);

INSERT INTO
    `authors` (`user_id`)
VALUES
    (1);