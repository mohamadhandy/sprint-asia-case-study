CREATE TABLE task (
  id          INT PRIMARY KEY AUTO_INCREMENT,
  title       VARCHAR(255),
  description TEXT,
  created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
  completed   BOOLEAN DEFAULT FALSE
);

CREATE TABLE deadline (
  id          INT PRIMARY KEY AUTO_INCREMENT,
  task_id     INT,
  due_at      DATETIME,
  FOREIGN KEY (task_id) REFERENCES task(id)
);

CREATE TABLE sub_task (
  id          INT PRIMARY KEY AUTO_INCREMENT,
  task_id     INT,
  title       VARCHAR(255),
  completed   BOOLEAN DEFAULT FALSE,
  FOREIGN KEY (task_id) REFERENCES task(id)
);

CREATE TABLE task_progress (
  task_id     INT,
  percentage  INT,
  PRIMARY KEY (task_id),
  FOREIGN KEY (task_id) REFERENCES task(id)
);

CREATE TABLE task_status (
  id          INT PRIMARY KEY,
  status      ENUM('ongoing', 'completed')
);

