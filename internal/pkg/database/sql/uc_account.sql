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

 Date: 29/03/2024 21:02:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_account
-- ----------------------------
DROP TABLE IF EXISTS `uc_account`;
CREATE TABLE `uc_account`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `identity_card_id` char(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '关联用户身份证信息ID',
  `username` char(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `password_hash` char(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户密码',
  `tel` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户手机号码',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '用户状态(0.启用1.停用)',
  `real_name_type` tinyint(1) NULL DEFAULT 1 COMMENT '实名状态1=未实名2=已上传未审核3=审核驳回4=实名成功',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `identity_card_id`(`identity_card_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 38 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户中心' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of uc_account
-- ----------------------------
INSERT INTO `uc_account` VALUES (1, '0', 'SA_15566036902', '$2a$10$PGcqoS3fqh4UkiimESrRxeNNduAj/Ljur9Z4WghFW9Ut7/UQjRz3i', 'SA_15566036902', 0, 1, '2023-02-22 23:36:37', '2024-03-04 00:05:47');
INSERT INTO `uc_account` VALUES (2, '0', 'SA_yuwen002', '$2a$10$ahUEJ36yYJfY/.CL/nyfk.q/tYd7PgccXrzY4DydeMTMOmh4DPCu.', 'SA_15566006666', 0, 1, '2023-02-24 11:31:00', '2023-03-20 17:40:28');
INSERT INTO `uc_account` VALUES (3, '0', 'SA_yuwen003', '$2a$10$UwsnYTpFIxPHPWm7ufGQr..kw0H7Q7BN3cPFb4wml./AF48fbJJ1K', 'SA_15566036904', 0, 1, '2023-02-25 23:48:15', '2023-02-25 23:48:15');
INSERT INTO `uc_account` VALUES (8, '0', 'EMP_15566661000', '$2a$10$cJS36JbDEXVa/D2iWvgA3OwnFqdAmVjcHKiUa6qeJEn5C2/fYEWry', 'EMP_15566661000', 0, 1, '2023-04-03 16:56:15', '2023-04-05 01:07:02');
INSERT INTO `uc_account` VALUES (9, '0', 'SA_yuwen004', '$2a$10$YxTM8W11IWORiMaNb8Gs/eoasvaJajTpU/NhEMeyqR06QkLOaVVeS', 'SA_15566006667', 0, 1, '2023-06-16 18:32:40', '2023-06-16 18:32:40');
INSERT INTO `uc_account` VALUES (10, '0', 'SA_yuwen005', '$2a$10$0W626TL2nQPXeBVgPl8ZIOEC8x2wQvAqJwVsra/uRHg7wdYGOXjtK', 'SA_15566036668', 0, 1, '2023-06-16 18:37:19', '2023-06-16 18:37:19');
INSERT INTO `uc_account` VALUES (11, '0', 'SA_15566661001', '$2a$10$b1TtiLVei4fKQGCe539K9OvZD7jzgfL9rUA9Ft4lzA46JgDbdzXD6', 'SA_15566661001', 0, 1, '2023-06-20 17:55:39', '2023-06-20 17:55:39');
INSERT INTO `uc_account` VALUES (12, '0', 'SA_15566661002', '$2a$10$uFhTq2ErtRchRJrPwa9NOemIso84LKO72qFmCjtDwoD38mSIgIIHq', 'SA_15566661002', 0, 1, '2023-06-20 17:57:47', '2023-06-20 17:57:47');
INSERT INTO `uc_account` VALUES (13, '0', 'SA_15566661003', '$2a$10$.lQ0vW210qCDoyhG1DelHeaNRhpqVR37BFiapiBroicSSJSambGeO', 'SA_15566661003', 0, 1, '2023-06-20 17:58:59', '2023-06-28 15:54:58');
INSERT INTO `uc_account` VALUES (14, '0', 'SA_15566661004', '$2a$10$q.sIavCa2IfkD5wd0z7VwOwVAbkgPMq7xT1SMsZTKG3NdYns0vEz.', 'SA_15566661004', 0, 1, '2023-06-20 18:10:06', '2023-06-28 15:55:00');
INSERT INTO `uc_account` VALUES (15, '0', 'SA_15566661005', '$2a$10$nrc.UIAXshzUji477ySU8ufqYLl9QeMJvk1g4tHRSFNjwd6LJ61Hq', 'SA_15566661005', 1, 1, '2023-06-20 18:14:48', '2023-07-18 18:05:34');
INSERT INTO `uc_account` VALUES (36, '0', 'SA_15566661006', '$2a$10$ELoJAnqbA.C8El58FloWL.hYXBMm2H.PfQe6KthoiW008d/54Cyxq', 'SA_15566661006', 0, 1, '2024-02-20 01:18:21', '2024-02-20 01:18:21');
INSERT INTO `uc_account` VALUES (37, '0', 'SA_15566660001', '$2a$10$h0Ht16Q8EMGltIJSquMNwebOUXI9mNKKgoz2mixz1w3OaMf.Sv8YS', 'SA_15566660001', 0, 1, '2024-03-23 00:46:18', '2024-03-23 00:46:18');

SET FOREIGN_KEY_CHECKS = 1;
