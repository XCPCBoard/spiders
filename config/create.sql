CREATE
DATABASE IF NOT EXISTS xcpc_board_mysql;
use
xcpc_board_mysql;

drop table if exists id_platform;
CREATE TABLE `id_platform`
(
    `uid`      VARCHAR(40) COMMENT '账号id',
    `name_id`  bigint COMMENT '成员id',
    `platform` VARCHAR(40) COMMENT '平台'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='平台及账号';

drop table if exists score;
CREATE TABLE `score`
(
    `id`              bigint COMMENT '成员id',
    `cf_problem`      int COMMENT '过题数',
    `cf_rank`         int COMMENT 'cf rating',
    `nk_problem`      int COMMENT '牛客过题数',
    `nk_rank`         int COMMENT '牛客rating',
    `luogu_problem`   int COMMENT '洛谷过题数',
    `luogu_rank`      int COMMENT '洛谷rating',
    `atcoder_problem` int COMMENT 'atcoder过题数',
    `atcoder_rank`    int COMMENT 'atcoder rating',
    `vjudge_problem`  int COMMENT 'vjudge过题数'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='平台及账号';