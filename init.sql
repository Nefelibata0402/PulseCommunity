CREATE TABLE userinfo (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id` bigint(20) NOT NULL COMMENT '用户id',
    `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
    `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
    `username` longtext DEFAULT NULL COMMENT '用户名',
    `password` longtext DEFAULT NULL COMMENT '密码',
    `avatar` longtext DEFAULT NULL COMMENT '头像',
    `background_image` longtext DEFAULT NULL COMMENT '背景图',
    `signature` longtext DEFAULT NULL COMMENT '个人简介',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_user_id` (`user_id`), -- 创建唯一索引
    KEY `idx_users_deleted_at` (`deleted_at`)
)ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = COMPACT;

CREATE TABLE article (
     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
     `article_id` bigint(20) NOT NULL COMMENT '文章id',
     `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
     `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
     `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
     `user_id` bigint(20) unsigned DEFAULT NULL COMMENT '作者id',
     `content` longtext DEFAULT NULL COMMENT '文章内容',
     `category` longtext DEFAULT NULL COMMENT '文章分类',
     `title` longtext DEFAULT NULL COMMENT '文章标题',
     PRIMARY KEY (`id`),
     KEY `idx_videos_deleted_at` (`deleted_at`),
     UNIQUE KEY `unique_article_id` (`article_id`) -- 创建唯一索引
) ENGINE=InnoDB CHARACTER SET=utf8 COLLATE=utf8_general_ci COMMENT='文章表' ROW_FORMAT=COMPACT;

CREATE TABLE `favorites` (
     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
     `favorite_id` bigint(20) NOT NULL COMMENT '喜欢id',
     `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
     `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
     `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
     `article_id` bigint(20) DEFAULT NULL COMMENT '文章id',
     `user_id` bigint(20) DEFAULT NULL COMMENT '作者id',
     PRIMARY KEY (`id`),
     KEY `idx_favorites_deleted_at` (`deleted_at`),
     UNIQUE KEY `unique_favorite_id` (`favorite_id`) -- 创建唯一索引
) ENGINE=InnoDB CHARACTER SET=utf8 COLLATE=utf8_general_ci COMMENT='喜欢列表' ROW_FORMAT=COMPACT;

CREATE TABLE `comments` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `comment_id` bigint(20) NOT NULL COMMENT '评论id',
    `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
    `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
    `article_id` bigint(20) DEFAULT NULL COMMENT '文章id',
    `user_id` bigint(20) DEFAULT NULL COMMENT '作者id',
    `content` longtext DEFAULT NULL COMMENT '评论内容',
    PRIMARY KEY (`id`),
    KEY `idx_comments_deleted_at` (`deleted_at`),
    UNIQUE KEY `unique_comment_id` (`comment_id`) -- 创建唯一索引
) ENGINE=InnoDB CHARACTER SET=utf8 COLLATE=utf8_general_ci COMMENT='评论表' ROW_FORMAT=COMPACT;