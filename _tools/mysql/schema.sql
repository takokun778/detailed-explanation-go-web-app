CREATE TABLE `user`
(
    `id`        BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザー識別子',
    `name`      VARCHAR(20) NOT NULL COMMENT 'ユーザー名',
    `password`  VARCHAR(80) NOT NULL COMMENT 'パスワードハッシュ',
    `role`      VARCHAR(80) NOT NULL COMMENT 'ロール',
    `created`   DATETIME NOT NULL COMMENT 'レコード作成日時',
    `modified`  DATETIME NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uix_name` (`name`) USING BTREE
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

CREATE TABLE `task`
(
    `id`        BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'タスク識別子',
    `title`     VARCHAR(128) NOT NULL COMMENT 'タスクタイトル',
    `status`    VARCHAR(80) NOT NULL COMMENT 'タスク状態',
    `created`   DATETIME NOT NULL COMMENT 'レコード作成日時',
    `modified`  DATETIME NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='タスク';
