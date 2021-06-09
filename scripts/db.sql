CREATE DATABASE IF NOT EXISTS go_vue_admin CHARSET utf8mb4;

USE go_vue_admin;

-- for mysql version 8.0-
-- CREATE USER IF NOT EXISTS go_vue_admin IDENTIFIED BY 'go_vue_admin';
-- GRANT ALL PRIVILEGES ON go_vue_admin.* TO 'go_vue_admin'@'%' IDENTIFIED BY 'go_vue_admin';

-- for mysql version 8.0+
CREATE USER IF NOT EXISTS 'go_vue_admin'@'%' IDENTIFIED WITH mysql_native_password BY 'go_vue_admin';
GRANT ALL PRIVILEGES ON go_vue_admin.* TO 'go_vue_admin'@'%';

FLUSH PRIVILEGES;

CREATE TABLE IF NOT EXISTS sys_user
(
    id                   int auto_increment primary key,
    user_name            varchar(64)  not null default '',
    nick_name            varchar(64)  not null default '',
    password             varchar(128) not null default '',
    salt                 varchar(128) not null default '',
    state                tinyint(2)   not null default 0,
    avatar               varchar(512) not null default '',
    email                varchar(128) not null default '',
    phone                varchar(64)  not null default '',
    remark               varchar(128) not null default '',
    modify_password_time bigint       not null default 0,
    last_login_time      bigint       not null default 0,
    last_login_ip        varchar(128) not null default '',
    updated              bigint       not null default 0,
    created              bigint       not null default 0,
    deleted              bigint       not null default 0,
    unique uq_sys_user_name (user_name)
);
