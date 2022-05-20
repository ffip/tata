CREATE TABLE `zg_queue` (
  `id` int NOT NULL AUTO_INCREMENT,
  `type` tinyint NOT NULL DEFAULT '0',
  `data` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `path` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `createAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `pushAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_type` (`type`) USING BTREE,
  KEY `ix_pushAt` (`pushAt`) USING BTREE,
  KEY `ix_createAt` (`createAt`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;