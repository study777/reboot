CREATE TABLE `task` (
      `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
      `namespace` varchar(256) NOT NULL DEFAULT '' COMMENT 'task 附属Namespace',
      `resource` varchar(16) NOT NULL DEFAULT '' COMMENT 'task 附属资源类型',
      `task_type` varchar(32) NOT NULL DEFAULT '' COMMENT '任务类型: create, update, delete , etc.',
      `spec` text COMMENT '任务参数',
      `status` text COMMENT '任务状态',
      `is_canceled` tinyint(4) NOT NULL DEFAULT '0' COMMENT '任务是否被取消',
      `is_paused` tinyint(4) NOT NULL DEFAULT '0' COMMENT '任务是否在暂停状态',
      `is_skip_paused` tinyint(4) NOT NULL DEFAULT '0' COMMENT '任务是否在跳过暂停标志',
      `is_urgent_skipped` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否紧急跳过',
      `urgent_skip_comment` varchar(64) NOT NULL DEFAULT '' COMMENT '紧急跳过备注',
      `is_closed` tinyint(4) NOT NULL DEFAULT '0' COMMENT '任务是否已经关闭',
      `is_closed_manually` tinyint(4) NOT NULL DEFAULT '0' COMMENT '手动关掉任务',
      `op_user` varchar(32) NOT NULL DEFAULT '' COMMENT '创建任务的用户',
      `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:01' COMMENT '发起时间',
      `last_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近修改时间',
      PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='任务表';
