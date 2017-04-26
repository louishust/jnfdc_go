DROP TABLE `jn_net_sign`;
CREATE TABLE `jn_net_sign` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `sign_type` varchar(256) NOT NULL COMMENT '网签类型',
  `sign_number` int unsigned NOT NULL COMMENT '网签数量',
  `sign_area` decimal(20,2) unsigned NOT NULL COMMENT '网签面积',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='济南网签表';
alter table 
