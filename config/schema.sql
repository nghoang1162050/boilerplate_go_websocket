use ecommerce_db;

DROP TABLE IF EXISTS `messages`;
DROP TABLE IF EXISTS `room_members`;
DROP TABLE IF EXISTS `rooms`;

CREATE TABLE `rooms` (
    `id` CHAR(36) PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT,
    `host_id` CHAR(36) COLLATE utf8mb4_unicode_ci NOT NULL,
    `is_closed` BIT DEFAULT 0
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT `room_user` FOREIGN KEY (`host_id`) REFERENCES `users`(`id`)
);

CREATE TABLE `room_members` (
    `room_id` CHAR(36),
    `user_id` CHAR(36) COLLATE utf8mb4_unicode_ci,
    `joined_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (room_id, user_id),
    CONSTRAINT `room_members_room` FOREIGN KEY (`room_id`) REFERENCES `rooms`(`id`) ON DELETE CASCADE,
    CONSTRAINT `room_members_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
);

CREATE TABLE `messages` (
    `id` CHAR(36) PRIMARY KEY,
    `room_id` CHAR(36),
    `sender_id` CHAR(36) COLLATE utf8mb4_unicode_ci,
    `content` TEXT NOT NULL,
    `sent_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT `messages_room` FOREIGN KEY (`room_id`) REFERENCES rooms(`id`) ON DELETE CASCADE,
    CONSTRAINT `messages_user` FOREIGN KEY (`sender_id`) REFERENCES users(`id`) ON DELETE SET NULL
);


