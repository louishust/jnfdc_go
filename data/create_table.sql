DROP TABLE `jn_net_sign`;
CREATE TABLE `jn_net_sign` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `sign_type` varchar(256) NOT NULL COMMENT '网签类型',
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


