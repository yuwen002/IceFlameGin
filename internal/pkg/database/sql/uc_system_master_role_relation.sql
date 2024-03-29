/*
 Navicat MySQL Data Transfer

 Source Server         : 测试
 Source Server Type    : MySQL
 Source Server Version : 50738
 Source Host           : 82.157.248.230:3306
 Source Schema         : go_test

 Target Server Type    : MySQL
 Target Server Version : 50738
 File Encoding         : 65001

 Date: 29/03/2024 21:05:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_system_master_role_relation
-- ----------------------------
DROP TABLE IF EXISTS `uc_system_master_role_relation`;
CREATE TABLE `uc_system_master_role_relation`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `account_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '管理员ID',
  `role_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '身份ID',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `account_id`(`account_id`) USING BTREE,
  INDEX `role_id`(`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色信息关联管理ID' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of uc_system_master_role_relation
-- ----------------------------
INSERT INTO `uc_system_master_role_relation` VALUES (1, 1, 1, '2023-03-09 13:27:27', '2024-03-22 22:40:40');
INSERT INTO `uc_system_master_role_relation` VALUES (2, 2, 2, '2023-03-09 13:51:41', '2023-03-09 14:36:36');
INSERT INTO `uc_system_master_role_relation` VALUES (3, 3, 1, '2023-03-09 13:56:00', '2023-03-09 13:56:00');
INSERT INTO `uc_system_master_role_relation` VALUES (4, 10, 2, '2023-07-06 17:28:39', '2023-07-06 17:28:39');
INSERT INTO `uc_system_master_role_relation` VALUES (5, 14, 3, '2024-03-22 11:35:51', '2024-03-22 23:16:00');

SET FOREIGN_KEY_CHECKS = 1;
