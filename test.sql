/*
Navicat MySQL Data Transfer

Source Server         : 127.0.0.1
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : bill

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2019-04-16 18:02:20
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `test`
-- ----------------------------
DROP TABLE IF EXISTS `test`;
CREATE TABLE `test` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `device_type` int(11) NOT NULL DEFAULT '0' COMMENT '设备类型：1-',
  `device_type_name` varchar(64) NOT NULL,
  `device_type_title` varchar(64) NOT NULL DEFAULT '' COMMENT '设备类型标题',
  `device_screen_width` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '屏宽',
  `device_screen_height` int(11) NOT NULL DEFAULT '0' COMMENT '屏高',
  `update_time` int(11) NOT NULL DEFAULT '0',
  `create_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `i_device_type` (`device_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COMMENT='设备类型表';

-- ----------------------------
-- Records of test
-- ----------------------------
INSERT INTO `test` VALUES ('1', '1', 'aa', 'aa', '1920', '1080', '1541647758', '1541646082');
INSERT INTO `test` VALUES ('2', '2', 'bb', 'bb', '1920', '1080', '1541746764', '1541746764');
INSERT INTO `test` VALUES ('3', '3', 'cc', 'cc', '1920', '1080', '1541753141', '1541749904');
INSERT INTO `test` VALUES ('4', '4', 'dd', 'dd', '1080', '1920', '1541753141', '1541753141');
INSERT INTO `test` VALUES ('5', '5', 'ee', 'ee', '1920', '1080', '1541753141', '1541753141');
INSERT INTO `test` VALUES ('6', '6', 'ff', 'ff', '1920', '360', '1541753141', '1541753141');
INSERT INTO `test` VALUES ('7', '7', 'gg', 'gg', '800', '1280', '1541753141', '1541753141');
INSERT INTO `test` VALUES ('8', '8', 'hh', 'hh', '1080', '1920', '1541753141', '1541753141');
INSERT INTO `test` VALUES ('10', '10', 'ii', 'ii', '1920', '1080', '1541753141', '1541753141');
