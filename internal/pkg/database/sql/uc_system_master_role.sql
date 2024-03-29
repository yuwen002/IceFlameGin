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

 Date: 29/03/2024 21:04:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_system_master_role
-- ----------------------------
DROP TABLE IF EXISTS `uc_system_master_role`;
CREATE TABLE `uc_system_master_role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '角色名称',
  `remark` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `supper_master` tinyint(1) NULL DEFAULT 0 COMMENT '1为超级管理员模块',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '管理员角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of uc_system_master_role
-- ----------------------------
INSERT INTO `uc_system_master_role` VALUES (1, '超级管理员', '超级管理员', NULL, '2023-03-07 16:53:22', '2023-04-02 01:04:34');
INSERT INTO `uc_system_master_role` VALUES (2, '普通管理员', '测试修改', NULL, '2023-03-07 16:53:49', '2023-07-03 18:00:12');
INSERT INTO `uc_system_master_role` VALUES (3, '测试管理员', '测试管理员', 0, '2023-07-04 15:56:23', '2023-07-04 15:56:23');
INSERT INTO `uc_system_master_role` VALUES (4, '测试管理员1', '测试管理员1', 0, '2023-07-04 15:56:51', '2024-03-18 01:08:35');
INSERT INTO `uc_system_master_role` VALUES (5, '工111', '寺', 0, '2023-07-04 16:06:48', '2024-03-18 23:20:38');
INSERT INTO `uc_system_master_role` VALUES (6, '圭', '寺', 0, '2023-07-04 16:31:40', '2023-07-04 16:31:40');
INSERT INTO `uc_system_master_role` VALUES (7, '222', '113', 0, '2023-07-04 16:34:21', '2024-03-26 15:18:21');
INSERT INTO `uc_system_master_role` VALUES (8, '2221', '4441', 0, '2023-07-04 16:34:47', '2024-03-19 00:06:24');
INSERT INTO `uc_system_master_role` VALUES (9, '124', '111ff', 0, '2023-07-04 16:38:25', '2024-03-22 22:37:33');
INSERT INTO `uc_system_master_role` VALUES (10, 'AAACC', '', 0, '2024-03-23 00:39:08', '2024-03-23 00:39:08');
INSERT INTO `uc_system_master_role` VALUES (11, 'AAACC', '', 0, '2024-03-23 00:41:35', '2024-03-28 01:08:06');

SET FOREIGN_KEY_CHECKS = 1;
