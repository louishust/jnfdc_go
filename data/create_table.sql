DROP TABLE `jn_net_sign`;
CREATE TABLE `jn_net_sign` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `sign_type` varchar(128) NOT NULL COMMENT '网签类型',
  `sign_number` int unsigned NOT NULL COMMENT '网签数量',
  `sign_area` decimal(20,2) unsigned NOT NULL COMMENT '网签面积',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='济南网签表';


CREATE TABLE dbversion(version int);

insert into dbversion values(1);

DROP TABLE `qd_net_sign_first_hand`;
CREATE TABLE `qd_net_sign_first_hand` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `sign_region` varchar(256) NOT NULL COMMENT '区县',
  `sign_type` varchar(256) NOT NULL COMMENT '网签类型',
  `sign_number` int unsigned NOT NULL COMMENT '网签数量',
  `sign_area` decimal(20,2) unsigned NOT NULL COMMENT '网签面积',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='青岛一手房网签';

DROP TABLE `qd_net_sign_second_hand`;
CREATE TABLE `qd_net_sign_second_hand` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `sign_region` varchar(256) NOT NULL COMMENT '区县',
  `sign_type` varchar(256) NOT NULL COMMENT '网签类型',
  `sign_number` int unsigned NOT NULL COMMENT '网签数量',
  `sign_area` decimal(20,2) unsigned NOT NULL COMMENT '网签面积',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='青岛二手房网签';



-- 2017.6.12 10:00 add by lou
insert into dbversion values(2);
CREATE TABLE `jn_net_sign_region` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `region_id` bigint(11) unsigned NOT NULL COMMENT '区域ID',
  `region_name` varchar(256) NOT NULL COMMENT '区域名称',
  `onsale_num` bigint(11) unsigned NOT NULL COMMENT '可售套数',
  `house_onsale_num` bigint(11) unsigned NOT NULL COMMENT '住宅可售套数',
  `sign_number` int unsigned NOT NULL COMMENT '网签数量',
  `sign_area` decimal(20,2) unsigned NOT NULL COMMENT '网签面积',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='济南网签区域表';



-- 2017.6.30 14:00 add by lou
insert into dbversion values(3);
CREATE TABLE `jn_net_sign_static` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `sign_type` varchar(128) NOT NULL COMMENT '网签类型',
  `sign_number` int unsigned NOT NULL COMMENT '网签数量',
  `sign_date` date NOT NULL COMMENT '日期',
  PRIMARY KEY (`id`),
  UNIQUE KEY(`sign_date`, sign_type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='济南网签统计表';

INSERT INTO jn_net_sign_static(sign_type, sign_number, sign_date)
SELECT sign_type, max(sign_number) as sign_number, date(created) as sign_date
FROM jn_net_sign group by sign_date, sign_type;



alter table jn_net_sign change column `sign_type` `sign_type` varchar(128) NOT NULL COMMENT '网签类型';
alter table jn_net_sign add index(created, sign_type);








