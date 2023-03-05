create xcpc_board;
create table id_platform
(
    id       int auto_increment
        primary key,
    uid      varchar(20) null,
    platform varchar(20) null,
    name_id  int         null,
    constraint id_platform_id_uindex
        unique (id)
)
    auto_increment = 4;

create index id_platform_score_id_fk
    on id_platform (name_id);

create table score
(
    id         int auto_increment
        primary key,
    name       varchar(10) null,
    cf_problem int         null,
    total      int         null,
    codeforces varchar(20) null,
    nowcoder   int         null,
    constraint score_id_uindex
        unique (id)
);

