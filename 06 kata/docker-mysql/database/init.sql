CREATE DATABASE IF NOT EXISTS todo;

USE todo;

CREATE TABLE IF NOT EXISTS todoer (
  id SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
  username VARCHAR(64) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS todo (
  id SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
  description TEXT NOT NULL,
  creator SMALLINT UNSIGNED NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (creator) REFERENCES todoer(id) ON DELETE CASCADE
);

INSERT INTO todoer(username)
VALUES ("Ziwei"),
       ("Mingwei");

INSERT INTO todo(description, creator)
VALUES ("Create goroutine examples for workshop", 1),
       ("Create reading material for workshop", 1),
       ("Update timesheet", 1),
       ("File transport claim", 1),
       ("Follow up on interview test questions with team", 2);