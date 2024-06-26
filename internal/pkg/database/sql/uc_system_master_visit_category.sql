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

 Date: 29/03/2024 21:05:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_system_master_visit_category
-- ----------------------------
DROP TABLE IF EXISTS `uc_system_master_visit_category`;
CREATE TABLE `uc_system_master_visit_category`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '访问类型标题',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '访问类型' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of uc_system_master_visit_category
-- ----------------------------
INSERT INTO `uc_system_master_visit_category` VALUES (1, '管理员登入', '2023-03-03 17:03:27', '2023-07-20 17:58:32');
INSERT INTO `uc_system_master_visit_category` VALUES (2, '管理员管理', '2023-03-03 17:15:09', '2023-07-20 17:58:48');
INSERT INTO `uc_system_master_visit_category` VALUES (3, '访问日志管理', '2023-03-06 22:05:31', '2023-03-06 22:05:31');
INSERT INTO `uc_system_master_visit_category` VALUES (4, '权限管理', '2023-03-06 22:15:49', '2023-03-06 22:15:49');
INSERT INTO `uc_system_master_visit_category` VALUES (5, '文章管理', '2023-03-22 23:44:18', '2023-03-22 23:44:18');
INSERT INTO `uc_system_master_visit_category` VALUES (6, '员工管理', '2023-03-30 18:22:21', '2023-03-30 18:22:21');
INSERT INTO `uc_system_master_visit_category` VALUES (7, '111', '2023-07-21 11:28:49', '2024-03-26 15:32:11');
INSERT INTO `uc_system_master_visit_category` VALUES (8, '222', '2024-03-26 17:56:51', '2024-03-26 17:58:37');

SET FOREIGN_KEY_CHECKS = 1;
