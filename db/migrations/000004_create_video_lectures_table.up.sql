CREATE TABLE IF NOT EXISTS `video_lectures` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `video_url` VARCHAR(255) NOT NULL,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    `course_id` BIGINT UNSIGNED NOT NULL,
    CONSTRAINT lecture_videos_course_id FOREIGN KEY(course_id) REFERENCES courses(id)
);