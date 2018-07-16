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

 Date: 16/07/2018 19:58:21
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
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '0' COMMENT '网站名',
  `page_index` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '首页',
  `path_rule` text CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '小说目录路径规则',
  `path_page_exp` varchar(100) NOT NULL COMMENT '目录样板页',
  `name_rule` text CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '小说名规则',
  `author_rule` text NOT NULL COMMENT '作者规则',
  `category_rule` text NOT NULL COMMENT '分类规则',
  `status_rule` text NOT NULL COMMENT '状态规则',
  `chapter_path_rule` text NOT NULL COMMENT '章节路径规则',
  `chapter_path_exp` varchar(100) NOT NULL COMMENT '章节路径样本页',
  `chapter_list_rule` text CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '章节列表规则',
  `content_exp` text NOT NULL COMMENT '内容样本页',
  `content_rule` text NOT NULL COMMENT '内容规则',
  `soso_exp` varchar(150) NOT NULL COMMENT '搜索样本页',
  `soso_kw_charset` varchar(5) NOT NULL COMMENT '搜索关键字编码',
  `soso_rule` text NOT NULL COMMENT '搜索页面规则',
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `update_time` int(11) NOT NULL DEFAULT '0',
  `create_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COMMENT='项目';

-- ----------------------------
-- Records of app_item
-- ----------------------------
BEGIN;
INSERT INTO `app_item` VALUES (2, '顶点小说', 'https://www.23us.so/', 'https://www.23us.so/xiaoshuo/(\\d*).(html)', 'https://www.23us.so/xiaoshuo/43.html', '<h1>(.*)全文阅读</h1>', '', '<th>小说类别</th>\\n<td>&nbsp;<a href=\".*\">(.*)</a></td>', '<th>小说状态</th>\\n<td>&nbsp;(.*)</td>', '<a class=\"read\" href=\"(.*)\" title=\".*最新章节\">最新章节</a>', 'https://www.23us.so/files/article/html/0/43/index.html', '(?U)<td class=\"L\"><a href=\"(.*)\">(.*)</a></td>', 'https://www.23us.so/files/article/html/0/43/3615670.html', '(?iUs)<dd id=\"contents\">(.*)</dd>', '', '', '', -1, 1531720780, 1515035827);
INSERT INTO `app_item` VALUES (6, '58小说网', 'http://www.5858xs.com/', '(?U)/(\\d*).(html)', 'http://www.5858xs.com/273530.html', '', '', '', '', '', 'http://www.5858xs.com/html/273/273530/index.html', '', 'http://www.5858xs.com/html/273/273530/40723761.html', '', '', '', '', 0, 1531720949, 1531706864);
INSERT INTO `app_item` VALUES (7, '顶点小说2', 'https://www.x23us.com/', '', 'https://www.x23us.com/html/68/68255/', '', '', '', '', '', 'https://www.x23us.com/html/68/68255/', '', 'https://www.x23us.com/html/68/68255/', '', 'https://www.x23us.com/modules/article/search.php?searchtype=keywords&searchkey={$KEYWORD}', 'gbk', '(?U)<td class=\"odd\"><a href=\"(.*)\"><b style=\"color:red\">.*</b>.*</a></td>', 0, 1531741362, 1531721940);
INSERT INTO `app_item` VALUES (8, '顶点3', 'https://www.dingdiann.com/', '<a href=\"/(.*)/\">(.*)</a>', 'https://www.dingdiann.com/ddk65139/', '<h1>(.*)</h1>', '<p>作&nbsp;&nbsp;者：(.*)</p>', '<meta property=\"og:novel:category\" content=\"(.*)\"/>', '<meta property=\"og:novel:status\" content=\"(.*)\"/>\r\n', '<meta property=\"og:novel:read_url\" content=\"(.*)\"/>\r\n', 'https://www.dingdiann.com/ddk65139/', '<dd> <a style=\"\" href=\"(.*)\">(.*)</a></dd>', 'https://www.dingdiann.com/ddk65139/3436909.html', '(?Us)<div id=\"content\">(.*)</div>', 'https://www.x23us.com/modules/article/search.php?searchtype=keywords&searchkey={$KEYWORD}', 'utf-8', '<tr><td class=\".*\">', 0, 1531739647, 1531723140);
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
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8 COMMENT='权限列表';

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
INSERT INTO `sys_func` VALUES (40, '验证功能', 18, 'appitem', 'verify', 0, 1, '', '验证功能', 0, 1, 1531710514, 1531710379);
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
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_logs
-- ----------------------------
BEGIN;
INSERT INTO `sys_logs` VALUES (1, 1, 1, '添加Item的ID:6', 1531706864);
INSERT INTO `sys_logs` VALUES (2, 1, 1, '更新Item的ID:6|&{%!s(int=6) 58小说网 http://www.5858xs.com/ 1      %!s(int=0) %!s(int64=1531707769) %!s(int64=1531706864)}', 1531707769);
INSERT INTO `sys_logs` VALUES (3, 1, 1, '添加功能的ID:40|:&{%!s(int=40) 验证功能 %!s(int=18) appitem Verify %!s(int=0) %!s(int=-1)  验证功能 %!s(int=0) %!s(int=0) %!s(int64=1531710379) %!s(int64=1531710379)}', 1531710379);
INSERT INTO `sys_logs` VALUES (4, 1, 1, '功能解锁成功', 1531710485);
INSERT INTO `sys_logs` VALUES (5, 1, 1, '更新功能的ID:40|&{%!s(int=40) 验证功能 %!s(int=18) appitem verify %!s(int=0) %!s(int=1)  验证功能 %!s(int=0) %!s(int=1) %!s(int64=1531710514) %!s(int64=1531710379)}', 1531710514);
INSERT INTO `sys_logs` VALUES (6, 1, 1, '更新Item的ID:6|&{%!s(int=6) 58小说网 http://www.5858xs.com/ http://www.5858xs.com/      %!s(int=0) %!s(int64=1531715706) %!s(int64=1531706864)}', 1531715706);
INSERT INTO `sys_logs` VALUES (7, 1, 1, '更新Item的ID:6|&{%!s(int=6) 58小说网 http://www.5858xs.com/ (?U)http://www.5858xs.com/(.*).(html)$      %!s(int=0) %!s(int64=1531715846) %!s(int64=1531706864)}', 1531715846);
INSERT INTO `sys_logs` VALUES (8, 1, 1, '更新Item的ID:6|&{%!s(int=6) 58小说网 http://www.5858xs.com/ (?U)/(\\d*).(html)      %!s(int=0) %!s(int64=1531716168) %!s(int64=1531706864)}', 1531716168);
INSERT INTO `sys_logs` VALUES (9, 1, 1, '更新Item的ID:2|&{%!s(int=2) 顶点小说 https://www.23us.so/ https://www.23us.so/ https://www.23us.so/xiaoshuo/(\\d*).(html) <h1>(.*)全文阅读</h1> <th>小说类别</th>\\n<td>&nbsp;<a href=\".*\">(.*)</a></td> <th>小说状态</th>\\n<td>&nbsp;(.*)</td> (?U)<td class=\"L\"><a href=\"(.*)\">(.*)</a></td> (?iUs)<dd id=\"contents\">(.*)</dd> %!s(int=-1) %!s(int64=1531716866) %!s(int64=1515035827)}', 1531716866);
INSERT INTO `sys_logs` VALUES (10, 1, 1, '更新Item的ID:2|&{%!s(int=2) 顶点小说 https://www.23us.so/ https://www.23us.so/xiaoshuo/43.html https://www.23us.so/xiaoshuo/(\\d*).(html) <h1>(.*)全文阅读</h1> <th>小说类别</th>\\n<td>&nbsp;<a href=\".*\">(.*)</a></td> <th>小说状态</th>\\n<td>&nbsp;(.*)</td> (?U)<td class=\"L\"><a href=\"(.*)\">(.*)</a></td> (?iUs)<dd id=\"contents\">(.*)</dd> %!s(int=-1) %!s(int64=1531716900) %!s(int64=1515035827)}', 1531716900);
INSERT INTO `sys_logs` VALUES (11, 1, 1, '更新Item的ID:2|&{%!s(int=2) 顶点小说 https://www.23us.so/ https://www.23us.so/xiaoshuo/43.html https://www.23us.so/xiaoshuo/(\\d*).(html) <h1>(.*)全文阅读</h1> <th>小说类别</th>\\n<td>&nbsp;<a href=\".*\">(.*)</a></td> <th>小说状态</th>\\n<td>&nbsp;(.*)</td>   (?U)<td class=\"L\"><a href=\"(.*)\">(.*)</a></td> (?iUs)<dd id=\"contents\">(.*)</dd> %!s(int=-1) %!s(int64=1531718279) %!s(int64=1515035827)}', 1531718279);
INSERT INTO `sys_logs` VALUES (12, 1, 1, '更新Item的ID:2|&{%!s(int=2) 顶点小说 https://www.23us.so/ https://www.23us.so/xiaoshuo/43.html https://www.23us.so/xiaoshuo/(\\d*).(html) <h1>(.*)全文阅读</h1> <th>小说类别</th>\\n<td>&nbsp;<a href=\".*\">(.*)</a></td> <th>小说状态</th>\\n<td>&nbsp;(.*)</td>  https://www.23us.so/files/article/html/0/43/index.html (?U)<td class=\"L\"><a href=\"(.*)\">(.*)</a></td> (?iUs)<dd id=\"contents\">(.*)</dd> %!s(int=-1) %!s(int64=1531718373) %!s(int64=1515035827)}', 1531718373);
INSERT INTO `sys_logs` VALUES (13, 1, 1, '更新Item的ID:2|&{%!s(int=2) 顶点小说 https://www.23us.so/ https://www.23us.so/xiaoshuo/43.html https://www.23us.so/xiaoshuo/(\\d*).(html) <h1>(.*)全文阅读</h1> <th>小说类别</th>\\n<td>&nbsp;<a href=\".*\">(.*)</a></td> <th>小说状态</th>\\n<td>&nbsp;(.*)</td> <a class=\"read\" href=\"(.*)\" title=\".*最新章节\">最新章节</a> https://www.23us.so/files/article/html/0/43/index.html (?U)<td class=\"L\"><a href=\"(.*)\">(.*)</a></td> (?iUs)<dd id=\"contents\">(.*)</dd> %!s(int=-1) %!s(int64=1531718562) %!s(int64=1515035827)}', 1531718562);
INSERT INTO `sys_logs` VALUES (14, 1, 1, '更新Item的ID:2|&{%!s(int=2) 顶点小说 https://www.23us.so/ https://www.23us.so/xiaoshuo/43.html https://www.23us.so/xiaoshuo/(\\d*).(html) <h1>(.*)全文阅读</h1> <th>小说类别</th>\\n<td>&nbsp;<a href=\".*\">(.*)</a></td> <th>小说状态</th>\\n<td>&nbsp;(.*)</td> <a class=\"read\" href=\"(.*)\" title=\".*最新章节\">最新章节</a> https://www.23us.so/files/article/html/0/43/index.html (?U)<td class=\"L\"><a href=\"(.*)\">(.*)</a></td> https://www.23us.so/files/article/html/0/43/3615670.html (?iUs)<dd id=\"contents\">(.*)</dd> %!s(int=-1) %!s(int64=1531719196) %!s(int64=1515035827)}', 1531719196);
INSERT INTO `sys_logs` VALUES (15, 1, 1, '更新Item的ID:2|&{%!s(int=2) 顶点小说 https://www.23us.so/ https://www.23us.so/xiaoshuo/43.html https://www.23us.so/xiaoshuo/(\\d*).(html) <h1>(.*)全文阅读</h1> <th>小说类别</th>\\n<td>&nbsp;<a href=\".*\">(.*)</a></td> <th>小说状态</th>\\n<td>&nbsp;(.*)</td> <a class=\"read\" href=\"(.*)\" title=\".*最新章节\">最新章节</a> https://www.23us.so/files/article/html/0/43/index.html (?U)<td class=\"L\"><a href=\"(.*)\">(.*)</a></td> https://www.23us.so/files/article/html/0/43/3615670.html (?iUs)<dd id=\"contents\">(.*)</dd> %!s(int=-1) %!s(int64=1531720780) %!s(int64=1515035827)}', 1531720780);
INSERT INTO `sys_logs` VALUES (16, 1, 1, '更新Item的ID:6|&{%!s(int=6) 58小说网 http://www.5858xs.com/ http://www.5858xs.com/273530.html (?U)/(\\d*).(html)     http://www.5858xs.com/html/273/273530/index.html  http://www.5858xs.com/html/273/273530/40723761.html  %!s(int=0) %!s(int64=1531720949) %!s(int64=1531706864)}', 1531720949);
INSERT INTO `sys_logs` VALUES (17, 1, 1, '添加Item的ID:7', 1531721940);
INSERT INTO `sys_logs` VALUES (18, 1, 1, '添加Item的ID:8', 1531723140);
INSERT INTO `sys_logs` VALUES (19, 1, 1, '更新Item的ID:8|&{%!s(int=8) 顶点3 https://www.dingdiann.com/ https://www.dingdiann.com/ddk65139/          %!s(int=0) %!s(int64=1531723176) %!s(int64=1531723140)}', 1531723176);
INSERT INTO `sys_logs` VALUES (20, 1, 1, '更新Item的ID:8|&{%!s(int=8) 顶点3 https://www.dingdiann.com/ https://www.dingdiann.com/ddk65139/ <a href=\"/(.*)/\">(.*)</a>         %!s(int=0) %!s(int64=1531723345) %!s(int64=1531723140)}', 1531723345);
INSERT INTO `sys_logs` VALUES (21, 1, 1, '更新Item的ID:8|&{%!s(int=8) 顶点3 https://www.dingdiann.com/ https://www.dingdiann.com/ddk65139/ <a href=\"/(.*)/\">(.*)</a> <h1>(.*)</h1>        %!s(int=0) %!s(int64=1531723558) %!s(int64=1531723140)}', 1531723558);
INSERT INTO `sys_logs` VALUES (22, 1, 1, '更新Item的ID:8|&{%!s(int=8) 顶点3 https://www.dingdiann.com/ https://www.dingdiann.com/ddk65139/ <a href=\"/(.*)/\">(.*)</a> <h1>(.*)</h1> <p>作&nbsp;&nbsp;者：(.*)</p>        %!s(int=0) %!s(int64=1531723908) %!s(int64=1531723140)}', 1531723908);
INSERT INTO `sys_logs` VALUES (23, 1, 1, '更新Item的ID:8|&{%!s(int=8) 顶点3 https://www.dingdiann.com/ https://www.dingdiann.com/ddk65139/ <a href=\"/(.*)/\">(.*)</a> <h1>(.*)</h1> <p>作&nbsp;&nbsp;者：(.*)</p>    https://www.dingdiann.com/ddk65139/  https://www.dingdiann.com/ddk5390/3903128.html  %!s(int=0) %!s(int64=1531724117) %!s(int64=1531723140)}', 1531724117);
INSERT INTO `sys_logs` VALUES (24, 1, 1, '更新Item的ID:8|&{%!s(int=8) 顶点3 https://www.dingdiann.com/ https://www.dingdiann.com/ddk65139/ <a href=\"/(.*)/\">(.*)</a> <h1>(.*)</h1> <p>作&nbsp;&nbsp;者：(.*)</p>    https://www.dingdiann.com/ddk65139/  https://www.dingdiann.com/ddk65139/3436909.html  %!s(int=0) %!s(int64=1531724145) %!s(int64=1531723140)}', 1531724145);
INSERT INTO `sys_logs` VALUES (25, 1, 1, '更新Item的ID:8|&{%!s(int=8) 顶点3 https://www.dingdiann.com/ https://www.dingdiann.com/ddk65139/ <a href=\"/(.*)/\">(.*)</a> <h1>(.*)</h1> <p>作&nbsp;&nbsp;者：(.*)</p>  <meta property=\"og:novel:status\" content=\"(.*)\"/>\r\n  https://www.dingdiann.com/ddk65139/  https://www.dingdiann.com/ddk65139/3436909.html  %!s(int=0) %!s(int64=1531724372) %!s(int64=1531723140)}', 1531724372);
INSERT INTO `sys_logs` VALUES (26, 1, 1, '更新Item的ID:8|&{%!s(int=8) 顶点3 https://www.dingdiann.com/ https://www.dingdiann.com/ddk65139/ <a href=\"/(.*)/\">(.*)</a> <h1>(.*)</h1> <p>作&nbsp;&nbsp;者：(.*)</p> <meta property=\"og:novel:category\" content=\"(.*)\"/> <meta property=\"og:novel:status\" content=\"(.*)\"/>\r\n  https://www.dingdiann.com/ddk65139/  https://www.dingdiann.com/ddk65139/3436909.html  %!s(int=0) %!s(int64=1531724402) %!s(int64=1531723140)}', 1531724402);
INSERT INTO `sys_logs` VALUES (27, 1, 1, '更新Item的ID:8|&{%!s(int=8) 顶点3 https://www.dingdiann.com/ https://www.dingdiann.com/ddk65139/ <a href=\"/(.*)/\">(.*)</a> <h1>(.*)</h1> <p>作&nbsp;&nbsp;者：(.*)</p> <meta property=\"og:novel:category\" content=\"(.*)\"/> <meta property=\"og:novel:status\" content=\"(.*)\"/>\r\n  https://www.dingdiann.com/ddk65139/ <dd> <a style=\"\" href=\"(.*)\">(.*)</a></dd> https://www.dingdiann.com/ddk65139/3436909.html  %!s(int=0) %!s(int64=1531724588) %!s(int64=1531723140)}', 1531724588);
INSERT INTO `sys_logs` VALUES (28, 1, 1, '更新Item的ID:8|&{%!s(int=8) 顶点3 https://www.dingdiann.com/ https://www.dingdiann.com/ddk65139/ <a href=\"/(.*)/\">(.*)</a> <h1>(.*)</h1> <p>作&nbsp;&nbsp;者：(.*)</p> <meta property=\"og:novel:category\" content=\"(.*)\"/> <meta property=\"og:novel:status\" content=\"(.*)\"/>\r\n  https://www.dingdiann.com/ddk65139/ <dd> <a style=\"\" href=\"(.*)\">(.*)</a></dd> https://www.dingdiann.com/ddk65139/3436909.html (?Us)<div id=\"content\">(.*)</div> %!s(int=0) %!s(int64=1531724670) %!s(int64=1531723140)}', 1531724670);
INSERT INTO `sys_logs` VALUES (29, 1, 1, '更新Item的ID:8|&{%!s(int=8) 顶点3 https://www.dingdiann.com/ https://www.dingdiann.com/ddk65139/ <a href=\"/(.*)/\">(.*)</a> <h1>(.*)</h1> <p>作&nbsp;&nbsp;者：(.*)</p> <meta property=\"og:novel:category\" content=\"(.*)\"/> <meta property=\"og:novel:status\" content=\"(.*)\"/>\r\n <meta property=\"og:novel:read_url\" content=\"(.*)\"/>\r\n https://www.dingdiann.com/ddk65139/ <dd> <a style=\"\" href=\"(.*)\">(.*)</a></dd> https://www.dingdiann.com/ddk65139/3436909.html (?Us)<div id=\"content\">(.*)</div> %!s(int=0) %!s(int64=1531727272) %!s(int64=1531723140)}', 1531727272);
INSERT INTO `sys_logs` VALUES (30, 1, 1, '更新Item的ID:7|&{%!s(int=7) 顶点小说2 https://www.x23us.com/ https://www.x23us.com/html/68/68255/       https://www.x23us.com/html/68/68255/  https://www.x23us.com/html/68/68255/  https://www.x23us.com/modules/article/search.php?searchtype=keywords&searchkey={$KEYWORD}  %!s(int=0) %!s(int64=1531735749) %!s(int64=1531721940)}', 1531735749);
INSERT INTO `sys_logs` VALUES (31, 1, 1, '更新Item的ID:8|&{%!s(int=8) 顶点3 https://www.dingdiann.com/ https://www.dingdiann.com/ddk65139/ <a href=\"/(.*)/\">(.*)</a> <h1>(.*)</h1> <p>作&nbsp;&nbsp;者：(.*)</p> <meta property=\"og:novel:category\" content=\"(.*)\"/> <meta property=\"og:novel:status\" content=\"(.*)\"/>\r\n <meta property=\"og:novel:read_url\" content=\"(.*)\"/>\r\n https://www.dingdiann.com/ddk65139/ <dd> <a style=\"\" href=\"(.*)\">(.*)</a></dd> https://www.dingdiann.com/ddk65139/3436909.html (?Us)<div id=\"content\">(.*)</div> https://www.x23us.com/modules/article/search.php?searchtype=keywords&searchkey={$KEYWORD}  %!s(int=0) %!s(int64=1531737720) %!s(int64=1531723140)}', 1531737720);
INSERT INTO `sys_logs` VALUES (32, 1, 1, '更新Item的ID:8|&{%!s(int=8) 顶点3 https://www.dingdiann.com/ https://www.dingdiann.com/ddk65139/ <a href=\"/(.*)/\">(.*)</a> <h1>(.*)</h1> <p>作&nbsp;&nbsp;者：(.*)</p> <meta property=\"og:novel:category\" content=\"(.*)\"/> <meta property=\"og:novel:status\" content=\"(.*)\"/>\r\n <meta property=\"og:novel:read_url\" content=\"(.*)\"/>\r\n https://www.dingdiann.com/ddk65139/ <dd> <a style=\"\" href=\"(.*)\">(.*)</a></dd> https://www.dingdiann.com/ddk65139/3436909.html (?Us)<div id=\"content\">(.*)</div> https://www.x23us.com/modules/article/search.php?searchtype=keywords&searchkey={$KEYWORD}  <tr><td %!s(int=0) %!s(int64=1531739579) %!s(int64=1531723140)}', 1531739579);
INSERT INTO `sys_logs` VALUES (33, 1, 1, '更新Item的ID:8|&{%!s(int=8) 顶点3 https://www.dingdiann.com/ https://www.dingdiann.com/ddk65139/ <a href=\"/(.*)/\">(.*)</a> <h1>(.*)</h1> <p>作&nbsp;&nbsp;者：(.*)</p> <meta property=\"og:novel:category\" content=\"(.*)\"/> <meta property=\"og:novel:status\" content=\"(.*)\"/>\r\n <meta property=\"og:novel:read_url\" content=\"(.*)\"/>\r\n https://www.dingdiann.com/ddk65139/ <dd> <a style=\"\" href=\"(.*)\">(.*)</a></dd> https://www.dingdiann.com/ddk65139/3436909.html (?Us)<div id=\"content\">(.*)</div> https://www.x23us.com/modules/article/search.php?searchtype=keywords&searchkey={$KEYWORD} utf-8 <tr><td class=\".*\"> %!s(int=0) %!s(int64=1531739647) %!s(int64=1531723140)}', 1531739647);
INSERT INTO `sys_logs` VALUES (34, 1, 1, '更新Item的ID:7|&{%!s(int=7) 顶点小说2 https://www.x23us.com/ https://www.x23us.com/html/68/68255/       https://www.x23us.com/html/68/68255/  https://www.x23us.com/html/68/68255/  https://www.x23us.com/modules/article/search.php?searchtype=keywords&searchkey={$KEYWORD} utf-8 <tr>\r\n    <td class=\"odd\"><a href=\"https://www.x23us.com/book/69626\"><b style=\"color:red\">凡人</b>修仙之仙界篇</a></td> %!s(int=0) %!s(int64=1531740096) %!s(int64=1531721940)}', 1531740096);
INSERT INTO `sys_logs` VALUES (35, 1, 1, '更新Item的ID:7|&{%!s(int=7) 顶点小说2 https://www.x23us.com/ https://www.x23us.com/html/68/68255/       https://www.x23us.com/html/68/68255/  https://www.x23us.com/html/68/68255/  https://www.x23us.com/modules/article/search.php?searchtype=keywords&searchkey={$KEYWORD} utf-8 <tr>\\n<td class=\".*\"><a href=\"(.*)\"><b style=\"color:red\">凡人</b>修仙之仙界篇</a></td> %!s(int=0) %!s(int64=1531740539) %!s(int64=1531721940)}', 1531740539);
INSERT INTO `sys_logs` VALUES (36, 1, 1, '更新Item的ID:7|&{%!s(int=7) 顶点小说2 https://www.x23us.com/ https://www.x23us.com/html/68/68255/       https://www.x23us.com/html/68/68255/  https://www.x23us.com/html/68/68255/  https://www.x23us.com/modules/article/search.php?searchtype=keywords&searchkey={$KEYWORD} utf-8 (?U)<td class=\"odd\"><a href=\"(.*)\"><b style=\"color:red\">.*</b>.*</a></td> %!s(int=0) %!s(int64=1531740922) %!s(int64=1531721940)}', 1531740922);
INSERT INTO `sys_logs` VALUES (37, 1, 1, '更新Item的ID:7|&{%!s(int=7) 顶点小说2 https://www.x23us.com/ https://www.x23us.com/html/68/68255/       https://www.x23us.com/html/68/68255/  https://www.x23us.com/html/68/68255/  https://www.x23us.com/modules/article/search.php?searchtype=keywords&searchkey={$KEYWORD} gbk (?U)<td class=\"odd\"><a href=\"(.*)\"><b style=\"color:red\">.*</b>.*</a></td> %!s(int=0) %!s(int64=1531741362) %!s(int64=1531721940)}', 1531741362);
COMMIT;

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
