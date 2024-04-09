/*
 Navicat MySQL Data Transfer

 Source Server         : kd_test
 Source Server Type    : MySQL
 Source Server Version : 50738
 Source Host           : 82.157.248.230:3306
 Source Schema         : go_test

 Target Server Type    : MySQL
 Target Server Version : 50738
 File Encoding         : 65001

 Date: 09/04/2024 15:30:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for single_page
-- ----------------------------
DROP TABLE IF EXISTS `single_page`;
CREATE TABLE `single_page`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '页面标题',
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '页面描述',
  `keyword` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '页面关键字',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '页面内容',
  `thumbnail` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '缩略图',
  `click` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '点击量',
  `status` tinyint(3) UNSIGNED NULL DEFAULT 1 COMMENT '显示状态（1=显示，2=隐藏）',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of single_page
-- ----------------------------
INSERT INTO `single_page` VALUES (1, '测试单页', NULL, NULL, NULL, NULL, 0, 1, '2023-03-22 16:18:01', '2023-09-27 16:51:11');
INSERT INTO `single_page` VALUES (2, '测试频道', NULL, NULL, NULL, NULL, 0, 1, '2023-03-24 13:39:57', '2023-03-24 13:39:57');
INSERT INTO `single_page` VALUES (3, '艾师傅', '123', '123', '<p>阿斯蒂芬</p>', NULL, 0, 1, '2023-08-28 16:56:01', '2023-08-28 16:56:01');
INSERT INTO `single_page` VALUES (4, 'single-page11', 'single-page', 'single-page', '<p>single-page<br></p>', '\\uploads\\2024\\04\\20240405_000031_942.png', 0, 1, '2024-04-01 21:23:44', '2024-04-06 23:03:09');

SET FOREIGN_KEY_CHECKS = 1;
