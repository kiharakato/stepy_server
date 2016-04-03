CREATE TABLE stepy.todo_groups(
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name TEXT NOT NULL,
    status VARCHAR(128) NOT NULL,
    users_id INTEGER,
    updated_at VARCHAR(255),
    created_at DATETIME
);
