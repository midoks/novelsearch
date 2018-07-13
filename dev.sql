-- phpMyAdmin SQL Dump
-- version 4.1.14.4
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: 2018-01-11 19:44:12
-- 服务器版本： 5.6.24
-- PHP Version: 5.5.29

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `o_webcron`
--

-- --------------------------------------------------------

--
-- 表的结构 `app_auth`
--

DROP TABLE IF EXISTS `app_auth`;
CREATE TABLE IF NOT EXISTS `app_auth` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) DEFAULT NULL COMMENT '授权用户',
  `item_id` int(11) DEFAULT NULL COMMENT '授权项目',
  `create_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `app_cron`
--

DROP TABLE IF EXISTS `app_cron`;
CREATE TABLE IF NOT EXISTS `app_cron` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '0',
  `desc` varchar(50) NOT NULL DEFAULT '0',
  `item_id` int(11) NOT NULL DEFAULT '0',
  `cron_spec` varchar(100) NOT NULL DEFAULT '0',
  `cmd` text NOT NULL,
  `concurrent` tinyint(4) DEFAULT NULL,
  `exec_num` int(11) NOT NULL DEFAULT '0',
  `prev_time` int(11) NOT NULL DEFAULT '0',
  `notify` tinyint(4) NOT NULL DEFAULT '0',
  `timeout` int(11) NOT NULL DEFAULT '0',
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `update_time` int(11) NOT NULL DEFAULT '0',
  `create_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='运行任务' AUTO_INCREMENT=3 ;

--
-- 转存表中的数据 `app_cron`
--

INSERT INTO `app_cron` (`id`, `name`, `desc`, `item_id`, `cron_spec`, `cmd`, `concurrent`, `exec_num`, `prev_time`, `notify`, `timeout`, `status`, `update_time`, `create_time`) VALUES
(1, 'ccc', 'cccc', 2, '*/2 * * * * ?', 'echo 12', 0, 3604, 1515670988, 0, 10, 1, 1515670198, 0),
(2, '测试任务', '测试任务', 4, '*/3 * * * * ?', 'ls', 0, 7282, 1515670989, 0, 10, 1, 1515670184, 1515378291);

-- --------------------------------------------------------

--
-- 表的结构 `app_cron_log`
--

DROP TABLE IF EXISTS `app_cron_log`;
CREATE TABLE IF NOT EXISTS `app_cron_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `cron_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '任务ID',
  `output` mediumtext NOT NULL COMMENT '任务输出',
  `error` text NOT NULL COMMENT '错误信息',
  `status` tinyint(4) NOT NULL COMMENT '状态',
  `process_time` int(11) NOT NULL DEFAULT '0' COMMENT '消耗时间/毫秒',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `app_debug`
--

DROP TABLE IF EXISTS `app_debug`;
CREATE TABLE IF NOT EXISTS `app_debug` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` tinyint(4) NOT NULL,
  `msg` text NOT NULL,
  `add_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `app_item`
--

DROP TABLE IF EXISTS `app_item`;
CREATE TABLE IF NOT EXISTS `app_item` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '0',
  `desc` varchar(100) NOT NULL DEFAULT '0',
  `type` tinyint(4) NOT NULL DEFAULT '0',
  `sign` char(32) DEFAULT NULL,
  `server_id` int(11) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `update_time` int(11) NOT NULL DEFAULT '0',
  `create_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='项目' AUTO_INCREMENT=6 ;

--
-- 转存表中的数据 `app_item`
--

INSERT INTO `app_item` (`id`, `name`, `desc`, `type`, `sign`, `server_id`, `status`, `update_time`, `create_time`) VALUES
(2, 'test', '测试', 1, '2', 0, 1, 1515474304, 1515035827),
(4, '测试', '192.168.57.91', 0, '1ad48a0957f19d5f1f110448d3fb7a4b', 1, 1, 1515640581, 1515035914);

-- --------------------------------------------------------

--
-- 表的结构 `app_server`
--

DROP TABLE IF EXISTS `app_server`;
CREATE TABLE IF NOT EXISTS `app_server` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ip` varchar(50) NOT NULL DEFAULT '0',
  `port` int(11) NOT NULL,
  `type` tinyint(4) NOT NULL,
  `desc` varchar(255) NOT NULL,
  `user` varchar(50) NOT NULL DEFAULT '0',
  `pwd` varchar(50) NOT NULL DEFAULT '0',
  `status` tinyint(4) NOT NULL,
  `update_time` int(11) NOT NULL DEFAULT '0',
  `create_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='服务器管理' AUTO_INCREMENT=2 ;

--
-- 转存表中的数据 `app_server`
--

INSERT INTO `app_server` (`id`, `ip`, `port`, `type`, `desc`, `user`, `pwd`, `status`, `update_time`, `create_time`) VALUES
(1, '192.168.57.91', 22, 0, '测试服务器', 'chenjiangshan', 'cjscjs123', 1, 1515670951, 0);

-- --------------------------------------------------------

--
-- 表的结构 `sys_func`
--

DROP TABLE IF EXISTS `sys_func`;
CREATE TABLE IF NOT EXISTS `sys_func` (
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
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='权限列表' AUTO_INCREMENT=38 ;

--
-- 转存表中的数据 `sys_func`
--

INSERT INTO `sys_func` (`id`, `name`, `pid`, `controller`, `action`, `type`, `is_menu`, `icon`, `desc`, `sort`, `status`, `update_time`, `create_time`) VALUES
(1, '系统设置', 0, 'sys', 'index', 0, 1, 'fa fa-cog', '系统相关参数设置', 0, 1, 1489429439, 1489429439),
(2, '管理员管理', 1, 'sysuser', 'index', 1, 1, 'fa fa-users', '添加、删除、编辑系统管理员的权限。', 0, 1, 1489429439, 1489429439),
(3, '系统功能添加', 1, 'sysfunc', 'add', 1, 0, 'glyphicon glyphicon-th', '系统功能添加', 6, 1, 1489429439, 1489429439),
(4, '功能管理', 1, 'sysfunc', 'index', 1, 1, '', '功能列表', 7, 1, 1489429439, 1489429439),
(5, '系统功能删除', 1, 'sysfunc', 'del', 1, 0, '', '系统功能删除', 8, 1, 1489429439, 1489429439),
(6, '添加管理员', 1, 'sysuser', 'add', 1, 0, 'glyphicon glyphicon-user', '添加管理员', 1, 1, 1489429439, 1489429439),
(7, '管理员删除', 1, 'sysuser', 'del', 1, 0, '', '管理员删除', 2, 1, 1489429439, 1489429439),
(8, '重置管理员密码', 1, 'sysuser', 'repwd', 1, 0, '', '重置管理员密码', 3, 1, 1489429439, 1489429439),
(9, '锁定管理员', 1, 'sysuser', 'lock', 1, 0, '', '锁定管理员', 4, 1, 1489429439, 1489429439),
(10, '系统功能锁定', 1, 'sysfunc', 'lock', 1, 0, '', '系统功能锁定', 9, 1, 1489429439, 1489429439),
(11, '角色管理', 1, 'sysrole', 'index', 1, 1, 'fa fa-users', '系统功能锁定', 10, 1, 1489429439, 1489429439),
(12, '添加角色', 1, 'sysrole', 'add', 1, 0, 'fa fa-users', '添加角色', 11, 1, 1489429439, 1489429439),
(13, '删除角色', 1, 'sysrole', 'del', 1, 0, 'fa fa-users', '删除角色', 12, 1, 1489429439, 1489429439),
(14, '锁定角色', 1, 'sysrole', 'lock', 1, 0, 'fa fa-users', '锁定角色', 13, 1, 1489429439, 1489429439),
(15, '功能设置菜单', 1, 'sysfunc', 'setmenu', 1, 0, '', '功能设置菜单', 9, 1, 1489429439, 1489429439),
(16, '功能升降序', 1, 'sysfunc', 'sort', 1, 0, '', '功能升降序', 5, 1, 1489429439, 1489429439),
(17, '日志管理', 1, 'syslog', 'index', 1, 1, '', '日志管理', 99, 1, 1489429439, 1489429439),
(18, '项目管理', 0, 'appitem', 'index', 0, 1, 'glyphicon glyphicon-th-large', '项目管理 - 列表', 2, 1, 1515040212, 1514871452),
(19, '项目列表', 18, 'appitem', 'index', 0, 1, '', '项目列表', 0, 1, 1514970432, 1514871585),
(20, '添加项目', 18, 'appitem', 'add', 0, 1, '', '', 0, 1, 1515033366, 1515033366),
(21, '服务器管理', 0, 'appserver', '', 0, 1, 'glyphicon glyphicon-road', '', 1, 1, 1515040219, 1515040192),
(22, '服务器列表', 21, 'appserver', 'index', 0, 1, '', '', 0, 1, 1515040259, 1515040259),
(23, '添加服务器', 21, 'appserver', 'add', 0, 1, '', '', 0, 1, 1515040299, 1515040299),
(24, '任务管理', 0, 'appcron', '', 0, 1, 'glyphicon glyphicon-refresh', '计划任务管理', 3, 1, 1515041719, 1515041701),
(25, '任务列表', 24, 'appcron', 'index', 0, 1, '', '', 0, 1, 1515041741, 1515041741),
(26, '添加任务', 24, 'appcron', 'add', 0, -1, '', '', 0, 1, 1515041766, 1515041766),
(27, '调试日志管理', 0, 'appdebug', '', 0, 1, 'fa fa-bell', '异常日志管理', 4, 1, 1515042025, 1515041927),
(28, '列表', 27, 'appdebug', 'index', 0, 1, '', '', 0, 1, 1515041961, 1515041961),
(29, '项目搜索', 18, 'appitem', 'searchAjax', 0, -1, '', '', 0, 1, 1515382967, 1515382967),
(30, '服务器锁定功能', 21, 'appserver', 'lock', 0, -1, '', '', 0, 1, 1515400454, 1515400454),
(31, '服务器删除功能', 21, 'appserver', 'del', 0, -1, '', '', 0, 1, 1515400476, 1515400476),
(32, '项目锁定', 18, 'appitem', 'lock', 0, -1, '', '', 0, 1, 1515400747, 1515400747),
(33, '项目删除', 18, 'appitem', 'del', 0, -1, '', '', 0, 1, 1515400760, 1515400760),
(34, '任务锁定', 24, 'appcron', 'lock', 0, -1, '', '', 0, 1, 1515403941, 1515403941),
(35, '任务删除', 24, 'appcron', 'del', 0, -1, '', '', 0, 1, 1515403953, 1515403953),
(36, '服务器搜索', 21, 'appserver', 'searchAjax', 0, -1, '', '', 0, 1, 1515405335, 1515405335),
(37, '执行日志', 24, 'appcronlog', 'index', 0, 1, '', '', 0, 1, 1515491096, 1515491096);

-- --------------------------------------------------------

--
-- 表的结构 `sys_logs`
--

DROP TABLE IF EXISTS `sys_logs`;
CREATE TABLE IF NOT EXISTS `sys_logs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `type` tinyint(4) NOT NULL,
  `msg` text NOT NULL,
  `add_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `sys_role`
--

DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE IF NOT EXISTS `sys_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '角色名称',
  `desc` varchar(200) DEFAULT NULL COMMENT '角色介绍',
  `list` text NOT NULL COMMENT '权限列表JSON',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态1有效0无效',
  `update_time` int(10) DEFAULT NULL,
  `create_time` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='管理员权限表' AUTO_INCREMENT=3 ;

--
-- 转存表中的数据 `sys_role`
--

INSERT INTO `sys_role` (`id`, `name`, `desc`, `list`, `status`, `update_time`, `create_time`) VALUES
(1, '管理员', '系统总管理员', '2,6,7,8,9,16,3,4,5,10,15,11,12,13,14,17,22,23,30,31,36,19,20,29,32,33,25,26,34,35,37,28', 1, 1515491473, 1489429439),
(2, '编辑', '普通编辑人员', '2,14,15,17,18,23,32', 1, 1489429439, 1489429439);

-- --------------------------------------------------------

--
-- 表的结构 `sys_user`
--

DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE IF NOT EXISTS `sys_user` (
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
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='管理员' AUTO_INCREMENT=3 ;

--
-- 转存表中的数据 `sys_user`
--

INSERT INTO `sys_user` (`id`, `username`, `password`, `nick`, `sex`, `mail`, `tel`, `roleid`, `status`, `update_time`, `create_time`) VALUES
(1, 'admin', '21232f297a57a5a743894a0e4a801fc3', '管理员', 1, 'admin@localhost', '13000000000', 1, 1, 1489429439, 1489429439),
(2, 'guest', '084e0343a0486ff05530df6c705c8bb4', 'guest', 1, '13800138000@qq.com', '13800138000', 2, 1, 1514976507, 1489429439);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
