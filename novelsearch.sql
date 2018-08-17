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

 Date: 02/08/2018 15:57:09
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
  `is_official` tinyint(4) DEFAULT '0' COMMENT '是否是官方网站(0,否;1,是)',
  `page_charset` varchar(5) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '搜索关键字编码',
  `page_index` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '首页',
  `page_index_rule` text NOT NULL COMMENT '首页爬取规则',
  `path_rule` text CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '小说目录路径规则',
  `path_tpl` varchar(150) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '小说目录路径模版',
  `path_page_exp` varchar(100) NOT NULL COMMENT '目录样板页',
  `name_rule` text CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '小说名规则',
  `desc_rule` text NOT NULL COMMENT '简介规则',
  `author_rule` text NOT NULL COMMENT '作者规则',
  `category_rule` text NOT NULL COMMENT '分类规则',
  `status_rule` text NOT NULL COMMENT '状态规则',
  `status_end_mark` varchar(10) NOT NULL DEFAULT '' COMMENT '完本标示',
  `chapter_path_rule` text NOT NULL COMMENT '章节路径规则',
  `chapter_path_exp` varchar(100) NOT NULL COMMENT '章节路径样本页',
  `chapter_list_rule` text CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '章节列表规则',
  `content_exp` text NOT NULL COMMENT '内容样本页',
  `content_rule` text NOT NULL COMMENT '内容规则',
  `soso_exp` varchar(150) NOT NULL COMMENT '搜索样本页',
  `soso_page_args` varchar(10) NOT NULL COMMENT '搜索分页参数',
  `soso_rule` text NOT NULL COMMENT '搜索页面规则',
  `spider_exp` varchar(100) NOT NULL COMMENT '爬取页',
  `spider_range` varchar(50) NOT NULL COMMENT '范围',
  `spider_rule` text NOT NULL COMMENT '爬取规则',
  `spider_progress` int(11) NOT NULL DEFAULT '0' COMMENT '爬取进度',
  `err_msg` varchar(100) NOT NULL DEFAULT '' COMMENT '记录错误信息',
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `cron_up_time` int(11) NOT NULL COMMENT '计划任务更新时间',
  `update_time` int(11) NOT NULL DEFAULT '0',
  `create_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 COMMENT='项目';

-- ----------------------------
-- Records of app_item
-- ----------------------------
BEGIN;
INSERT INTO `app_item` VALUES (2, '顶点小说', 0, 'utf8', 'https://www.23us.so/', '<li><p>.*</p><a href=\"https://www.23us.so/xiaoshuo/(.*).html\" target=\"_blank\">(.*)</a></li>\r\n', 'https://www.23us.so/xiaoshuo/\\d*.html', 'https://www.23us.so/xiaoshuo/{$ID}.html', 'https://www.23us.so/xiaoshuo/43.html', '<h1>(.*)全文阅读</h1>', '(?Uis)<p>&nbsp;&nbsp;&nbsp;&nbsp;(.*)<br />', '<th>小说作者</th>\\n<td>&nbsp;(.*)</td>', '<th>小说类别</th>\\n<td>&nbsp;<a href=\".*\">(.*)</a></td>', '<th>小说状态</th>\\n<td>&nbsp;(.*)</td>', '完本', '<a class=\"read\" href=\"(.*)\" title=\".*最新章节\">最新章节</a>', 'https://www.23us.so/files/article/html/0/43/index.html', '(?U)<td class=\"L\"><a href=\"(.*)\">(.*)</a></td>', 'https://www.23us.so/files/article/html/0/43/3615670.html', '(?iUs)<dd id=\"contents\">(.*)</dd>', 'http://zhannei.baidu.com/cse/search?s=8053757951023821596&q={$KEYWORD}', 'p', '(?iU)<a cpos=\"title\" href=\"http://.*/html/\\d+/(.*)/index.html\" title=\"(.*)\" class=\"result-game-item-title-link\" target=\"_blank\">', 'https://www.23us.so/top/allvisit_{$RANGE}.html', '1,1200', '(?Uis)<td class=\"L\"><a href=\"https://www.23us.so/xiaoshuo/(\\d*).html\">(.*)</a></td>', 108, '', 1, 0, 1533028692, 1515035827);
INSERT INTO `app_item` VALUES (6, '58小说网', 0, 'gbk', 'http://www.5858xs.com/', '<a href=\"/pinglun/(.*).html\" target=\"_blank\">(.*)</a>', 'http://www.5858xs.com/(\\d*).html', 'http://www.5858xs.com/{$ID}.html', 'http://www.5858xs.com/273530.html', '<h1><b>(.*)</b></h1>', '(?Uis)<strong>.*</strong>(.*)本站www', '<li>作者：<span>(.*) </span></li>', '(?U)<li>类别：(.*)</li>', '<li>.*状态：(.*)中</li>', '完本', '<li><a href=(.*) title=点击阅读.*>点击阅读</a></li>', 'http://www.5858xs.com/html/273/273530/index.html', '(?Uis)<td><a href=\"(.*\\.html)\">(.*)</a></td>', 'http://www.5858xs.com/html/273/273530/40723695.html', '(?Uis)<div Id=content>.*</script></div></br>(.*)5858xs\\.com.*</div>', 'http://www.5858xs.com/modules/article/search.php?searchtype=articlename&searchkey={$KEYWORD}', 'page', '(?U)<tr>\\s*<td class=odd><a href=http://www.5858xs.com/(.*).html target=_blank><b>(.*)</b></a></td>\r\n', 'http://www.5858xs.com/xiaoshuosort0/0/{$RANGE}.html', '1,100', '(?Uis)<td class=odd><a href=/(.*).html target=_blank>(.*)</a></td>', 68, '', 1, 0, 1533028309, 1531706864);
INSERT INTO `app_item` VALUES (9, '起点', 1, 'utf8', 'https://www.qidian.com/', '(?U)<h4><a href=\"//(.*)\" target=\"_blank\" data-eid=\".*\" data-bid=\".*\">(.*)</a></h4>\r\n(?U)<div class=\"name-box\"><a class=\"name\" href=\"//(.*)\" target=\"_blank\" data-eid=\".*\" data-bid=\".*\">(.*)</a><i class=\"total\">.*</i></div>', 'https://book.qidian.com/info/([\\d]*)', 'https://book.qidian.com/info/{{$ID}}', 'https://book.qidian.com/info/1004608738', '(?U)<div class=\"book-info\\s*\">\\s*<h1>\\s*<em>(.*)</em>\\s*<span>', '(?U)<div class=\"book-intro\">(.*)</div>', '(?U)<a class=\"writer\" href=\"//.*\" target=\"_blank\" data-eid=\".*\">(.*)</a>', '(?Usi)</em>\\s*<a href=\"//.*\" target=\"_blank\">(.*)</a><em class=\"iconfont\">', '(?U)<p class=\"tag\"><span class=\"blue\">(.*)</span>', '完本', '', 'https://book.qidian.com/info/1010327039#Catalog', '(?Uis)<li data-rid=\"\\d*\"><a href=\"//(.*)\" target=\"_blank\" data-eid=\".*\" data-cid=\".*\" title=\".*\">(.*)</a>\\s*</li>', 'https://read.qidian.com/chapter/VyR3pTTz0W7ywypLIF-xfQ2/vDlmIPJ20R1p4rPq4Fd4KQ2', '(?Uis)<div class=\"read-content j_readContent\">(.*)</div>', 'https://www.qidian.com/search?kw={$KEYWORD}', 'page', '(?isU)<h4><a href=\"//(.*)\" target=\"_blank\" data-eid=\".*\" data-bid=\".*\" data-algrid=\".*\">.*</a></h4>', 'https://www.qidian.com/all?pageSize=20&page={$RANGE}', '1,1000000', '(?Uis)<h4><a href=\"//(.*)\" target=\"_blank\" data-eid=\".*\" data-bid=\".*\">(.*)</a></h4>', 0, '', -1, 0, 1532942939, 1531804748);
INSERT INTO `app_item` VALUES (11, '35文学', 0, 'gbk', 'https://www.35wx.com/', '<dt><span>.*</span><a href=\"https://www.35wx.com/book/(\\d*)/\">(.*)</a></dt>\r\n', 'https://www.35wx.com/book/\\d*/', 'https://www.35wx.com/book/{$ID}/', 'https://www.35wx.com/book/3765/', '', '', '', '', '', '完本', '', '', '', '', '', '', 'page', '', '', '1,100', '', 0, '', 0, 0, 1533104267, 1533104267);
COMMIT;

-- ----------------------------
-- Table structure for app_novel
-- ----------------------------
DROP TABLE IF EXISTS `app_novel`;
CREATE TABLE `app_novel` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `from_id` int(11) NOT NULL COMMENT '来源ID',
  `name` varchar(50) COLLATE utf8_unicode_ci NOT NULL COMMENT '小说名',
  `url` text COLLATE utf8_unicode_ci NOT NULL COMMENT '来源地址',
  `unique_id` char(32) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '唯一标示',
  `desc` text COLLATE utf8_unicode_ci NOT NULL COMMENT '简介',
  `category` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '小说分类',
  `author` varchar(50) COLLATE utf8_unicode_ci NOT NULL COMMENT '作者',
  `list` longtext COLLATE utf8_unicode_ci NOT NULL COMMENT '列表',
  `chapter_num` int(11) NOT NULL COMMENT '章节数量',
  `last_chapter` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '最新章节名',
  `last_chapter_url` varchar(200) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '最新章节url',
  `book_status` tinyint(5) DEFAULT '0' COMMENT '书的状态(0,连载中,1,完本)',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  `update_time` int(11) DEFAULT NULL,
  `create_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `t1` (`from_id`,`name`),
  UNIQUE KEY `t2` (`name`, `author`),
  UNIQUE KEY `unique_id` (`unique_id`),
  INDEX `create_time`(`create_time`),
  KEY `author` (`author`)
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
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8 COMMENT='权限列表';

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
INSERT INTO `sys_func` VALUES (18, '网站管理', 0, 'appitem', 'index', 0, 1, 'glyphicon glyphicon-th-large', '项目管理 - 列表', 2, 1, 1532083623, 1514871452);
INSERT INTO `sys_func` VALUES (19, '列表', 18, 'appitem', 'index', 0, 1, '', '网站列表', 0, 1, 1531393703, 1514871585);
INSERT INTO `sys_func` VALUES (20, '添加', 18, 'appitem', 'add', 0, 1, '', '添加网站', 0, 1, 1531393758, 1515033366);
INSERT INTO `sys_func` VALUES (27, '日志管理', 0, 'appdebug', '', 0, 1, 'fa fa-bell', '异常日志管理', 4, 1, 1531449077, 1515041927);
INSERT INTO `sys_func` VALUES (28, '列表', 27, 'appdebug', 'index', 0, 1, '', '', 0, 1, 1515041961, 1515041961);
INSERT INTO `sys_func` VALUES (29, '搜索', 18, 'appitem', 'searchAjax', 0, -1, '', '', 0, 1, 1531393726, 1515382967);
INSERT INTO `sys_func` VALUES (32, '锁定', 18, 'appitem', 'lock', 0, -1, '', '', 0, 1, 1531393687, 1515400747);
INSERT INTO `sys_func` VALUES (33, '删除', 18, 'appitem', 'del', 0, -1, '', '', 0, 1, 1531393695, 1515400760);
INSERT INTO `sys_func` VALUES (38, '小说管理', 0, 'appnovel', '', 0, 1, 'glyphicon glyphicon-file', '小说列表', 3, 1, 1531394032, 1531393922);
INSERT INTO `sys_func` VALUES (39, '列表', 38, 'appnovel', 'index', 0, 1, '', '小说列表', 0, 1, 1531394020, 1531393957);
INSERT INTO `sys_func` VALUES (40, '验证功能', 18, 'appitem', 'verify', 0, -1, '', '验证功能', 0, 1, 1531710514, 1531710379);
INSERT INTO `sys_func` VALUES (41, '基本设置', 0, 'syssetting', '', 0, 1, 'glyphicon glyphicon-edit', '基本设置', 1, 1, 1532057224, 1532056813);
INSERT INTO `sys_func` VALUES (42, '基本信息', 41, 'syssetting', 'index', 0, 1, '', '', 0, 1, 1532057439, 1532057071);
INSERT INTO `sys_func` VALUES (43, '邮件设置', 41, 'syssetting', 'mail', 0, 1, '', '', 0, 1, 1532065874, 1532065874);
INSERT INTO `sys_func` VALUES (44, '删除', 38, 'appnovel', 'del', 0, -1, '', '', 0, 1, 1532083601, 1532083601);
INSERT INTO `sys_func` VALUES (45, '全站爬取功能', 18, 'appitem', 'allspider', 0, -1, '', '权限爬取功能', 0, 1, 1532932359, 1532932290);
INSERT INTO `sys_func` VALUES (46, '重爬', 38, 'appnovel', 'spider', 0, -1, '', '重新爬取', 0, 1, 1533135384, 1533135342);
INSERT INTO `sys_func` VALUES (47, '单篇抓取', 18, 'appitem', 'onespider', 0, -1, '', '单篇抓取', 0, 1, 1533192364, 1533192364);
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
-- Table structure for sys_option
-- ----------------------------
DROP TABLE IF EXISTS `sys_option`;
CREATE TABLE `sys_option` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '配置名字',
  `value` text COLLATE utf8_unicode_ci COMMENT '值',
  `update_time` int(11) DEFAULT NULL,
  `create_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of sys_option
-- ----------------------------
BEGIN;
INSERT INTO `sys_option` VALUES (1, 'web_name', '小说搜索', 1532928203, 1532064476);
INSERT INTO `sys_option` VALUES (2, 'web_keyword', '小说关键字', 1532928203, 1532065440);
INSERT INTO `sys_option` VALUES (3, 'web_desc', '小说描述', 1532928203, 1532065440);
INSERT INTO `sys_option` VALUES (4, 'web_stat', '<script>\r\nvar _hmt = _hmt || [];\r\n(function() {\r\n  var hm = document.createElement(\"script\");\r\n  hm.src = \"https://hm.baidu.com/hm.js?76d1a0e1913a4b43c42669b51ef1734b\";\r\n  var s = document.getElementsByTagName(\"script\")[0]; \r\n  s.parentNode.insertBefore(hm, s);\r\n})();\r\n</script>', 1532928203, 1532065440);
INSERT INTO `sys_option` VALUES (5, 'web_notice', '', 1532928203, 1532065440);
INSERT INTO `sys_option` VALUES (6, 'mail_host', 'smtp.qq.com', 1532850274, 1532066709);
INSERT INTO `sys_option` VALUES (7, 'mail_port', '25', 1532850274, 1532066709);
INSERT INTO `sys_option` VALUES (8, 'mail_user', '627293072@qq.com', 1532850274, 1532066709);
INSERT INTO `sys_option` VALUES (9, 'mail_pwd', 'jnigycrzraecbfci', 1532850274, 1532066709);
INSERT INTO `sys_option` VALUES (10, 'web_notice_mail', 'midoks@163.com', 1532928203, 1532067228);
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
INSERT INTO `sys_role` VALUES (1, '管理员', '系统总管理员', '2,6,7,8,9,16,3,4,5,10,15,11,12,13,14,17,42,43,19,20,29,32,33,40,45,47,39,44,46,28', 1, 1533192391, 1489429439);
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
INSERT INTO `sys_user` VALUES (1, 'admin', '21232f297a57a5a743894a0e4a801fc3', '管理员', 1, 'xxx@163.com', '13000000000', 1, 1, 1532951331, 1489429439);
INSERT INTO `sys_user` VALUES (2, 'guest', '4297f44b13955235245b2497399d7a93', 'guest', 1, '13800138000@qq.com', '13800138000', 2, 1, 1531386195, 1489429439);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
