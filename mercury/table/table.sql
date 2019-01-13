/*
Navicat MySQL Data Transfer

Source Server         : golang
Source Server Version : 80011
Source Host           : localhost:3306
Source Database       : mercury

Target Server Type    : MYSQL
Target Server Version : 80011
File Encoding         : 65001

Date: 2019-01-12 11:07:36
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for answer
-- ----------------------------
DROP TABLE IF EXISTS `answer`;
CREATE TABLE `answer` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `answer_id` bigint(20) unsigned NOT NULL,
  `content` text COLLATE utf8mb4_general_ci NOT NULL,
  `comment_count` int(10) unsigned NOT NULL,
  `voteup_count` int(11) NOT NULL,
  `author_id` bigint(20) NOT NULL,
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1',
  `can_comment` tinyint(3) unsigned NOT NULL DEFAULT '1',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_answer_id` (`answer_id`),
  KEY `idx_author_Id` (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of answer
-- ----------------------------

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `category_id` int(10) unsigned NOT NULL,
  `category_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_category_id` (`category_id`),
  UNIQUE KEY `idx_category_name` (`category_name`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES ('1', '1', '技术', '2019-01-01 08:30:40', '2019-01-01 08:30:40');
INSERT INTO `category` VALUES ('2', '2', '情感', '2019-01-01 08:31:07', '2019-01-01 08:31:07');
INSERT INTO `category` VALUES ('3', '3', '王者荣耀', '2019-01-01 08:31:25', '2019-01-01 08:31:25');
INSERT INTO `category` VALUES ('4', '4', '吃鸡', '2019-01-01 15:45:13', '2019-01-01 15:45:13');
INSERT INTO `category` VALUES ('5', '5', '科幻', '2019-01-05 23:02:43', '2019-01-05 23:02:43');

-- ----------------------------
-- Table structure for question
-- ----------------------------
DROP TABLE IF EXISTS `question`;
CREATE TABLE `question` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `question_id` bigint(20) NOT NULL COMMENT '问题id',
  `caption` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '问题标题',
  `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '问题内容',
  `author_id` bigint(20) NOT NULL COMMENT '作者的用户id',
  `category_id` bigint(20) NOT NULL COMMENT '所属栏目',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '问题状态',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_author_id` (`author_id`),
  KEY `idx_question_id` (`question_id`),
  KEY `idx_category_id` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of question
-- ----------------------------
INSERT INTO `question` VALUES ('1', '229544174527971329', '上岛咖啡斯拉夫斯卡就开始了是开始实施上课试试看是', '开始了打发时间浪费时间浪费多少反倒是浪费时间开放时间发生的雷锋精神离开房间是否是否', '224533709452214273', '1', '1', '2019-01-01 21:16:29', '2019-01-01 21:16:29');
INSERT INTO `question` VALUES ('2', '230134441849126913', '未来三十年内，哪些行业的工作人员可能会被人工智能取代？', '人工智能这些年的快速发展，在某些领域已经开始渐渐取代人类的工作岗位了。未来这种情况是否会越来越严重？以后的人类会进入空虚的享乐时代么？', '224533709452214273', '1', '1', '2019-01-05 23:00:16', '2019-01-05 23:00:16');
INSERT INTO `question` VALUES ('3', '230134511709454337', '你见过最渣的渣女有多渣？', '我一个玩的挺好的舍友，在一家旁边很多酒吧的电玩电玩城上班，酒吧多，帅哥，男生也就多了，喜欢她的男生就也还挺多的，她说她自己是个渣女，因为同时跟四五个男生暧昧，出去玩都是跟这个男生玩完，又去跟另一个男生玩。她长得一般，但身材可以，你们觉得她渣吗？\n\n\n', '224533709452214273', '1', '1', '2019-01-05 23:00:58', '2019-01-05 23:00:58');
INSERT INTO `question` VALUES ('4', '230134710704013313', '你觉得《三体》中最残忍的一句话是什么？', '你觉得《三体》中最残忍的一句话是什么？', '224533709452214273', '5', '1', '2019-01-05 23:02:56', '2019-01-05 23:02:56');

-- ----------------------------
-- Table structure for question_answer_rel
-- ----------------------------
DROP TABLE IF EXISTS `question_answer_rel`;
CREATE TABLE `question_answer_rel` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `question_id` bigint(20) NOT NULL,
  `answer_id` bigint(20) NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_question_answer` (`question_id`,`answer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of question_answer_rel
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `sex` tinyint(4) NOT NULL DEFAULT '0',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`) USING BTREE,
  UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('3', '223181233645944833', 'admin2', '', '21968b082e3af16563cad001184ddfa8', 'admin', '0', '2018-11-18 23:46:21', '2018-11-18 23:46:21');
INSERT INTO `user` VALUES ('4', '224533709452214273', 'admin', '', '21968b082e3af16563cad001184ddfa8', 'admin', '0', '2018-11-28 07:41:59', '2018-11-28 07:41:59');
INSERT INTO `user` VALUES ('5', '224533754784251905', 'admin22', '', 'c8814aea2aa8a4feae0569b18a97cd00', 'admin22', '0', '2018-11-28 07:42:26', '2018-11-28 07:42:26');
INSERT INTO `user` VALUES ('6', '224533951312560129', 'admin222', '', 'c8814aea2aa8a4feae0569b18a97cd00', 'admin222', '0', '2018-11-28 07:44:24', '2018-11-28 07:44:24');
INSERT INTO `user` VALUES ('7', '224534423356309505', 'admin2221', '', 'c8814aea2aa8a4feae0569b18a97cd00', 'admin2221', '0', '2018-11-28 07:49:05', '2018-11-28 07:49:05');
INSERT INTO `user` VALUES ('8', '224534702126530561', 'admin22212', '', 'c8814aea2aa8a4feae0569b18a97cd00', 'admin22212', '0', '2018-11-28 07:51:51', '2018-11-28 07:51:51');
INSERT INTO `user` VALUES ('9', '224534879344263169', 'admin222121', '', 'c8814aea2aa8a4feae0569b18a97cd00', 'admin222121', '0', '2018-11-28 07:53:37', '2018-11-28 07:53:37');
INSERT INTO `user` VALUES ('10', '224535034466402305', 'admin2221212', '', 'c8814aea2aa8a4feae0569b18a97cd00', 'admin2221212', '0', '2018-11-28 07:55:09', '2018-11-28 07:55:09');
INSERT INTO `user` VALUES ('11', '224535766808657921', 'admin22212121', '', 'c8814aea2aa8a4feae0569b18a97cd00', 'admin22212121', '0', '2018-11-28 08:02:26', '2018-11-28 08:02:26');
INSERT INTO `user` VALUES ('12', '224536026788397057', 'admin111', '', '21968b082e3af16563cad001184ddfa8', 'admin111', '0', '2018-11-28 08:05:01', '2018-11-28 08:05:01');
INSERT INTO `user` VALUES ('13', '224536207143469057', 'admin1111', '', '21968b082e3af16563cad001184ddfa8', 'admin1111', '0', '2018-11-28 08:06:48', '2018-11-28 08:06:48');
INSERT INTO `user` VALUES ('14', '224536320255459329', 'admin11112', '', '21968b082e3af16563cad001184ddfa8', 'admin11112', '0', '2018-11-28 08:07:56', '2018-11-28 08:07:56');
INSERT INTO `user` VALUES ('15', '224536455865696257', 'admin111122', '', '21968b082e3af16563cad001184ddfa8', 'admin111122', '0', '2018-11-28 08:09:16', '2018-11-28 08:09:16');
INSERT INTO `user` VALUES ('16', '224536664054169601', 'admin1111223', '', '21968b082e3af16563cad001184ddfa8', 'admin1111223', '0', '2018-11-28 08:11:21', '2018-11-28 08:11:21');
INSERT INTO `user` VALUES ('17', '224536736212975617', 'admin11112234', '', '21968b082e3af16563cad001184ddfa8', 'admin11112234', '0', '2018-11-28 08:12:04', '2018-11-28 08:12:04');
INSERT INTO `user` VALUES ('18', '225180933609750529', 'admi', '', '46f19871f7a3088f42ecf44102d2e6e9', 'admin', '0', '2018-12-02 18:51:35', '2018-12-02 18:51:35');
INSERT INTO `user` VALUES ('19', '225181104049487873', 'admi2', '', '46f19871f7a3088f42ecf44102d2e6e9', 'admin2', '0', '2018-12-02 18:53:17', '2018-12-02 18:53:17');
INSERT INTO `user` VALUES ('20', '225181127923466241', 'admi22', '', '46f19871f7a3088f42ecf44102d2e6e9', 'admin2', '0', '2018-12-02 18:53:31', '2018-12-02 18:53:31');
INSERT INTO `user` VALUES ('21', '225182107310227457', 'admin2222', '', '21968b082e3af16563cad001184ddfa8', 'admin', '0', '2018-12-02 19:03:15', '2018-12-02 19:03:15');
INSERT INTO `user` VALUES ('22', '225183383855038465', 'admin1111222', '', '6a74dfafca685f959141472ffa1d907b', 'admin', '0', '2018-12-02 19:15:56', '2018-12-02 19:15:56');
INSERT INTO `user` VALUES ('23', '225183813502763009', 'nadmi', '', '8b6a48f398cf7b5b9ae62adf75791e88', 'admi', '0', '2018-12-02 19:20:12', '2018-12-02 19:20:12');
INSERT INTO `user` VALUES ('24', '225184375774380033', 'admin434444', '', '21968b082e3af16563cad001184ddfa8', 'admin', '0', '2018-12-02 19:25:47', '2018-12-02 19:25:47');
INSERT INTO `user` VALUES ('25', '225186212074225665', 'DKDK', '', '9c0ac25434be0cf1ceeafaa11c80a578', 'SMIN', '0', '2018-12-02 19:44:01', '2018-12-02 19:44:01');
INSERT INTO `user` VALUES ('26', '225190952694710273', 'admin222222222', '', '21968b082e3af16563cad001184ddfa8', 'admin', '0', '2018-12-02 20:31:07', '2018-12-02 20:31:07');
INSERT INTO `user` VALUES ('27', '225191497299918849', 'sherlockhua', '拈花湾', 'e46445d7555d22b48bea1997cd607420', 'sherlockhua@163.com', '1', '2018-12-02 20:36:32', '2018-12-02 20:36:32');
INSERT INTO `user` VALUES ('28', '225191580145811457', 'kala', 'kalo', '76bac7d7ee3687f8c583769d410a2c0f', 'kala@qq.com', '2', '2018-12-02 20:37:21', '2018-12-02 20:37:21');
