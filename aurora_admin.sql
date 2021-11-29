/*
 Navicat Premium Data Transfer

 Source Server         : myServer
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : 47.111.153.65:3306
 Source Schema         : aurora_admin

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 29/11/2021 15:00:54
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for api
-- ----------------------------
DROP TABLE IF EXISTS `api`;
CREATE TABLE `api` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `method` varchar(255) NOT NULL,
  `module` int(11) DEFAULT NULL,
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of api
-- ----------------------------
BEGIN;
INSERT INTO `api` VALUES (1, '/permissions/getMenuList', '获取菜单数据', 'GET', 3, '2021-10-29 14:15:08', '2021-10-29 14:15:08');
INSERT INTO `api` VALUES (3, '/permissions/addMenuRouter', '新增菜单', 'POST', 3, '2021-10-29 14:20:52', '2021-10-29 14:20:52');
INSERT INTO `api` VALUES (4, '/permissions/GetSingleMenuRouter', '获取单条菜单数据', 'GET', 3, '2021-10-29 14:22:21', '2021-10-29 14:22:21');
INSERT INTO `api` VALUES (5, '/permissions/updateMenuRouter', '更新菜单数据', 'POST', 3, '2021-10-29 14:22:45', '2021-10-29 14:23:30');
INSERT INTO `api` VALUES (6, '/permissions/deleteMenuRouter', '删除菜单', 'GET', 3, '2021-10-29 14:23:16', '2021-10-29 14:23:16');
INSERT INTO `api` VALUES (17, '/permissions/createRole', '创建角色', 'POST', 4, '2021-10-29 16:37:20', '2021-10-29 16:37:20');
INSERT INTO `api` VALUES (18, '/permissions/getRolesList', '获取角色列表', 'GET', 4, '2021-10-29 16:38:31', '2021-10-29 16:38:31');
INSERT INTO `api` VALUES (19, '/permissions/updateRole', '更新角色数据', 'POST', 4, '2021-10-29 16:39:02', '2021-10-29 16:39:02');
INSERT INTO `api` VALUES (20, '/permissions/deleteRole', '删除角色', 'GET', 4, '2021-10-29 16:39:38', '2021-10-29 16:39:38');
INSERT INTO `api` VALUES (21, '/permissions/createUser', '创建用户', 'POST', 5, '2021-10-29 16:40:07', '2021-10-29 16:40:07');
INSERT INTO `api` VALUES (22, '/permissions/updateUser', '更新用户数据', 'POST', 5, '2021-10-29 16:40:29', '2021-10-29 16:40:29');
INSERT INTO `api` VALUES (23, '/permissions/updateUserState', '设置用户账号状态', 'POST', 5, '2021-10-29 16:41:22', '2021-10-29 16:41:22');
INSERT INTO `api` VALUES (24, '/permissions/deleteUser', '删除用户', 'GET', 5, '2021-10-29 16:42:12', '2021-10-29 16:42:12');
INSERT INTO `api` VALUES (25, '/permissions/queryUser', '获取用户列表', 'POST', 5, '2021-10-29 16:42:42', '2021-10-29 16:42:42');
INSERT INTO `api` VALUES (26, '/permissions/getApiList', '获取api列表', 'GET', 21, '2021-10-29 16:46:05', '2021-10-29 16:46:05');
INSERT INTO `api` VALUES (27, '/permissions/createApi', '创建api', 'POST', 21, '2021-10-29 16:46:35', '2021-10-29 16:46:35');
INSERT INTO `api` VALUES (28, '/permissions/deleteApi', '删除api', 'POST', 21, '2021-10-29 16:46:56', '2021-10-29 16:46:56');
INSERT INTO `api` VALUES (29, '/permissions/updateApi', '更新用户数据', 'POST', 21, '2021-10-29 16:47:35', '2021-10-29 16:47:35');
INSERT INTO `api` VALUES (30, '/permissions/getQiNiuToken', '获取七牛token（必选）', 'GET', 1, '2021-11-03 13:56:55', '2021-11-03 14:14:44');
INSERT INTO `api` VALUES (31, '/permissions/getMenuRouter', '获取用户可访问的菜单（必选）', 'GET', 1, '2021-11-03 13:59:47', '2021-11-03 14:14:51');
INSERT INTO `api` VALUES (32, '/permissions/getUserData', '获取当前登录用户信息（必选）', 'GET', 1, '2021-11-03 14:12:56', '2021-11-03 14:14:58');
COMMIT;

-- ----------------------------
-- Table structure for menu_router
-- ----------------------------
DROP TABLE IF EXISTS `menu_router`;
CREATE TABLE `menu_router` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_number` int(11) NOT NULL DEFAULT '0',
  `path` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `redirect` varchar(255) DEFAULT NULL,
  `component` varchar(255) NOT NULL,
  `parent_id` int(11) DEFAULT '0',
  `title` varchar(255) DEFAULT NULL,
  `en_title` varchar(255) DEFAULT NULL,
  `icon` varchar(255) DEFAULT NULL,
  `keep` tinyint(2) NOT NULL DEFAULT '0',
  `hidden` tinyint(2) NOT NULL DEFAULT '0',
  `affix` tinyint(2) NOT NULL DEFAULT '0',
  `nested` tinyint(2) DEFAULT '0',
  `state` int(255) DEFAULT '1',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of menu_router
-- ----------------------------
BEGIN;
INSERT INTO `menu_router` VALUES (1, 99, 'home', 'home', '/home', 'home', 0, '首页', 'Home', 'icon-shouyefill', 0, 0, 1, 0, 1, '2021-10-18 15:29:17', '2021-11-03 16:34:37');
INSERT INTO `menu_router` VALUES (2, 97, '/auth', 'auth', '/auth/menu-settings', 'layout', 0, '权限管理', 'Permission', 'icon-quanxian1', 0, 0, 0, 1, 1, '2021-10-18 16:17:08', '2021-10-25 17:19:44');
INSERT INTO `menu_router` VALUES (3, 10, 'menu-settings', 'menu-settings', '', 'auth/menu-settings', 2, '菜单设置', 'Menu settings', '', 1, 0, 0, 0, 1, '2021-10-18 16:19:32', '2021-11-03 14:41:15');
INSERT INTO `menu_router` VALUES (4, 9, 'role-settings', 'role-settings', '', 'auth/role-settings', 2, '角色设置', 'Role settings', '', 0, 0, 0, 0, 1, '2021-10-18 16:21:20', '2021-11-03 14:41:38');
INSERT INTO `menu_router` VALUES (5, 8, 'user-settings', 'user-settings', '', 'auth/user-settings', 2, '用户设置', 'User settings', '', 1, 0, 0, 0, 1, '2021-10-18 16:22:22', '2021-11-03 14:41:27');
INSERT INTO `menu_router` VALUES (10, 96, '/components-demo', 'components-demo', '/components-demo/editor', 'layout', 0, '组件', 'Components', 'icon-menu_cjgl', 0, 0, 0, 1, 1, '2021-10-20 09:18:12', '2021-11-03 16:34:17');
INSERT INTO `menu_router` VALUES (11, 12, 'editor', 'editor', '', 'components-demo/editor', 10, '富文本编辑器', 'Editor', '', 1, 0, 0, 0, 1, '2021-10-20 09:20:12', '2021-10-22 15:45:11');
INSERT INTO `menu_router` VALUES (12, 0, 'picture-upload', 'picture-upload', '', 'components-demo/picture-upload', 10, '头像上传', 'Picture upload', '', 0, 0, 0, 0, 1, '2021-10-20 09:22:38', '2021-11-03 16:36:09');
INSERT INTO `menu_router` VALUES (13, 0, 'kanban', 'kanban', '', 'components-demo/kanban', 10, '可拖拽看板', 'Drag panel', '', 1, 0, 0, 0, 1, '2021-10-20 09:25:35', '2021-11-03 16:36:14');
INSERT INTO `menu_router` VALUES (21, 0, 'api-settings', 'api-settings', '', 'auth/api-settings', 2, 'Api管理', 'Api settings', '', 1, 0, 0, 0, 1, '2021-10-28 14:35:17', '2021-11-03 14:41:33');
INSERT INTO `menu_router` VALUES (36, 0, 'user', 'user', '/user', 'user', 0, '个人中心', 'User', '', 0, 1, 0, 0, 1, '2021-11-03 16:37:26', '2021-11-03 16:37:26');
COMMIT;

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_key` varchar(255) NOT NULL,
  `role_name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `menu` varchar(255) NOT NULL,
  `api` varchar(255) NOT NULL,
  `state` int(2) NOT NULL DEFAULT '1',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `role_key` (`role_key`) USING BTREE,
  UNIQUE KEY `role_name` (`role_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of roles
-- ----------------------------
BEGIN;
INSERT INTO `roles` VALUES (1, 'super_admin', '超级管理员', '系统权限最高级别拥有者，请勿随便修改该角色拥有的的权限。', '1,2,3,4,5,21,10,11,12,13,36', '30,31,32,1,3,4,5,6,17,18,19,20,21,22,23,24,25,26,27,28,29', 1, '2021-10-21 16:05:13', '2021-11-03 16:39:27');
INSERT INTO `roles` VALUES (2, 'warrior', '战士', '战士型英雄攻守兼备，属于输出能力和生存能力都比较突出的近攻英雄', '1,2,3,4,5,21,10,11,12,13,36', '30,31,32,1,4,18,25,26', 1, '2021-10-21 17:17:41', '2021-11-03 16:39:34');
INSERT INTO `roles` VALUES (4, 'assassin', '刺客', '刺客型英雄属于生存能力较弱，但输出能力却极具爆炸性的英雄。在团战中可绕往敌方阵容后面，找准时机一套技能秒杀敌方后排的脆皮输出。', '1,10,11,12,13,36', '30,31,32', 1, '2021-10-23 15:21:41', '2021-11-23 05:49:27');
INSERT INTO `roles` VALUES (18, 'master', '法师', '法师型英雄属于生存能力较弱，但输出强力且兼具控制技能的英雄。在团战中占据后排的位置，对敌方英雄进行输出控制。', '1,10,13,36', '30,31,32', 1, '2021-10-25 16:05:22', '2021-11-23 06:17:04');
INSERT INTO `roles` VALUES (19, 'shooter', '射手', '射手型英雄属于远程输出能力极高且带有控制技能的英雄。在团战中利用自身“手长”的优势站于后排进行输出控制。', '1,10,11,36', '30,31,32', 1, '2021-10-25 16:06:18', '2021-11-23 06:17:23');
INSERT INTO `roles` VALUES (20, 'auxiliary', '辅助', '辅助型英雄属于生存能力和输出能力略平庸的英雄，单靠强力的技能效果来为队友增加状态及控制敌方英雄。', '10,11,12,13,36', '30,31,32', 1, '2021-10-25 16:06:59', '2021-11-23 06:17:34');
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) NOT NULL,
  `number` varchar(255) NOT NULL,
  `avatar` varchar(255) NOT NULL DEFAULT 'http://img.joooyoo.com/FjXaEOdLB4y3b27qIuZeVUCGkBXy',
  `phone` varchar(255) DEFAULT NULL,
  `role_key` varchar(255) NOT NULL,
  `password` varchar(255) DEFAULT '123456',
  `sex` int(2) DEFAULT NULL,
  `state` int(11) NOT NULL DEFAULT '1',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_name` (`user_name`) USING BTREE,
  UNIQUE KEY `number` (`number`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, '吕布', 'lvbu', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/123/123.jpg', '18200009999', 'warrior', '123456', 1, 1, '2021-10-25 12:05:13', '2021-10-28 14:14:56');
INSERT INTO `user` VALUES (2, '貂蝉', 'diaochan', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/141/141.jpg', '13200000999', 'assassin', '123456', 2, 1, '2021-10-25 12:08:26', '2021-10-27 09:22:23');
INSERT INTO `user` VALUES (4, '司空震', 'sikongzhen', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/537/537.jpg', '13498762435', 'warrior', '123456', 1, 1, '2021-10-25 15:45:23', '2021-10-25 15:59:22');
INSERT INTO `user` VALUES (5, '蒙恬', 'mengtian', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/527/527.jpg', '15699998888', 'warrior', '123456', 1, 1, '2021-10-25 17:23:53', '2021-10-26 15:40:32');
INSERT INTO `user` VALUES (6, '马超', 'machao', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/518/518.jpg', '13534567890', 'warrior', '123456', 1, 1, '2021-10-25 17:33:09', '2021-10-28 13:27:07');
INSERT INTO `user` VALUES (7, '马可波罗', 'makeboluo', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/132/132.jpg', '15645327896', 'shooter', '123456', 1, 0, '2021-10-25 17:35:25', '2021-10-26 13:03:09');
INSERT INTO `user` VALUES (8, '瑶', 'yao', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/505/505.jpg', '13245678945', 'auxiliary', '123456', 2, 1, '2021-10-25 17:39:59', '2021-10-25 17:48:03');
INSERT INTO `user` VALUES (9, '关羽', 'guanyu', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/140/140.jpg', '15932263166', 'warrior', '123456', 1, 1, '2021-10-25 17:41:49', '2021-10-25 17:47:53');
INSERT INTO `user` VALUES (10, '不知火舞', 'buzhihuowu', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/157/157.jpg', '17978652345', 'assassin', '123456', 2, 1, '2021-10-25 17:43:03', '2021-10-25 17:47:38');
INSERT INTO `user` VALUES (11, '安琪拉', 'anqila', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/142/142.jpg', '18976785643', 'master', '123456', 2, 1, '2021-10-25 17:46:31', '2021-10-25 17:47:30');
INSERT INTO `user` VALUES (12, '蔡文姬', 'caiwenji', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/184/184.jpg', '13456784657', 'auxiliary', '123456', 2, 1, '2021-10-25 17:49:38', '2021-10-25 17:49:49');
INSERT INTO `user` VALUES (13, '澜', 'lan', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/528/528.jpg', '13567894578', 'assassin', '123456', 1, 0, '2021-10-26 09:22:05', '2021-10-27 13:55:32');
INSERT INTO `user` VALUES (14, '镜', 'jing', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/531/531.jpg', '15678765435', 'assassin', '123456', 2, 1, '2021-10-26 09:23:19', '2021-10-26 09:25:09');
INSERT INTO `user` VALUES (15, '露娜', 'luna', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/146/146.jpg', '18230987876', 'warrior', '123456', 2, 1, '2021-10-26 09:24:18', '2021-10-26 09:25:28');
INSERT INTO `user` VALUES (16, '艾琳', 'ailin', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/155/155.jpg', '13456734578', 'shooter', '123456', 2, 1, '2021-10-26 09:26:24', '2021-10-26 09:28:41');
INSERT INTO `user` VALUES (17, '迦罗', 'jialuo', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/508/508.jpg', '17867534256', 'shooter', '123456', 2, 1, '2021-10-26 09:27:12', '2021-10-26 09:28:42');
INSERT INTO `user` VALUES (18, '孙膑', 'sunbin', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/118/118.jpg', '13245678975', 'auxiliary', '123456', 1, 1, '2021-10-26 09:28:05', '2021-10-26 09:28:52');
INSERT INTO `user` VALUES (19, '妲己', 'daji', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/109/109.jpg', '15878976532', 'master', '123456', 2, 1, '2021-10-26 15:31:12', '2021-10-26 15:32:25');
INSERT INTO `user` VALUES (20, '嬴政', 'yinzheng', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/110/110.jpg', '15789873546', 'master', '123456', 1, 1, '2021-10-26 15:33:21', '2021-10-26 15:33:34');
INSERT INTO `user` VALUES (21, '武则天', 'wuzetian', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/136/136.jpg', '18976743298', 'master', '123456', 2, 1, '2021-10-26 15:34:20', '2021-10-27 09:45:04');
INSERT INTO `user` VALUES (22, '晚风', 'wanfeng', 'http://img.joooyoo.com/FjXaEOdLB4y3b27qIuZeVUCGkBXy', '19876543210', 'super_admin', 'wenfangxin', 1, 1, '2021-10-26 15:35:25', '2021-11-03 15:39:41');
INSERT INTO `user` VALUES (23, '宫本武藏', 'gongbenwuzang', 'https://game.gtimg.cn/images/yxzj/img201606/heroimg/130/130.jpg', '17867532222', 'warrior', '12345612', 1, 1, '2021-10-27 14:23:19', '2021-11-03 15:39:25');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
