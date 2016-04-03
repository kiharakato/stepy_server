CREATE TABLE stepy.users(
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    uuid VARCHAR(128) NOT NULL,
    name TEXT,
    email VARCHAR(255),
    created_at DATETIME
);
