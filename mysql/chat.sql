/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 50724
 Source Host           : localhost:3306
 Source Schema         : chat

 Target Server Type    : MySQL
 Target Server Version : 50724
 File Encoding         : 65001

 Date: 09/11/2019 00:58:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for community
-- ----------------------------
DROP TABLE IF EXISTS `community`;
CREATE TABLE `community` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(30) DEFAULT NULL COMMENT '名称',
  `ownerid` bigint(20) DEFAULT NULL COMMENT '群主ID',
  `icon` varchar(250) DEFAULT NULL COMMENT '群logo',
  `cate` int(11) DEFAULT NULL COMMENT '类别',
  `memo` varchar(120) DEFAULT NULL COMMENT '描述',
  `createat` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='群';

-- ----------------------------
-- Table structure for contact
-- ----------------------------
DROP TABLE IF EXISTS `contact`;
CREATE TABLE `contact` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `ownerid` bigint(20) DEFAULT NULL COMMENT '自己id',
  `dstobj` bigint(20) DEFAULT NULL COMMENT '对方id',
  `cate` int(11) DEFAULT NULL COMMENT '类型',
  `memo` varchar(120) DEFAULT NULL COMMENT '备注',
  `createat` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of contact
-- ----------------------------
BEGIN;
INSERT INTO `contact` VALUES (17, 5, 4, 1, '', '2019-11-09 00:35:47');
INSERT INTO `contact` VALUES (18, 4, 5, 1, '', '2019-11-09 00:35:47');
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(64) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `mobile` varchar(20) DEFAULT NULL COMMENT '手机号码 账号',
  `passwd` varchar(40) DEFAULT NULL COMMENT '密码',
  `avatar` varchar(150) DEFAULT NULL COMMENT '图像',
  `sex` varchar(2) DEFAULT NULL COMMENT '性别',
  `nickname` varchar(20) DEFAULT NULL COMMENT '网名',
  `salt` varchar(10) DEFAULT NULL COMMENT '加盐随机字符串6',
  `online` int(10) DEFAULT NULL COMMENT '是否在线',
  `token` varchar(40) DEFAULT NULL COMMENT 'token',
  `memo` varchar(140) DEFAULT NULL COMMENT '备注',
  `createat` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (4, '18822855252', '69a5f32d315c5761e1a4245a21e0d912', '/mnt/15732297721298498081.jpg', 'M', '游浩', '007887', 0, 'EAFD33DF5180B4BF8E90A554B293FAEF', '', '2019-11-09 00:17:48');
INSERT INTO `user` VALUES (5, '18822855251', '231a8746dcb083fb5a6ac0f80bdbf9a8', '/static/images/user.jpg', 'W', '李艳玲', '002081', 0, '37AE27F74F1D67322AAB8BB25EA0DAC2', '', '2019-11-09 00:22:06');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
