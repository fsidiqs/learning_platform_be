package queryutil

const (
	INSERT_NEW_USER = "INSERT INTO users" +
		" (name, email, password, role_id)" +
		" VALUES(?, ?, ?, ?)"
	user                            = "id, name, email, created_at, updated_at, deleted_at"
	user_with_pw                    = "id, name, email, password, created_at, updated_at, deleted_at"
	GET_ALL_USER                    = "SELECT " + user + " FROM users"
	GET_USER_BY_ID                  = "SELECT " + user + " FROM users where id = ?"
	GET_USER_BY_EMAIL               = "SELECT " + user + " FROM users where email = ?"
	GET_USER_BY_EMAIL_WITH_PASSWORD = "SELECT " + user_with_pw + " FROM users where email = ?"

	UPDATE_USER_BY_ID = "UPDATE users" +
		" SET name = ?, email = ?, updated_at=NOW()" +
		" WHERE id = ?"
	UPDATE_USER_NAME_BY_ID = "UPDATE users" +
		" SET name = ?" +
		" WHERE id = ?"
	DEACTIVATE_USER = "UPDATE users SET updated_at=NOW(), deleted_at=NOW() WHERE id=?"
	ACTIVATE_USER   = "UPDATE users SET updated_at=NOW(), deleted_at=NULL WHERE id=?"

	DELETE_USER_BY_ID = "DELETE FROM users WHERE id=?"
)

const (
	INSERT_NEW_STUDENT = "INSERT INTO students" +
		" (user_id)" +
		" VALUES(?)"

	student                = "id, user_id"
	GET_STUDENT_BY_ID      = "SELECT " + student + " FROM students where id = ?"
	GET_STUDENT_BY_USER_ID = "SELECT " + student + " FROM students where user_id = ?"
)

const (
	INSERT_NEW_COURSE = "INSERT INTO courses" +
		" (title, description, price, image_url, created_at, updated_at, author_id)" +
		" VALUES(?, ?, ?, ?, ?, ?, ?)"
	course          = "id, author_id,title, description, price, image_url, created_at, updated_at, deleted_at"
	GET_ALL_COURSES = "SELECT " + course + " FROM courses"
)

const (
	GET_ALL_COURSES_AGGR_STUDENT_COURSE_AGGR_STUDENT_BY_USER_ID = `
SELECT
	c.id,
	c.author_id,
	c.title,
	c.description,
	c.price,
	c.image_url,
	c.created_at,
	c.updated_at,
	c.deleted_at
FROM
	students AS s
INNER JOIN students_courses AS sc ON
	s.id = sc.student_id
INNER JOIN courses c ON
	sc.course_id = c.id
WHERE
	s.user_id = ?
`
)

const (
	INSERT_NEW_VIDEO_LECTURE = "INSERT INTO video_lectures" +
		" (title,video_url, created_at, updated_at, course_id)" +
		" VALUES(?, ?, ?, ?, ?)"
	GET_ALL_VIDEO_LECTURES_BY_COURSE_ID = `
	SELECT
		vl.*
	FROM
		courses c
	INNER JOIN video_lectures vl ON
		vl.course_id = c.id
	WHERE c.id = ?
	`
)

const (
	INSERT_STUDENT_COURSE = "INSERT INTO students_courses" +
		" (student_id, course_id)" +
		" VALUES(?, ?)"
)
