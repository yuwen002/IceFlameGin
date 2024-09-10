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

 Date: 09/04/2024 23:45:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article_tag
-- ----------------------------
DROP TABLE IF EXISTS `article_tag`;
CREATE TABLE `article_tag`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'tag内容',
  `remark` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `sort` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '排序',
  `status` tinyint(4) UNSIGNED NULL DEFAULT 0 COMMENT '显示状态（0=显示，1=隐藏）',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文章tags' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of article_tag
-- ----------------------------
INSERT INTO `article_tag` VALUES (1, '热门', NULL, NULL, 0, '2023-03-27 16:00:11', '2023-03-27 16:00:11');
INSERT INTO `article_tag` VALUES (2, '热搜', '测试修改', 1, 0, '2023-03-27 16:00:39', '2023-03-27 16:55:26');
INSERT INTO `article_tag` VALUES (4, '热门11', NULL, NULL, 0, '2023-03-28 10:54:33', '2023-03-28 10:54:33');
INSERT INTO `article_tag` VALUES (5, '情感', NULL, NULL, 0, '2023-03-28 10:55:11', '2023-03-28 10:55:11');
INSERT INTO `article_tag` VALUES (6, '写实', NULL, NULL, 0, '2023-03-28 10:55:17', '2023-03-28 10:55:17');

SET FOREIGN_KEY_CHECKS = 1;
