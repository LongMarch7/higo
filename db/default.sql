--
-- 表的结构 `micro_admin_menu`
--

CREATE TABLE IF NOT EXISTS `micro_admin_menu` (
  `id` int(10) UNSIGNED NOT NULL ,
  `parent_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '父菜单id',
  `type` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '菜单类型;1:有界面可访问菜单,2:无界面可访问菜单,0:只作为菜单',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '状态;1:显示,0:不显示',
  `list_order` float NOT NULL DEFAULT '10000' COMMENT '排序',
  `router` varchar(50) NOT NULL DEFAULT '' COMMENT '规则唯一英文标识',
  `method` varchar(100) NOT NULL DEFAULT '' COMMENT '规则方法(大写)GET、POST、PUT、 (PUT) | (GET)',
  `param` varchar(50) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '额外参数',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `icon` varchar(20) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '菜单图标',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `status` (`status`),
  KEY `parent_id` (`parent_id`),
  KEY `router` (`router`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台菜单表';

--
-- 转存表中的数据 `micro_admin_menu`
--
INSERT INTO `micro_admin_menu` (`id`, `parent_id`, `type`, `status`, `list_order`, `router`, `method`, `param`, `name`, `icon`, `remark`) VALUES
(1,     0,   0, 1, 0,     'admin/setting/SettingIndex',     'GET',    '', '设置',              'cogs', '系统设置入口'),
(100,   1,   1, 1, 100,   'admin/setting/MailIndex',        'GET',    '', '邮箱配置',             '', '邮箱配置'),
(10000, 100, 2, 0, 10000, 'admin/setting/MailConfPOST',     'POST',   '', '邮箱配置提交保存',      '', '邮箱配置提交保存'),
(10001, 100, 1, 0, 10000, 'admin/setting/MailTemplate',     'GET',    '', '邮件模板',             '', '邮件模板'),
(10002, 100, 0, 0, 10000, 'admin/setting/MailTemplatePOST', 'POST',   '', '邮件模板提交',          '', '邮件模板提交'),
(10003, 100, 1, 0, 10000, 'admin/setting/MailTest',         'GET',    '', '邮件测试',             '', '邮件发送测试'),
(10004, 100, 1, 0, 10000, 'admin/setting/MailSendTest',     'POST',   '', '邮件发送测试',          '', '邮件发送测试'),
(101,   1,   0, 1, 101,   'admin/setting/LinkIndex',        'GET',    '', '友情链接',             '', '友情链接管理'),
(10100, 101, 1, 0, 10000, 'admin/setting/LinkAdd',          'GET',    '', '添加友情链接',         '', '添加友情链接'),
(10101, 101, 2, 0, 10000, 'admin/setting/LinkAddPost',      'POST',   '', '添加友情链接提交保存',  '', '添加友情链接提交保存'),
(10102, 101, 1, 0, 10000, 'admin/setting/LinkEdit',         'GET',    '', '编辑友情链接',         '', '编辑友情链接'),
(10103, 101, 2, 0, 10000, 'admin/setting/LinkEditPost',     'POST',   '', '编辑友情链接提交保存',  '', '编辑友情链接提交保存'),
(10104, 101, 2, 0, 10000, 'admin/setting/LinkDelete',       'POST',   '', '删除友情链接',         '', '删除友情链接'),
(10105, 101, 2, 0, 10000, 'admin/setting/LinkListOrder',    'POST',   '', '友情链接排序',         '', '友情链接排序'),
(10106, 101, 2, 0, 10000, 'admin/setting/LinkToggle',       'POST',   '', '友情链接显示隐藏',      '', '友情链接显示隐藏'),
(102,   1,   1, 1, 0,     'admin/setting/SiteIndex',        'GET',    '', '网站信息',             '', '网站信息'),
(10200, 102, 2, 0, 10000, 'admin/setting/SitePost',         'POST',   '', '网站信息设置提交',      '', '网站信息设置提交'),
(103,   1,   1, 0, 10000, 'admin/setting/PasswordIndex',    'GET',    '', '密码修改',             '', '密码修改'),
(10300, 103, 2, 0, 10000, 'admin/setting/PasswordPost',     'POST',   '', '密码修改提交',          '', '密码修改提交'),
(104,   1,   1, 1, 10000, 'admin/setting/UploadIndex',      'GET',    '', '上传设置',             '', '上传设置'),
(10400, 104, 2, 0, 10000, 'admin/setting/UploadPost',       'POST',   '', '上传设置提交',          '', '上传设置提交'),
(105,   1,   1, 0, 10000, 'admin/setting/AssetIndex',       'GET',    '', '资源管理',             'file', '资源管理列表'),
(10500, 105, 2, 0, 10000, 'admin/setting/AssetDel',         'POST',   '', '删除文件',             '', '删除文件'),


(2,     0,   0, 1, 10,    'admin/user/Index',               'GET',    '', '用户管理', 'group', '用户管理'),
(200,   2,   1, 1, 10000, 'admin/user/RoleIndex',           'GET',    '', '角色管理', '', '角色管理'),
(20000, 200, 1, 0, 10000, 'admin/user/RoleAdd',             'GET',    '', '添加角色', '', '添加角色'),
(20001, 200, 2, 0, 10000, 'admin/user/RoleAddPost',         'POST',   '', '添加角色提交', '', '添加角色提交'),
(20002, 200, 1, 0, 10000, 'admin/user/RoleEdit',            'GET',    '', '编辑角色', '', '编辑角色'),
(20003, 200, 2, 0, 10000, 'admin/user/RoleEditPost',        'POST',   '', '编辑角色提交', '', '编辑角色提交'),
(20004, 200, 2, 0, 10000, 'admin/user/RoleDelete',          'POST',   '', '删除角色', '', '删除角色'),
(20005, 200, 1, 0, 10000, 'admin/user/RoleAuthorize',       'GET',    '', '设置角色权限', '', '设置角色权限'),
(20006, 200, 2, 0, 10000, 'admin/user/RoleAuthorizePost',   'POST',   '', '角色授权提交', '', '角色授权提交'),
(201,   2,   1, 1, 10000, 'admin/user/SiteUser',            'GET',    '', '本站用户', '', '本站用户'),
(20100, 201, 2, 0, 10000, 'admin/user/SiteUserBan',         'POST',   '', '本站用户拉黑', '', '本站用户拉黑'),
(20101, 201, 2, 0, 10000, 'admin/user/SiteUserCancelBan',   'POST',   '', '本站用户启用', '', '本站用户启用'),

(3,     0,   0, 1, 30,    'admin/portal/Incdex',            'GET',    '', '门户管理', 'th', '门户管理'),
(300,   3,   1, 1, 10000, 'admin/portal/ArticleIndex',      'GET',    '', '文章管理', '', '文章列表'),
(30000, 300, 1, 0, 10000, 'admin/portal/ArticleAdd',        'GET',    '', '添加文章', '', '添加文章'),
(30001, 300, 2, 0, 10000, 'admin/portal/ArticleAddPost',    'POST',   '', '添加文章提交', '', '添加文章提交'),
(30002, 300, 1, 0, 10000, 'admin/portal/ArticleEdit',       'GET',    '', '编辑文章', '', '编辑文章'),
(30003, 300, 2, 0, 10000, 'admin/portal/ArticleEditPost',   'POST',   '', '编辑文章提交', '', '编辑文章提交'),
(30004, 300, 2, 0, 10000, 'admin/portal/ArticleDelete',     'POST',   '', '文章删除', '', '文章删除'),
(30005, 300, 2, 0, 10000, 'admin/portal/ArticlePublish',    'POST',   '', '文章发布', '', '文章发布'),
(30006, 300, 2, 0, 10000, 'admin/portal/ArticleTop',        'POST',   '', '文章置顶', '', '文章置顶'),
(30007, 300, 2, 0, 10000, 'admin/portal/ArticleRecommend',  'POST',   '', '文章推荐', '', '文章推荐'),
(30008, 300, 2, 0, 10000, 'admin/portal/ArticleOrder',      'POST',   '', '文章排序', '', '文章排序'),
(301,   3,   1, 1, 10000, 'admin/portal/CategoryIndex',     'GET',    '', '分类管理', '', '文章分类列表'),
(30100, 301, 1, 0, 10000, 'admin/portal/CategoryAdd',       'GET',    '', '添加文章分类', '', '添加文章分类'),
(30101, 301, 2, 0, 10000, 'admin/portal/CategoryAddPost',   'POST',   '', '添加文章分类提交', '', '添加文章分类提交'),
(30102, 301, 1, 0, 10000, 'admin/portal/CategoryEdit',      'GET',    '', '编辑文章分类', '', '编辑文章分类'),
(30103, 301, 2, 0, 10000, 'admin/portal/CategoryEditPost',  'POST',   '', '编辑文章分类提交', '', '编辑文章分类提交'),
(30104, 301, 1, 0, 10000, 'admin/portal/CategorySelect',    'GET',    '', '文章分类选择对话框', '', '文章分类选择对话框'),
(30105, 301, 2, 0, 10000, 'admin/portal/CategoryOrder',     'POST',   '', '文章分类排序', '', '文章分类排序'),
(30106, 301, 2, 0, 10000, 'admin/portal/CategoryDelete',    'POST',   '', '删除文章分类', '', '删除文章分类'),
(302,   3,   1, 1, 10000, 'admin/portal/TagIndex',          'GET',    '', '文章标签', '', '文章标签'),
(30200, 302, 1, 0, 10000, 'admin/portal/TagAdd',            'GET',    '', '添加文章标签', '', '添加文章标签'),
(30201, 302, 2, 0, 10000, 'admin/portal/TagAddPost',        'POST',   '', '添加文章标签提交', '', '添加文章标签提交'),
(30202, 302, 2, 0, 10000, 'admin/portal/TagUpStatus',       'POST',   '', '更新标签状态', '', '更新标签状态'),
(30203, 302, 2, 0, 10000, 'admin/portal/TagDelete',         'POST',   '', '删除文章标签', '', '删除文章标签');


-- --------------------------------------------------------

--
-- 表的结构 `micro_asset`
--

CREATE TABLE IF NOT EXISTS `micro_asset` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户id',
  `file_size` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '文件大小,单位B',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '上传时间',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态;1:可用,0:不可用',
  `download_times` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '下载次数',
  `file_key` varchar(64) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '文件惟一码',
  `filename` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '文件名',
  `file_path` varchar(100) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '文件路径,相对于upload目录,可以为url',
  `file_md5` varchar(32) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '文件md5值',
  `file_sha1` varchar(40) CHARACTER SET utf8 NOT NULL DEFAULT '',
  `suffix` varchar(10) NOT NULL DEFAULT '' COMMENT '文件后缀名,不包括点',
  `more` text COMMENT '其它详细信息,JSON格式',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资源表';


--
-- 表的结构 `micro_user_like_201904`
--

CREATE TABLE IF NOT EXISTS `micro_casbin_rule` (
  `p_type` varchar(100) NOT NULL DEFAULT '' COMMENT '规则类型',
  `v0` varchar(100) COMMENT '规则0',
  `v1` varchar(100) COMMENT '规则1',
  `v2` varchar(100) COMMENT '规则2',
  `v3` varchar(100) COMMENT '规则3',
  `v4` varchar(100) COMMENT '规则4',
  `v5` varchar(100) COMMENT '规则4',
  UNIQUE KEY `rule_key` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='授权表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_comment`
--

CREATE TABLE IF NOT EXISTS `micro_comment` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '被回复的评论id',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '发表评论的用户id',
  `to_user_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '被评论的用户id',
  `object_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '评论内容 id',
  `like_count` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '点赞数',
  `dislike_count` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '吐槽数',
  `floor` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '楼层数',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '评论时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态,1:已审核,0:未审核',
  `type` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '评论类型；1实名评论',
  `tab_name` varchar(64) NOT NULL DEFAULT '' COMMENT '评论内容所在表，不带表前缀',
  `full_name` varchar(50) NOT NULL DEFAULT '' COMMENT '评论者昵称',
  `url` text COMMENT '原文地址',
  `content` text CHARACTER SET utf8mb4 COMMENT '评论内容',
  `more` text CHARACTER SET utf8mb4 COMMENT '扩展属性',
  PRIMARY KEY (`id`),
  KEY `table_id_status` (`tab_name`,`object_id`,`status`),
  KEY `object_id` (`object_id`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `parent_id` (`parent_id`) USING BTREE,
  KEY `create_time` (`create_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='评论表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_link`
--

CREATE TABLE IF NOT EXISTS `micro_link` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态;1:显示;0:不显示',
  `rating` int(11) NOT NULL DEFAULT '0' COMMENT '友情链接评级',
  `list_order` float NOT NULL DEFAULT '10000' COMMENT '排序',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '友情链接描述',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '友情链接地址',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '友情链接名称',
  `image` varchar(100) NOT NULL DEFAULT '' COMMENT '友情链接图标',
  `target` varchar(10) NOT NULL DEFAULT '' COMMENT '友情链接打开方式',
  `rel` varchar(50) NOT NULL DEFAULT '' COMMENT '链接与网站的关系',
  PRIMARY KEY (`id`),
  KEY `status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='友情链接表';

--
-- 转存表中的数据 `micro_link`
--

INSERT INTO `micro_link` (`id`, `status`, `rating`, `list_order`, `description`, `url`, `name`, `image`, `target`, `rel`) VALUES
(1, 1, 1, 8, '百度', 'http://www.baidu.com', 'baidu', '', '_blank', '');

-- --------------------------------------------------------

--
-- 表的结构 `micro_portal_category`
--

CREATE TABLE IF NOT EXISTS `micro_portal_category` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `parent_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '分类父id',
  `post_count` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '分类文章数',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态,1:发布,0:不发布',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
  `list_order` float NOT NULL DEFAULT '10000' COMMENT '排序',
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分类名称',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '分类描述',
  `path` varchar(255) NOT NULL DEFAULT '' COMMENT '分类层级关系路径',
  `seo_title` varchar(100) NOT NULL DEFAULT '',
  `seo_keywords` varchar(255) NOT NULL DEFAULT '',
  `seo_description` varchar(255) NOT NULL DEFAULT '',
  `list_tpl` varchar(50) NOT NULL DEFAULT '' COMMENT '分类列表模板',
  `one_tpl` varchar(50) NOT NULL DEFAULT '' COMMENT '分类文章页模板',
  `more` text COMMENT '扩展属性',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='portal应用 文章分类表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_portal_category_post`
--

CREATE TABLE IF NOT EXISTS `micro_portal_category_post` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `post_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '文章id',
  `category_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '分类id',
  `list_order` float NOT NULL DEFAULT '10000' COMMENT '排序',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态,1:发布;0:不发布',
  PRIMARY KEY (`id`),
  KEY `term_taxonomy_id` (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='portal应用 分类文章对应表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_portal_post`
--

CREATE TABLE IF NOT EXISTS `micro_portal_post` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '父级id',
  `post_type` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '类型,1:文章;2:页面',
  `post_format` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '内容格式;1:html;2:md',
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '发表者用户id',
  `post_status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态;1:已发布;0:未发布;',
  `comment_status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '评论状态;1:允许;0:不允许',
  `is_top` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否置顶;1:置顶;0:不置顶',
  `recommended` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否推荐;1:推荐;0:不推荐',
  `post_hits` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '查看数',
  `post_favorites` INT UNSIGNED NOT NULL DEFAULT '0' COMMENT '收藏数' ,
  `post_like` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '点赞数',
  `comment_count` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '评论数',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
  `published_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '发布时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
  `post_title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'post标题',
  `post_keywords` varchar(150) NOT NULL DEFAULT '' COMMENT 'seo keywords',
  `post_excerpt` varchar(500) NOT NULL DEFAULT '' COMMENT 'post摘要',
  `post_source` varchar(150) NOT NULL DEFAULT '' COMMENT '转载文章的来源',
  `thumbnail` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '缩略图',
  `post_content` text COMMENT '文章内容',
  `post_content_filtered` text COMMENT '处理过的文章内容',
  `more` text COMMENT '扩展属性,如缩略图;格式为json',
  PRIMARY KEY (`id`),
  KEY `type_status_date` (`post_type`,`post_status`,`create_time`,`id`),
  KEY `parent_id` (`parent_id`),
  KEY `user_id` (`user_id`),
  KEY `create_time` (`create_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='portal应用 文章表' ROW_FORMAT=COMPACT;

-- --------------------------------------------------------

--
-- 表的结构 `micro_portal_tag`
--

CREATE TABLE IF NOT EXISTS `micro_portal_tag` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '标签id',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态,1:发布,0:不发布',
  `recommended` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否推荐;1:推荐;0:不推荐',
  `post_count` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '标签文章数',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标签名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='portal应用 文章标签表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_portal_tag_post`
--

CREATE TABLE IF NOT EXISTS `micro_portal_tag_post` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `tag_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '标签 id',
  `post_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '文章 id',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态,1:发布;0:不发布',
  PRIMARY KEY (`id`),
  KEY `post_id` (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='portal应用 标签文章对应表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_recycle_bin`
--

CREATE TABLE IF NOT EXISTS `micro_recycle_bin` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `object_id` int(11) DEFAULT '0' COMMENT '删除内容 id',
  `create_time` int(10) UNSIGNED DEFAULT '0' COMMENT '创建时间',
  `tab_name` varchar(60) DEFAULT '' COMMENT '删除内容所在表名',
  `name` varchar(255) DEFAULT '' COMMENT '删除内容名称',
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT=' 回收站';

-- --------------------------------------------------------

--
-- 表的结构 `micro_role`
--

CREATE TABLE IF NOT EXISTS `micro_role` (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '父角色ID',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '状态;0:禁用;1:正常',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
  `list_order` float NOT NULL DEFAULT '0' COMMENT '排序',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `parent_id` (`parent_id`),
  KEY `status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

--
-- 转存表中的数据 `micro_role`
--

INSERT INTO `micro_role` (`id`, `parent_id`, `status`, `create_time`, `update_time`, `list_order`, `name`, `remark`) VALUES
(1, 0, 1, 1329633709, 1329633709, 0, '超级管理员', '拥有网站最高管理员权限！'),
(2, 0, 1, 1329633709, 1329633709, 0, '普通管理员', '权限由最高管理员分配！');

-- --------------------------------------------------------

--
-- 表的结构 `micro_role_user`
--

CREATE TABLE IF NOT EXISTS `micro_role_user` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `role_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '角色 id',
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `role_id` (`role_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户角色对应表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_user`
--

CREATE TABLE IF NOT EXISTS `micro_user` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `bind_id` bigint(20) UNSIGNED NOT NULL  COMMENT '绑定第三方账号ID',
  `user_type` tinyint(3) UNSIGNED NOT NULL DEFAULT '2' COMMENT '用户类型;1:admin;2:大v;3:广告商',
  `sex` tinyint(2) NOT NULL DEFAULT '0' COMMENT '性别;0:保密,1:男,2:女',
  `birthday` int(11) NOT NULL DEFAULT '0' COMMENT '生日',
  `last_login_time` int(11) NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `score` int(11) NOT NULL DEFAULT '0' COMMENT '用户等级',
  `coin` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '金币',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '余额',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '注册时间',
  `user_status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '用户状态;0:禁用,1:正常,2:冻结,3：未验证',
  `user_login` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `user_pass` varchar(64) NOT NULL DEFAULT '' COMMENT '登录密码;micro_password加密',
  `pay_pass` varchar(64) NOT NULL DEFAULT '' COMMENT '支付密码;micro_password加密',
  `freeze_time` int(11) NOT NULL DEFAULT '0' COMMENT '冻结时间',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `user_email` varchar(100) NOT NULL DEFAULT '' COMMENT '用户登录邮箱',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像',
  `signature` varchar(255) NOT NULL DEFAULT '' COMMENT '个性签名',
  `last_login_ip` varchar(15) NOT NULL DEFAULT '' COMMENT '最后登录ip',
  `user_activation_key` varchar(60) NOT NULL DEFAULT '' COMMENT '激活码',
  `mobile` VARCHAR(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '中国手机不带国家代码，国际手机号格式为：国家代码-手机号',
  `more` text COMMENT '扩展属性',
  PRIMARY KEY (`id`),
  KEY `user_login` (`user_login`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';


-- --------------------------------------------------------
--
-- 表的结构 `micro_bind_user`
--

CREATE TABLE IF NOT EXISTS `micro_bind_user` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` varchar(20) NOT NULL COMMENT '绑定账号',
  `start` int(11) NOT NULL DEFAULT '2' COMMENT '粉丝数',
  `region_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户地域',
  `class_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户分类',
  `bind_status` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '绑定状态;0:未绑定,1: 审核中,2:已绑定',
  `pass_time` int(11) NOT NULL DEFAULT '0' COMMENT '自动审核通过日期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='绑定账号表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_task_category`
--

CREATE TABLE IF NOT EXISTS `micro_task_category` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `parent_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '分类父id',
  `count` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '分类任务数',
  `valid_count` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '未完成任务数',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态,2:删除,1:发布,0:不发布',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
  `list_order` float NOT NULL DEFAULT '10000' COMMENT '排序',
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分类名称',
  `description` varchar(200) NOT NULL DEFAULT '' COMMENT '分类描述',
  `seo_title` varchar(100) NOT NULL DEFAULT '',
  `seo_keywords` varchar(200) NOT NULL DEFAULT '',
  `seo_description` varchar(200) NOT NULL DEFAULT '',
  `one_tpl` varchar(50) NOT NULL DEFAULT '' COMMENT '分类模板',
  `more` text COMMENT '扩展属性',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='任务分类表';


-- --------------------------------------------------------
--
-- 表的结构 `micro_task_list_201904`
--

CREATE TABLE IF NOT EXISTS `micro_task_list_201904` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '1' COMMENT '发布者id',
  `class_id` bigint(20) UNSIGNED NOT NULL DEFAULT '1'  COMMENT '分类id',
  `count` int(15) NOT NULL DEFAULT '1' COMMENT '任务数量',
  `consume_count` int(15) NOT NULL DEFAULT '0' COMMENT '已接单任务数量',
  `check_count` int(15) NOT NULL DEFAULT '0' COMMENT '提交审核任务数量',
  `finish_count` int(15) NOT NULL DEFAULT '0' COMMENT '完成任务数量',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '任务单价',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '发布日期',
  `end_time` int(11) NOT NULL DEFAULT '0' COMMENT '结束日期',
  `list_order` float NOT NULL DEFAULT '10000' COMMENT '排序',
  `region_rule` bigint(10) UNSIGNED NOT NULL DEFAULT '0'  COMMENT '地域id限制',
  `class_rule` tinyint(2) UNSIGNED NOT NULL DEFAULT '0' COMMENT '类别限制;0: 不限制，1: 同类别限制',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '发布状态;0:未发布,1: 审核中,2:已发布',
  `task_title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '任务标题',
  `task_describe` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '任务简述',
  `thumbnail` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '缩略图',
  `task_content` text COMMENT '任务内容',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='任务表';

-- --------------------------------------------------------
--
-- 表的结构 `micro_task_black_list`
--

CREATE TABLE IF NOT EXISTS `micro_task_black_list` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='任务黑名单';

-- --------------------------------------------------------
--
-- 表的结构 `micro_task_action_201904`
--

CREATE TABLE IF NOT EXISTS `micro_task_action_201904` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `task_id` bigint(20) UNSIGNED NOT NULL DEFAULT '1' COMMENT '任务id',
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '1' COMMENT '用户id',
  `region_id` int(10) UNSIGNED NOT NULL DEFAULT '1' COMMENT '用户地域',
  `class_id` int(10) UNSIGNED NOT NULL DEFAULT '1' COMMENT '用户分类',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '任务单价',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '接单日期',
  `end_time` int(11) NOT NULL DEFAULT '0' COMMENT '结束日期',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '状态;0:未完成,1: 自动审核中,2:已完成,3:复审,4: 无效',
  `check_count` tinyint(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '复审次数',
  `qr_url` varchar(100) NOT NULL DEFAULT 'noset' COMMENT '二维码路径',
  `comment_level` tinyint(3) UNSIGNED NOT NULL DEFAULT '4' COMMENT '评论等级',
  `comment_text` text COMMENT '评论内容',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '状态;0:未完成,1: 审核中,2:已完成',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='任务接单表';

-- --------------------------------------------------------
--
-- 表的结构 `micro_task_check_log_201904`
--

CREATE TABLE IF NOT EXISTS `micro_task_check_log_201904` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '1' COMMENT '用户id',
  `task_id` bigint(20) UNSIGNED NOT NULL DEFAULT '1' COMMENT '任务id',
  `reason_describe` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '原因描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='任务复审表';


-- --------------------------------------------------------

--
-- 表的结构 `micro_user_action_log_201904`
--

CREATE TABLE IF NOT EXISTS `micro_user_action_log_201904` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户id',
  `count` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '访问次数',
  `last_visit_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '最后访问时间',
  `parameter` varchar(100) NOT NULL DEFAULT '' COMMENT '参数',
  `action` varchar(50) NOT NULL DEFAULT '' COMMENT '操作名称;url',
  `ip` varchar(15) NOT NULL DEFAULT '' COMMENT '用户ip',
  PRIMARY KEY (`id`),
  KEY `user_parameter_action` (`user_id`,`parameter`,`action`),
  KEY `user_parameter_action_ip` (`user_id`,`parameter`,`action`,`ip`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='访问记录表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_user_balance_log_201904`
--

CREATE TABLE IF NOT EXISTS `micro_user_balance_log_201904` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户 id',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `change` decimal(20,2) NOT NULL DEFAULT '0.00' COMMENT '更改余额',
  `balance` decimal(20,2) NOT NULL DEFAULT '0.00' COMMENT '更改后余额',
  `fee` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '手续费',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户余额变更日志表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_user_favorite`
--

CREATE TABLE IF NOT EXISTS `micro_user_favorite` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户 id',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '收藏内容的标题',
  `thumbnail` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '缩略图',
  `url` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '收藏内容的地址，JSON格式',
  `description` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '收藏内容的描述',
  `create_time` int(10) UNSIGNED DEFAULT '0' COMMENT '收藏时间',
  PRIMARY KEY (`id`),
  KEY `uid` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户收藏表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_user_login_attempt`
--

CREATE TABLE IF NOT EXISTS `micro_user_login_attempt` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户 id',
  `login_attempts` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '尝试次数',
  `attempt_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '尝试登录时间',
  `locked_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '锁定时间',
  `ip` varchar(15) NOT NULL DEFAULT '' COMMENT '用户 ip',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户登录尝试表' ROW_FORMAT=COMPACT;

-- --------------------------------------------------------

--
-- 表的结构 `micro_user_score_log_201904`
--

CREATE TABLE IF NOT EXISTS `micro_user_score_log_201904` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户 id',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `action` varchar(50) NOT NULL DEFAULT '' COMMENT '用户操作名称',
  `score` int(11) NOT NULL DEFAULT '0' COMMENT '更改积分，可以为负',
  `coin` int(11) NOT NULL DEFAULT '0' COMMENT '更改金币，可以为负',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户操作奖励日志表';


-- --------------------------------------------------------

--
-- 表的结构 `micro_verification_code_201904`
--

CREATE TABLE IF NOT EXISTS `micro_verification_code_201904` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '表id',
  `count` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '当天已经发送成功的次数',
  `send_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '最后发送成功时间',
  `expire_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '验证码过期时间',
  `code` varchar(8) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '最后发送成功的验证码',
  `account` varchar(100) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '手机号或者邮箱',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='手机邮箱数字验证码表';


-- --------------------------------------------------------

--
-- 表的结构 `micro_user_like_201904`
--
CREATE TABLE IF NOT EXISTS `micro_user_like_201904` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户 id',
  `object_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '内容原来的主键id',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '内容的地址',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '内容的标题',
  `thumbnail` varchar(100) NOT NULL DEFAULT '' COMMENT '缩略图',
  `description` text COMMENT '内容的描述',
  PRIMARY KEY (`id`),
  KEY `uid` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户点赞表';


