/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 80013
Source Host           : localhost:3306
Source Database       : problem_123

Target Server Type    : MYSQL
Target Server Version : 80013
File Encoding         : 65001

Date: 2020-05-07 16:35:29
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `customers`
-- ----------------------------
DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers` (
  `Id` int(11) DEFAULT NULL,
  `Name` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of customers
-- ----------------------------
INSERT INTO `customers` VALUES ('1', 'Joe');
INSERT INTO `customers` VALUES ('2', 'Henry');
INSERT INTO `customers` VALUES ('3', 'Sam');
INSERT INTO `customers` VALUES ('4', 'Max');

-- ----------------------------
-- Table structure for `orders`
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `Id` int(11) DEFAULT NULL,
  `CustomerId` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of orders
-- ----------------------------
INSERT INTO `orders` VALUES ('1', '3');
INSERT INTO `orders` VALUES ('2', '1');
