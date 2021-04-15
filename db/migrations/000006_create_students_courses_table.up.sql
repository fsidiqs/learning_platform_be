CREATE TABLE IF NOT EXISTS `students_courses` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT,
    `student_id` BIGINT UNSIGNED NOT NULL,
    `course_id` BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY(`id`),
    CONSTRAINT sc_student_id FOREIGN KEY(student_id) references students(id),
    CONSTRAINT sc_course_id FOREIGN KEY(course_id) references courses(id)
);

