CREATE USER 'sqlsheet'@'localhost' IDENTIFIED WITH caching_sha2_password BY 'sqlsheet';
CREATE DATABASE sqlsheet;
GRANT ALL PRIVILEGES ON sqlsheet.* TO 'sqlsheet'@'localhost';
