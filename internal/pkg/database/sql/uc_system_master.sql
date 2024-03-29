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

 Date: 29/03/2024 21:04:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_system_master
-- ----------------------------
DROP TABLE IF EXISTS `uc_system_master`;
CREATE TABLE `uc_system_master`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `account_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '关联account id',
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户姓名',
  `email` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户邮箱',
  `tel` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户电话',
  `super_master` tinyint(1) NULL DEFAULT 0 COMMENT '超级管理员(=1为超级管理员)',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `account_id`(`account_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '管理员用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of uc_system_master
-- ----------------------------
INSERT INTO `uc_system_master` VALUES (1, 1, '流星雨', 'yuwen002@163.com', '15566036902', 1, '2023-02-22 23:36:37', '2024-03-11 17:52:48');
INSERT INTO `uc_system_master` VALUES (2, 2, '测试修改33', NULL, '15566006661', 0, '2023-02-24 11:31:00', '2023-02-28 13:42:11');
INSERT INTO `uc_system_master` VALUES (3, 3, '流星雨11', NULL, '15566036904', 0, '2023-02-25 23:48:15', '2023-02-25 23:48:15');
INSERT INTO `uc_system_master` VALUES (4, 9, '测试姓名', NULL, '15566006667', 0, '2023-06-16 18:32:40', '2023-06-16 18:32:40');
INSERT INTO `uc_system_master` VALUES (5, 10, '测试姓名2', NULL, '15566036668', 0, '2023-06-16 18:37:19', '2023-06-16 18:37:19');
INSERT INTO `uc_system_master` VALUES (6, 11, '123456', NULL, '15566661001', 0, '2023-06-20 17:55:39', '2023-06-20 17:55:39');
INSERT INTO `uc_system_master` VALUES (7, 12, '111111', NULL, '15566661002', 0, '2023-06-20 17:57:47', '2023-06-20 17:57:47');
INSERT INTO `uc_system_master` VALUES (8, 13, '1111', NULL, '15566661003', 0, '2023-06-20 17:58:59', '2023-06-20 17:58:59');
INSERT INTO `uc_system_master` VALUES (9, 14, '123456', NULL, '15566661004', 0, '2023-06-20 18:10:06', '2023-06-25 17:53:50');
INSERT INTO `uc_system_master` VALUES (10, 15, '123456', NULL, '15566661005', 0, '2023-06-20 18:14:48', '2023-07-18 17:08:41');
INSERT INTO `uc_system_master` VALUES (24, 36, '123456', 'yuwen002@126.com', '15566661006', 0, '2024-02-20 01:18:22', '2024-02-20 01:18:22');
INSERT INTO `uc_system_master` VALUES (25, 37, 'AAACC', 'yuwen0012@126.com', '15566660001', 0, '2024-03-23 00:46:18', '2024-03-25 11:30:00');

SET FOREIGN_KEY_CHECKS = 1;
