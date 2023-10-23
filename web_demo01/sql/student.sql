DROP TABLE IF EXISTS student;
CREATE TABLE students(
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(20),
    age INT,
    PRIMARY KEY (id)
);
ALTER TABLE student RENAME TO students;