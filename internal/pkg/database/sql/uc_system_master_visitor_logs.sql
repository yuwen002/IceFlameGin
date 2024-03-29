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

 Date: 29/03/2024 21:05:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_system_master_visitor_logs
-- ----------------------------
DROP TABLE IF EXISTS `uc_system_master_visitor_logs`;
CREATE TABLE `uc_system_master_visitor_logs`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `account_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '关联用户ID',
  `os_category` tinyint(4) UNSIGNED NULL DEFAULT NULL COMMENT '系统访问类型（1=web端，2=android端，3=IOS端）',
  `visit_category` smallint(5) UNSIGNED NULL DEFAULT NULL COMMENT '访问类型',
  `union_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '关联ID',
  `description` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '访问信息描述',
  `ip_long` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `ip` varchar(47) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '管理员访问记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of uc_system_master_visitor_logs
-- ----------------------------
INSERT INTO `uc_system_master_visitor_logs` VALUES (1, 1, 1, 2, NULL, '编辑角色信息', '2130706433', '127.0.0.1', '2024-03-28 01:08:06', '2024-03-28 01:08:06');
INSERT INTO `uc_system_master_visitor_logs` VALUES (2, 1, 1, 1, 0, '管理员登入', '1', '::1', '2024-03-29 11:22:02', '2024-03-29 11:22:02');
INSERT INTO `uc_system_master_visitor_logs` VALUES (3, 1, 1, 1, 0, '管理员登入', '2130706433', '127.0.0.1', '2024-03-29 15:21:05', '2024-03-29 15:21:05');
INSERT INTO `uc_system_master_visitor_logs` VALUES (4, 1, 1, 1, 0, '管理员登入', '2130706433', '127.0.0.1', '2024-03-29 17:21:55', '2024-03-29 17:21:55');
INSERT INTO `uc_system_master_visitor_logs` VALUES (5, 1, 1, 1, 0, '管理员登入', '2130706433', '127.0.0.1', '2024-03-29 20:31:52', '2024-03-29 20:31:52');
INSERT INTO `uc_system_master_visitor_logs` VALUES (6, 1, 1, 1, 0, '管理员登入', '2130706433', '127.0.0.1', '2024-03-29 20:31:55', '2024-03-29 20:31:55');

SET FOREIGN_KEY_CHECKS = 1;
