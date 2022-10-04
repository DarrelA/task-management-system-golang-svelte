-- CREATE DATABASE AND TABLES

CREATE DATABASE IF NOT EXISTS c3_database;
USE c3_database;

CREATE TABLE IF NOT EXISTS accounts (
  username varchar(50) NOT NULL,
  password varchar(255) NOT NULL,
  email varchar(100) DEFAULT NULL,
  admin_privilege tinyint DEFAULT '0',   -- boolean ('0' - user, '1' - admin)
  user_group varchar(255) DEFAULT NULL,
  status enum('Active','Inactive') NOT NULL DEFAULT 'Active',
  timestamp datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (username),
  UNIQUE KEY username_UNIQUE (username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE IF NOT EXISTS groupnames (
  user_group varchar(255) NOT NULL,
  PRIMARY KEY (user_group)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE IF NOT EXISTS usergroup (
  username varchar(255) NOT NULL,
  user_group varchar(255) NOT NULL,
  PRIMARY KEY (user_group,username),
  KEY username_idx (username),
  KEY groupname_idx (user_group),
  CONSTRAINT username FOREIGN KEY (username) REFERENCES accounts (username),
  CONSTRAINT user_group FOREIGN KEY (user_group) REFERENCES groupnames (user_group)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE IF NOT EXISTS application (
  app_acronym varchar(255) NOT NULL,
  app_description longtext,
  app_Rnum int DEFAULT NULL,
  app_startDate date DEFAULT NULL,
  app_endDate date DEFAULT NULL,
  app_permitCreate varchar(255) DEFAULT NULL,
  app_permitOpen varchar(255) DEFAULT NULL,
  app_permitToDo varchar(255) DEFAULT NULL,
  app_permitDoing varchar(255) DEFAULT NULL,
  app_permitDone varchar(255) DEFAULT NULL,
  app_createdDate datetime DEFAULT CURRENT_TIMESTAMP,   -- insert with `now()`
  PRIMARY KEY (app_acronym)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE IF NOT EXISTS plan (
  plan_mvp_name varchar(255) NOT NULL,
  plan_app_acronym varchar(255) DEFAULT NULL,
  plan_color varchar(7) DEFAULT NULL,
  plan_startDate datetime NOT NULL,  -- "sort"
  plan_endDate datetime NOT NULL,
  PRIMARY KEY (plan_mvp_name),
  KEY plan_app_acronym (plan_app_acronym),   -- index (find rows with specific column value quickly)
  KEY plan_color (plan_color),
  CONSTRAINT plan_app_acronym FOREIGN KEY (plan_app_acronym) REFERENCES application (app_acronym)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE IF NOT EXISTS task (
  task_app_acronym varchar(255) NOT NULL,
  task_id varchar(255) DEFAULT NULL,
  task_name varchar(255) NOT NULL,
  task_description longtext,
  task_notes longtext,
  task_plan varchar(255) DEFAULT NULL,
  task_color varchar(7) DEFAULT NULL,
  task_state enum('Open','ToDo','Doing','Done','Closed') DEFAULT 'Open', -- enum --> options
  task_creator varchar(255) DEFAULT NULL,
  task_owner varchar(255) DEFAULT NULL,
  task_createDate datetime DEFAULT NULL,
  PRIMARY KEY (task_app_acronym,task_name),
  KEY task_name (task_name),
  KEY task_plan (task_plan),
  KEY task_color_idx (task_color),
  CONSTRAINT task_app_acronym FOREIGN KEY (task_app_acronym) REFERENCES application (app_acronym),
  CONSTRAINT task_plan FOREIGN KEY (task_plan) REFERENCES plan (plan_mvp_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE IF NOT EXISTS task_notes (
  task_notes_id int NOT NULL AUTO_INCREMENT,
  task_name varchar(255) NOT NULL,
  task_note longtext,
  task_owner varchar(255) NOT NULL,
  task_state enum('Open','ToDo','Doing','Done','Closed'),
  task_app_acronym varchar(255) NOT NULL,
  last_updated datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (task_notes_id),
  CONSTRAINT task_name FOREIGN KEY (task_name) REFERENCES task (task_name)
  CONSTRAINT task_app_acronym FOREIGN KEY (task_app_acronym) REFERENCES application (app_acronym)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

INSERT INTO application (app_acronym, app_description, app_Rnum, app_startDate, app_endDate, app_permitCreate, app_permitOpen, app_permitToDo, app_permitDoing, app_permitDone, app_createdDate)
VALUES ("apple","","2", NOW(), NOW(), null, null, null, null, null, NOW());

INSERT INTO application (app_acronym, app_description, app_Rnum, app_startDate, app_endDate, app_permitCreate, app_permitOpen, app_permitToDo, app_permitDoing, app_permitDone, app_createdDate)
VALUES ("pear","","2", NOW(), NOW(), "Project Lead", "Project Manager", "Team Member", "Team Member", "Project Lead", NOW());