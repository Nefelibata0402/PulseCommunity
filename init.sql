CREATE TABLE userinfo (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
    `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
    `username` longtext DEFAULT NULL COMMENT '用户名',
    `password` longtext DEFAULT NULL COMMENT '密码',
    `avatar` longtext DEFAULT NULL COMMENT '头像',
    `background_image` longtext DEFAULT NULL COMMENT '背景图',
    `signature` longtext DEFAULT NULL COMMENT '个人简介',
    PRIMARY KEY (`id`),
)ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = COMPACT;

CREATE TABLE article (
     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
     `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
     `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
     `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
     `user_id` bigint(20) unsigned DEFAULT NULL COMMENT '作者id',
     `content` longtext DEFAULT NULL COMMENT '文章内容',
     `category` longtext DEFAULT NULL COMMENT '文章分类',
     `title` longtext DEFAULT NULL COMMENT '文章标题',
     PRIMARY KEY (`id`),
) ENGINE=InnoDB CHARACTER SET=utf8 COLLATE=utf8_general_ci COMMENT='文章表' ROW_FORMAT=COMPACT;

CREATE TABLE `favorites` (
     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
     `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
     `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
     `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
     `article_id` bigint(20) DEFAULT NULL COMMENT '文章id',
     `user_id` bigint(20) DEFAULT NULL COMMENT '作者id',
     PRIMARY KEY (`id`),
) ENGINE=InnoDB CHARACTER SET=utf8 COLLATE=utf8_general_ci COMMENT='喜欢列表' ROW_FORMAT=COMPACT;

CREATE TABLE `comments` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
    `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
    `article_id` bigint(20) DEFAULT NULL COMMENT '文章id',
    `user_id` bigint(20) DEFAULT NULL COMMENT '作者id',
    `content` longtext DEFAULT NULL COMMENT '评论内容',
    PRIMARY KEY (`id`),
) ENGINE=InnoDB CHARACTER SET=utf8 COLLATE=utf8_general_ci COMMENT='评论表' ROW_FORMAT=COMPACT;


CREATE TABLE `like` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `biz_id` BIGINT,
    `biz` VARCHAR(128),
    `status` INT,
    `created_at` BIGINT,
    `updated_at` BIGINT,
    UNIQUE KEY `uid_biz_type_id` (`id`, `biz_id`, `biz`)
)ENGINE=InnoDB CHARACTER SET=utf8 COLLATE=utf8_general_ci COMMENT='喜欢表' ROW_FORMAT=COMPACT;

CREATE TABLE `collection`(
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `biz_id` BIGINT,
    `biz` VARCHAR(128),
    `cid` BIGINT,
    `created_at` BIGINT,
    `updated_at` BIGINT,
    UNIQUE KEY `uid_biz_type_id` (`id`, `biz_id`, `biz`)
)ENGINE=InnoDB CHARACTER SET=utf8 COLLATE=utf8_general_ci COMMENT='收藏表' ROW_FORMAT=COMPACT;

CREATE TABLE `interactive` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `biz_id` BIGINT,
    `biz` VARCHAR(128),
    `read_cnt` BIGINT,
    `like_cnt` BIGINT,
    `collect_cnt` BIGINT,
    `created_at` BIGINT,
    `updated_at` BIGINT,
    UNIQUE KEY `biz_type_id` (`biz_id`, `biz`)
)ENGINE=InnoDB CHARACTER SET=utf8 COLLATE=utf8_general_ci COMMENT='交互表' ROW_FORMAT=COMPACT;
