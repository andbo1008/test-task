CREATE TABLE films(
    id serial PRIMARY KEY,
    title varchar(255) NOT NULL,
    releaseDate DATETIME DEFAULT NULL,
    createdAt DATETIME NOT NULL
);