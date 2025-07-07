-- start set up init script
GRANT ALL PRIVILEGES ON *.*  TO 'template_us'@'%';
SET GLOBAL log_bin_trust_function_creators = 1;
FLUSH PRIVILEGES;
-- end set up init script


CREATE DATABASE  IF NOT EXISTS foo;
USE foo;
CREATE TABLE IF NOT EXISTS users (
                                     id int NOT NULL AUTO_INCREMENT,
                                     name VARCHAR(30) NOT NULL,
    lastname VARCHAR(30) NOT NULL,

    PRIMARY KEY (id)
    );
