create table article
(
    uuid     int          not null,
    user_id  int          not null,
    name     varchar(50)  not null comment '文章名',
    info     text         null,
    tag      varchar(200) null comment '标签，分类',
    time     datetime     not null,
    like_sum int          null,
    primary key (`uuid`),
    constraint uuid
        unique (uuid)
);

create index article__name
    on article (name);

create index article__userid
    on article (user_id);