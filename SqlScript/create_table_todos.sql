CREATE TABLE stepy.todos(
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    action TEXT NOT NULL,
    status VARCHAR(128) NOT NULL,
    todo_gropus_id INTEGER,
    users_id INTEGER,
    updated_at VARCHAR(255),
    created_at DATETIME
);
