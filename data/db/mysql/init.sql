CREATE DATABASE IF NOT EXISTS Device_Scalable;

USE Device_Scalable;

CREATE TABLE devices(
    `id` bigint(20) PRIMARY KEY AUTO_INCREMENT,
    `atm_id` varchar(255) not null unique,
    `address` varchar(255),
    `agency` varchar(255),
    `created_at` timestamp default CURRENT_TIMESTAMP,
    `updated_at` timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE device_phone_access(
    `id` bigint(20) PRIMARY KEY AUTO_INCREMENT,
    `device_id` bigint(20) NOT NULL,
    `value` varchar(255) NOT NULL,
    FOREIGN KEY (`device_id`) REFERENCES devices(`id`)
)
