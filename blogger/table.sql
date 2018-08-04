/*
 Navicat Premium Data Transfer

 Source Server         : golang
 Source Server Type    : MySQL
 Source Server Version : 80011
 Source Host           : localhost:3306
 Source Schema         : blogger

 Target Server Type    : MySQL
 Target Server Version : 80011
 File Encoding         : 65001

 Date: 04/08/2018 22:20:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章内容',
  `title` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章标题',
  `view_count` int(255) UNSIGNED NOT NULL COMMENT '阅读次数',
  `create_time` timestamp(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '发布时间',
  `comment_count` int(255) UNSIGNED NOT NULL COMMENT '评论次数',
  `username` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '作者',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '评论id',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论内容',
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论作者',
  `create_time` timestamp(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '评论发布时间',
  `status` int(255) UNSIGNED NOT NULL COMMENT '评论状态: 0, 删除；1， 正常',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
