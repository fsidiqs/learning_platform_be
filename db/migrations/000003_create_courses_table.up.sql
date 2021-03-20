CREATE TABLE IF NOT EXISTS `courses` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `title` VARCHAR(255) NOT NULL,
    `description` TEXT NOT NULL,
    `price` BIGINT DEFAULT 0,
    `image_url` VARCHAR(255) NOT NULL,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    `author_id` BIGINT UNSIGNED NOT NULL,
    CONSTRAINT courses_author_id FOREIGN KEY(author_id) REFERENCES authors(id)
);

INSERT INTO
    `courses` (`title`, `description`, `price`, `image_url`, `author_id`)
VALUES
    (
        "OTTOMAN CUISINE, TURKISH COOKING CLASS",
        "<p>Turkish cuisine one of the world's great cuisines. It reflects the long history of this land...</p><p>A great variety of mouth watering dishes in Turkish cuisine which is mostly the heritage of Ottoman cuisine. It is the mixture and refinement of Central Asian, Middle Eastern and Balkan cuisines. Therefore it is impossible to fit Turkish cuisine into a short list.</p><p>Anatolia is blessed with varied climate which allows the country to get almost everything on its land. Turkey is one of the few countries in the world that has been self sustaining, producing all its own food. Turkish cuisine traditionally is NOT spicy, except in the southeast part of the country, where preparations can reflect a hot Middle Eastern(Arabic) food influence. But now in evertwhere you are able to find restaurants that sell that kind of spicy dishes</p><p>Turks have a big diversity of vegetables and of course this reflects on the dishes. One very important detail about vegetable dishes is whether they have meat in them or not.</p>",
        159000,
        "https://i.udemycdn.com/course/480x270/2702904_7dd8.jpg",
        1
    );