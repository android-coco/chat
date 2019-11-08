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

 Date: 07/11/2019 11:09:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for community
-- ----------------------------
DROP TABLE IF EXISTS `community`;
CREATE TABLE `community` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) DEFAULT NULL,
  `ownerid` bigint(20) DEFAULT NULL,
  `icon` varchar(250) DEFAULT NULL,
  `cate` int(11) DEFAULT NULL,
  `memo` varchar(120) DEFAULT NULL,
  `createat` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for contact
-- ----------------------------
DROP TABLE IF EXISTS `contact`;
CREATE TABLE `contact` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ownerid` bigint(20) DEFAULT NULL,
  `dstobj` bigint(20) DEFAULT NULL,
  `cate` int(11) DEFAULT NULL,
  `memo` varchar(120) DEFAULT NULL,
  `createat` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(64) NOT NULL AUTO_INCREMENT,
  `mobile` varchar(20) DEFAULT NULL,
  `passwd` varchar(40) DEFAULT NULL,
  `avatar` varchar(150) DEFAULT NULL,
  `sex` varchar(2) DEFAULT NULL,
  `nickname` varchar(20) DEFAULT NULL,
  `salt` varchar(10) DEFAULT NULL,
  `online` int(10) DEFAULT NULL,
  `token` varchar(40) DEFAULT NULL,
  `memo` varchar(140) DEFAULT NULL,
  `createat` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
