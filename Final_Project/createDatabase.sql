CREATE DATABASE IF NOT EXISTS `my_gram_db`;

USE `my_gram_db`;

CREATE TABLE IF NOT EXISTS `tb_users` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `username` VARCHAR(255) UNIQUE NOT NULL,
    `email` VARCHAR(255) UNIQUE NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `age` INT NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `tb_photos` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `title` VARCHAR(255) NOT NULL,
    `caption` TEXT,
    `photo_url` VARCHAR(255) NOT NULL,
    `user_id` INT NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (`user_id`) REFERENCES `tb_users`(`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    INDEX (`user_id`)  -- Tambahkan indeks pada user_id
);



CREATE TABLE IF NOT EXISTS `tb_photo` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `created_at` DATETIME(3) NULL,
    `updated_at` DATETIME(3) NULL,
    `title` LONGTEXT NOT NULL,
    `caption` LONGTEXT,
    `photo_url` LONGTEXT NOT NULL,
    `user_id` BIGINT UNSIGNED,
    CONSTRAINT `fk_tb_users_photos` FOREIGN KEY (`user_id`) REFERENCES `tb_users`(`id`) ON DELETE SET NULL ON UPDATE CASCADE
);


CREATE TABLE IF NOT EXISTS `tb_comments` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `message` TEXT NOT NULL,
    `photo_id` BIGINT UNSIGNED,
    `user_id` BIGINT UNSIGNED,
    `created_at` DATETIME(3) NULL,
    `updated_at` DATETIME(3) NULL,
    CONSTRAINT `fk_tb_photos_comments` FOREIGN KEY (`photo_id`) REFERENCES `tb_photo`(`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT `fk_tb_users_comments` FOREIGN KEY (`user_id`) REFERENCES `tb_users`(`id`) ON DELETE SET NULL ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS `tb_social_media` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `social_media_url` VARCHAR(255) NOT NULL,
    `user_id` BIGINT UNSIGNED,
    `created_at` DATETIME(3) NULL,
    `updated_at` DATETIME(3) NULL,
    CONSTRAINT `fk_tb_users_social_media` FOREIGN KEY (`user_id`) REFERENCES `tb_users`(`id`) ON DELETE SET NULL ON UPDATE CASCADE
);

