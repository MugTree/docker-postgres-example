--  Simple testing database and dummy data
CREATE DATABASE docker_db_test;
CREATE TABLE "user" (user_name text primary key, password text);
INSERT INTO TABLE "user" (user_name, password)
VALUES ("Matt", "123456");