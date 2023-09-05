create table user
(
    id       int auto_increment comment '主键id',
    uuid     int          not null comment 'uuid',
    name     varchar(30)  null comment '用户名',
    password varchar(100) null comment '密码',
    sex      varchar(10)  null comment '性别',
    age      int          null comment '年龄',
    phone    int          null comment '手机号',
    mail     varchar(30)  null comment '邮箱',
    identity int          null comment '0为在校生,1为应届生,2为毕业生',
    need     text         null comment '用户感兴趣的职位',
    auth     int          null comment '1为管理员,0为用户',
    PRIMARY KEY (`id`),
    constraint uuid
        unique (uuid)
)
    comment '用户信息';

create index user_id_index
    on user (id);

create index user_uuid
    on user (uuid);

