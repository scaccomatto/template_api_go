CREATE DATABASE  IF NOT EXISTS foo;
USE foo;
CREATE TABLE IF NOT EXISTS users (
         id int NOT NULL AUTO_INCREMENT,
         name VARCHAR(30) NOT NULL,
         lastname VARCHAR(30) NOT NULL,

    PRIMARY KEY (id)
    );
