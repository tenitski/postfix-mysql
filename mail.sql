CREATE DATABASE mail;
USE DATABASE mail;

GRANT ALL ON mail.* TO 'mail'@'localhost' IDENTIFIED BY 'mail';
FLUSH PRIVILEGES;

CREATE USER 'mail'@'localhost' IDENTIFIED BY 'mail';
GRANT ALL ON mail.* TO 'mail'@'localhost';

CREATE TABLE users (
    user varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    PRIMARY KEY (user)
);

INSERT INTO users (user, password) VALUES
    ("mailer@example.com", "password1"),
    ("mailer@example.net", "password2"),
    ("mailer@example.org", "password3")
;

CREATE TABLE senders (
    sender varchar(255) NOT NULL,
    user varchar(255) NOT NULL,
    PRIMARY KEY (sender,user)
);

INSERT INTO senders (sender, user) VALUES
    ("@example.com","mailer@example.com"),
    ("@example.net","mailer@example.net"),
    ("@mail.example.net","mailer@example.net"),
    ("tenitski@gmail.com","mailer@example.org")
;


-- Table structure v2

drop table user;
drop table sender;

CREATE TABLE user (
    login varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    PRIMARY KEY (login)
);

CREATE TABLE sender (
    login varchar(255) NOT NULL,
    sender varchar(255) NOT NULL,
    PRIMARY KEY (login, sender),
    FOREIGN KEY (login) REFERENCES user(login)
);

INSERT INTO user (login, password) VALUES
    ("mailer@example.com", "password1"),
    ("mailer@example.net", "password2"),
    ("mailer@example.org", "password3")
;


INSERT INTO sender (sender, login) VALUES
    ("@example.com","mailer@example.com"),
    ("@example.net","mailer@example.net"),
    ("@mail.example.net","mailer@example.net"),
    ("tenitski@gmail.com","mailer@example.org")
;

