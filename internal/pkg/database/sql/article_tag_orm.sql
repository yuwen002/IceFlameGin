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

 Date: 09/04/2024 23:45:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article_tag_orm
-- ----------------------------
DROP TABLE IF EXISTS `article_tag_orm`;
CREATE TABLE `article_tag_orm`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) NULL DEFAULT NULL COMMENT 'tag ID',
  `article_id` int(10) NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文章tags' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of article_tag_orm
-- ----------------------------
INSERT INTO `article_tag_orm` VALUES (1, 1, 14, '2023-03-28 14:13:59', '2023-03-28 14:13:59');
INSERT INTO `article_tag_orm` VALUES (2, 2, 14, '2023-03-28 14:13:59', '2023-03-28 14:13:59');

SET FOREIGN_KEY_CHECKS = 1;
