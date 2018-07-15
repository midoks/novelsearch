/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50634
 Source Host           : localhost:3306
 Source Schema         : novelsearch

 Target Server Type    : MySQL
 Target Server Version : 50634
 File Encoding         : 65001

 Date: 15/07/2018 22:48:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app_debug
-- ----------------------------
DROP TABLE IF EXISTS `app_debug`;
CREATE TABLE `app_debug` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` tinyint(4) NOT NULL,
  `msg` text NOT NULL,
  `add_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for app_item
-- ----------------------------
DROP TABLE IF EXISTS `app_item`;
CREATE TABLE `app_item` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '0' COMMENT '网站名',
  `page_index` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '首页',
  `path_rule` text CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '小说目录路径规则',
  `name_rule` text CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '小说名规则',
  `category_rule` text COMMENT '分类规则',
  `status_rule` text NOT NULL COMMENT '状态规则',
  `chapter_rule` text CHARACTER SET utf8 COLLATE utf8_unicode_ci COMMENT '章节规则',
  `content_rule` text COMMENT '内容规则',
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `update_time` int(11) NOT NULL DEFAULT '0',
  `create_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='项目';

-- ----------------------------
-- Records of app_item
-- ----------------------------
BEGIN;
INSERT INTO `app_item` VALUES (2, '顶点小说', 'https://www.23us.so/', 'https://www.23us.so/xiaoshuo/(\\d*).(html)', '<h1>(.*)全文阅读</h1>', '<th>小说类别</th>\\n<td>&nbsp;<a href=\".*\">(.*)</a></td>', '<th>小说状态</th>\\n<td>&nbsp;(.*)</td>', '(?U)<td class=\"L\"><a href=\"(.*)\">(.*)</a></td>', '(?iUs)<dd id=\"contents\">(.*)</dd>', -1, 1531647641, 1515035827);
COMMIT;

-- ----------------------------
-- Table structure for app_novel
-- ----------------------------
DROP TABLE IF EXISTS `app_novel`;
CREATE TABLE `app_novel` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `from_id` int(11) DEFAULT NULL,
  `name` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL,
  `desc` varchar(500) COLLATE utf8_unicode_ci DEFAULT NULL,
  `author` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL,
  `list` text COLLATE utf8_unicode_ci,
  `day_click` int(11) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  `add_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Table structure for sys_func
-- ----------------------------
DROP TABLE IF EXISTS `sys_func`;
CREATE TABLE `sys_func` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL COMMENT '名称',
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '父ID 0顶级',
  `controller` varchar(100) NOT NULL,
  `action` varchar(100) DEFAULT NULL,
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型：0controller 1action',
  `is_menu` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否是菜单：0不是1是',
  `icon` varchar(100) DEFAULT NULL,
  `desc` varchar(200) DEFAULT NULL COMMENT '介绍',
  `sort` int(10) NOT NULL DEFAULT '0',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：0无效1有效-1软删除',
  `update_time` int(10) DEFAULT NULL,
  `create_time` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `pid` (`pid`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8 COMMENT='权限列表';

-- ----------------------------
-- Records of sys_func
-- ----------------------------
BEGIN;
INSERT INTO `sys_func` VALUES (1, '系统设置', 0, 'sys', 'index', 0, 1, 'fa fa-cog', '系统相关参数设置', 0, 1, 1489429439, 1489429439);
INSERT INTO `sys_func` VALUES (2, '管理员管理', 1, 'sysuser', 'index', 1, 1, 'fa fa-users', '添加、删除、编辑系统管理员的权限。', 0, 1, 1489429439, 1489429439);
INSERT INTO `sys_func` VALUES (3, '系统功能添加', 1, 'sysfunc', 'add', 1, 0, 'glyphicon glyphicon-th', '系统功能添加', 6, 1, 1489429439, 1489429439);
INSERT INTO `sys_func` VALUES (4, '功能管理', 1, 'sysfunc', 'index', 1, 1, '', '功能列表', 7, 1, 1489429439, 1489429439);
INSERT INTO `sys_func` VALUES (5, '系统功能删除', 1, 'sysfunc', 'del', 1, 0, '', '系统功能删除', 8, 1, 1489429439, 1489429439);
INSERT INTO `sys_func` VALUES (6, '添加管理员', 1, 'sysuser', 'add', 1, 0, 'glyphicon glyphicon-user', '添加管理员', 1, 1, 1489429439, 1489429439);
INSERT INTO `sys_func` VALUES (7, '管理员删除', 1, 'sysuser', 'del', 1, 0, '', '管理员删除', 2, 1, 1489429439, 1489429439);
INSERT INTO `sys_func` VALUES (8, '重置管理员密码', 1, 'sysuser', 'repwd', 1, 0, '', '重置管理员密码', 3, 1, 1489429439, 1489429439);
INSERT INTO `sys_func` VALUES (9, '锁定管理员', 1, 'sysuser', 'lock', 1, 0, '', '锁定管理员', 4, 1, 1489429439, 1489429439);
INSERT INTO `sys_func` VALUES (10, '系统功能锁定', 1, 'sysfunc', 'lock', 1, 0, '', '系统功能锁定', 9, 1, 1489429439, 1489429439);
INSERT INTO `sys_func` VALUES (11, '角色管理', 1, 'sysrole', 'index', 1, 1, 'fa fa-users', '系统功能锁定', 10, 1, 1489429439, 1489429439);
INSERT INTO `sys_func` VALUES (12, '添加角色', 1, 'sysrole', 'add', 1, 0, 'fa fa-users', '添加角色', 11, 1, 1489429439, 1489429439);
INSERT INTO `sys_func` VALUES (13, '删除角色', 1, 'sysrole', 'del', 1, 0, 'fa fa-users', '删除角色', 12, 1, 1489429439, 1489429439);
INSERT INTO `sys_func` VALUES (14, '锁定角色', 1, 'sysrole', 'lock', 1, 0, 'fa fa-users', '锁定角色', 13, 1, 1489429439, 1489429439);
INSERT INTO `sys_func` VALUES (15, '功能设置菜单', 1, 'sysfunc', 'setmenu', 1, 0, '', '功能设置菜单', 9, 1, 1489429439, 1489429439);
INSERT INTO `sys_func` VALUES (16, '功能升降序', 1, 'sysfunc', 'sort', 1, 0, '', '功能升降序', 5, 1, 1489429439, 1489429439);
INSERT INTO `sys_func` VALUES (17, '日志管理', 1, 'syslog', 'index', 1, 1, '', '日志管理', 99, 1, 1489429439, 1489429439);
INSERT INTO `sys_func` VALUES (18, '小说网站管理', 0, 'appitem', 'index', 0, 1, 'glyphicon glyphicon-th-large', '项目管理 - 列表', 0, 1, 1531385434, 1514871452);
INSERT INTO `sys_func` VALUES (19, '列表', 18, 'appitem', 'index', 0, 1, '', '网站列表', 0, 1, 1531393703, 1514871585);
INSERT INTO `sys_func` VALUES (20, '添加', 18, 'appitem', 'add', 0, 1, '', '添加网站', 0, 1, 1531393758, 1515033366);
INSERT INTO `sys_func` VALUES (27, '日志管理', 0, 'appdebug', '', 0, 1, 'fa fa-bell', '异常日志管理', 4, 1, 1531449077, 1515041927);
INSERT INTO `sys_func` VALUES (28, '列表', 27, 'appdebug', 'index', 0, 1, '', '', 0, 1, 1515041961, 1515041961);
INSERT INTO `sys_func` VALUES (29, '搜索', 18, 'appitem', 'searchAjax', 0, -1, '', '', 0, 1, 1531393726, 1515382967);
INSERT INTO `sys_func` VALUES (32, '锁定', 18, 'appitem', 'lock', 0, -1, '', '', 0, 1, 1531393687, 1515400747);
INSERT INTO `sys_func` VALUES (33, '删除', 18, 'appitem', 'del', 0, -1, '', '', 0, 1, 1531393695, 1515400760);
INSERT INTO `sys_func` VALUES (38, '小说管理', 0, 'appnovel', '', 0, 1, 'glyphicon glyphicon-file', '小说列表', 0, 1, 1531394032, 1531393922);
INSERT INTO `sys_func` VALUES (39, '列表', 38, 'appnovel', 'index', 0, 1, '', '小说列表', 0, 1, 1531394020, 1531393957);
COMMIT;

-- ----------------------------
-- Table structure for sys_logs
-- ----------------------------
DROP TABLE IF EXISTS `sys_logs`;
CREATE TABLE `sys_logs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `type` tinyint(4) NOT NULL,
  `msg` text NOT NULL,
  `add_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '角色名称',
  `desc` varchar(200) DEFAULT NULL COMMENT '角色介绍',
  `list` text NOT NULL COMMENT '权限列表JSON',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态1有效0无效',
  `update_time` int(10) DEFAULT NULL,
  `create_time` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='管理员权限表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` VALUES (1, '管理员', '系统总管理员', '2,6,7,8,9,16,3,4,5,10,15,11,12,13,14,17,39,19,20,29,32,33,28', 1, 1531394056, 1489429439);
INSERT INTO `sys_role` VALUES (2, '编辑', '普通编辑人员', '2,14,15,17,18,23,32', 1, 1489429439, 1489429439);
COMMIT;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(150) NOT NULL COMMENT '登录名',
  `password` varchar(32) NOT NULL COMMENT '密码',
  `nick` varchar(50) DEFAULT NULL COMMENT '昵称',
  `sex` tinyint(4) DEFAULT '0' COMMENT '1男0女',
  `mail` varchar(150) DEFAULT NULL COMMENT '邮箱',
  `tel` varchar(11) DEFAULT NULL COMMENT '手机号',
  `roleid` int(11) DEFAULT '0' COMMENT '所属角色',
  `status` tinyint(4) DEFAULT '1' COMMENT '状体1有效0无效',
  `update_time` int(10) DEFAULT NULL,
  `create_time` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='管理员';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` VALUES (1, 'admin', '4297f44b13955235245b2497399d7a93', '管理员', 1, '627293072@qq.com', '13000000000', 1, 1, 1531386175, 1489429439);
INSERT INTO `sys_user` VALUES (2, 'guest', '4297f44b13955235245b2497399d7a93', 'guest', 1, '13800138000@qq.com', '13800138000', 2, 1, 1531386195, 1489429439);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
