CREATE TABLE `user` (
    `id` bigint(20) NOT NULL ,
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL ,
    `password` varchar(256) COLLATE utf8mb4_general_ci NOT NULL ,
    `phone` varchar(64) COLLATE utf8mb4_general_ci NOT NULL ,
    `email` varchar(64) COLLATE utf8mb4_general_ci ,
    `phone_valid` tinyint(1) DEFAULT '0' ,
    `email_valid` tinyint(1) DEFAULT '0' ,
    `gender` tinyint(4) NOT NULL DEFAULT '0' ,
    `signup_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
    `last_active` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,
    `status` int(16) NOT NULL DEFAULT '0' ,
    PRIMARY KEY (`id`) ,
    UNIQUE KEY `idx_phone` (`phone`) ,
    KEY `idx_status` (`status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;