
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for gen_table
-- ----------------------------
DROP TABLE IF EXISTS `gen_table`;
CREATE TABLE `gen_table`  (
                              `table_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
                              `table_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '表名称',
                              `parent_menu_id` bigint(20) NULL DEFAULT NULL COMMENT '父菜单ID',
                              `table_comment` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '表描述',
                              `sub_table_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '关联子表的表名',
                              `sub_table_fk_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '子表关联的外键名',
                              `struct_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '结构体名称',
                              `tpl_category` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'crud' COMMENT '使用的模板（crud单表操作 tree树表操作 sub主子表操作）',
                              `package_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '生成包路径',
                              `module_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '生成模块名',
                              `business_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '生成业务名',
                              `function_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '生成功能名',
                              `function_author` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '生成功能作者',
                              `options` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '其它生成选项',
                              `create_by` bigint(20) NULL DEFAULT NULL COMMENT '创建者',
                              `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                              `update_by` bigint(20) NULL DEFAULT NULL COMMENT '更新者',
                              `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                              `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
                              PRIMARY KEY (`table_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 221635920305065985 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '代码生成业务表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gen_table
-- ----------------------------

-- ----------------------------
-- Table structure for gen_table_column
-- ----------------------------
DROP TABLE IF EXISTS `gen_table_column`;
CREATE TABLE `gen_table_column`  (
                                     `column_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
                                     `table_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '归属表编号',
                                     `column_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '列名称',
                                     `column_comment` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '列描述',
                                     `column_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '列类型',
                                     `go_type` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'go类型',
                                     `go_field` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'go字段名',
                                     `is_pk` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '是否主键（1是）',
                                     `is_required` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '是否必填（1是）',
                                     `is_insert` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '是否为插入字段（1是）',
                                     `is_edit` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '是否编辑字段（1是）',
                                     `is_list` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '是否列表字段（1是）',
                                     `is_query` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '是否查询字段（1是）',
                                     `query_type` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'EQ' COMMENT '查询方式（等于、不等于、大于、小于、范围）',
                                     `html_type` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）',
                                     `dict_type` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '字典类型',
                                     `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
                                     `create_by` bigint(20) NULL DEFAULT NULL COMMENT '创建者',
                                     `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                                     `update_by` bigint(20) NULL DEFAULT NULL COMMENT '更新者',
                                     `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                                     `html_field` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
                                     PRIMARY KEY (`column_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 221635920317648906 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '代码生成业务表字段' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gen_table_column
-- ----------------------------

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
                               `config_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '参数主键',
                               `config_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '参数名称',
                               `config_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '参数键名',
                               `config_value` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '参数键值',
                               `config_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'N' COMMENT '系统内置（Y是 N否）',
                               `create_by` bigint(20) NULL DEFAULT NULL COMMENT '创建者',
                               `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                               `update_by` bigint(20) NULL DEFAULT NULL COMMENT '更新者',
                               `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                               `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
                               PRIMARY KEY (`config_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '参数配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (1, '主框架页-默认皮肤样式名称', 'sys.index.skinName', 'skin-blue', 'Y', 1, '2024-02-08 04:10:56', 1, NULL, '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow');
INSERT INTO `sys_config` VALUES (2, '用户管理-账号初始密码', 'sys.user.initPassword', '123456', 'Y', 1, '2024-02-08 04:10:56', 1, NULL, '初始化密码 123456');
INSERT INTO `sys_config` VALUES (3, '主框架页-侧边栏主题', 'sys.index.sideTheme', 'theme-dark', 'Y', 1, '2024-02-08 04:10:56', 1, NULL, '深色主题theme-dark，浅色主题theme-light');
INSERT INTO `sys_config` VALUES (4, '账号自助-验证码开关', 'sys.account.captchaEnabled', 'false', 'Y', 1, '2024-02-08 04:10:56', 1, '2024-03-09 13:17:45', '是否开启验证码功能（true开启，false关闭）');
INSERT INTO `sys_config` VALUES (5, '账号自助-是否开启用户注册功能', 'sys.account.registerUser', 'true', 'Y', 1, '2024-02-08 04:10:56', 1, '2024-03-09 13:24:50', '是否开启注册用户功能（true开启，false关闭）');

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
                             `dept_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '部门id',
                             `parent_id` bigint(20) NULL DEFAULT 0 COMMENT '父部门id',
                             `ancestors` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '祖级列表',
                             `dept_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '部门名称',
                             `order_num` int(11) NULL DEFAULT 0 COMMENT '显示顺序',
                             `leader` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '负责人',
                             `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '联系电话',
                             `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '邮箱',
                             `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
                             `del_flag` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
                             `create_by` bigint(20) NULL DEFAULT NULL COMMENT '创建者',
                             `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                             `update_by` bigint(20) NULL DEFAULT NULL COMMENT '更新者',
                             `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                             PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 110 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '部门表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (100, 0, '0', '白泽科技', 0, '白泽', '15888888888', 'bz@qq.com', '0', '0', 1, '2024-02-08 04:10:55', 1, NULL);
INSERT INTO `sys_dept` VALUES (101, 100, '0,100', '深圳总公司', 1, '白泽', '15888888888', 'bz@qq.com', '0', '0', 1, '2024-02-08 04:10:55', 1, NULL);
INSERT INTO `sys_dept` VALUES (102, 100, '0,100', '长沙分公司', 2, '白泽', '15888888888', 'bz@qq.com', '0', '0', 1, '2024-02-08 04:10:55', 1, NULL);
INSERT INTO `sys_dept` VALUES (103, 101, '0,100,101', '研发部门', 1, '白泽', '15888888888', 'bz@qq.com', '0', '0', 1, '2024-02-08 04:10:55', 1, NULL);
INSERT INTO `sys_dept` VALUES (104, 101, '0,100,101', '市场部门', 2, '白泽', '15888888888', 'bz@qq.com', '0', '0', 1, '2024-02-08 04:10:55', 1, NULL);
INSERT INTO `sys_dept` VALUES (105, 101, '0,100,101', '测试部门', 3, '白泽', '15888888888', 'bz@qq.com', '0', '0', 1, '2024-02-08 04:10:55', 1, NULL);
INSERT INTO `sys_dept` VALUES (106, 101, '0,100,101', '财务部门', 4, '白泽', '15888888888', 'bz@qq.com', '0', '0', 1, '2024-02-08 04:10:55', 1, NULL);
INSERT INTO `sys_dept` VALUES (107, 101, '0,100,101', '运维部门', 5, '白泽', '15888888888', 'bz@qq.com', '0', '0', 1, '2024-02-08 04:10:55', 1, NULL);
INSERT INTO `sys_dept` VALUES (108, 102, '0,100,102', '市场部门', 1, '白泽', '15888888888', 'bz@qq.com', '0', '0', 1, '2024-02-08 04:10:55', 1, NULL);
INSERT INTO `sys_dept` VALUES (109, 102, '0,100,102', '财务部门', 2, '白泽', '15888888888', 'bz@qq.com', '0', '0', 1, '2024-02-08 04:10:55', 1, NULL);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
                                  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典编码',
                                  `dict_sort` int(11) NULL DEFAULT 0 COMMENT '字典排序',
                                  `dict_label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '字典标签',
                                  `dict_value` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '字典键值',
                                  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '字典类型',
                                  `css_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '样式属性（其他样式扩展）',
                                  `list_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '表格回显样式',
                                  `is_default` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'N' COMMENT '是否默认（Y是 N否）',
                                  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
                                  `create_by` bigint(20) NULL DEFAULT NULL COMMENT '创建者',
                                  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                                  `update_by` bigint(20) NULL DEFAULT NULL COMMENT '更新者',
                                  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                                  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '备注',
                                  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 200266456896638976 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '字典数据表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES (1, 1, '男', '0', 'sys_user_sex', '', '', 'Y', '0', 1, '2024-02-08 04:10:56', 1, NULL, '性别男');
INSERT INTO `sys_dict_data` VALUES (2, 2, '女', '1', 'sys_user_sex', '', '', 'N', '0', 1, '2024-02-08 04:10:56', 1, NULL, '性别女');
INSERT INTO `sys_dict_data` VALUES (3, 3, '未知', '2', 'sys_user_sex', '', '', 'N', '0', 1, '2024-02-08 04:10:56', 1, NULL, '性别未知');
INSERT INTO `sys_dict_data` VALUES (6, 1, '正常', '0', 'sys_normal_disable', '', 'primary', 'Y', '0', 1, '2024-02-08 04:10:56', 1, NULL, '正常状态');
INSERT INTO `sys_dict_data` VALUES (7, 2, '停用', '1', 'sys_normal_disable', '', 'danger', 'N', '0', 1, '2024-02-08 04:10:56', 1, NULL, '停用状态');
INSERT INTO `sys_dict_data` VALUES (8, 1, '正常', '0', 'sys_job_status', '', 'primary', 'Y', '0', 1, '2024-02-08 04:10:56', 1, NULL, '正常状态');
INSERT INTO `sys_dict_data` VALUES (9, 2, '暂停', '1', 'sys_job_status', '', 'danger', 'N', '0', 1, '2024-02-08 04:10:56', 1, NULL, '停用状态');
INSERT INTO `sys_dict_data` VALUES (10, 1, '默认', 'DEFAULT', 'sys_job_group', '', '', 'Y', '0', 1, '2024-02-08 04:10:56', 1, NULL, '默认分组');
INSERT INTO `sys_dict_data` VALUES (11, 2, '系统', 'SYSTEM', 'sys_job_group', '', '', 'N', '0', 1, '2024-02-08 04:10:56', 1, NULL, '系统分组');
INSERT INTO `sys_dict_data` VALUES (12, 1, '是', 'Y', 'sys_yes_no', '', 'primary', 'Y', '0', 1, '2024-02-08 04:10:56', 1, NULL, '系统默认是');
INSERT INTO `sys_dict_data` VALUES (13, 2, '否', 'N', 'sys_yes_no', '', 'danger', 'N', '0', 1, '2024-02-08 04:10:56', 1, NULL, '系统默认否');
INSERT INTO `sys_dict_data` VALUES (14, 1, '通知', '1', 'sys_notice_type', '', 'warning', 'Y', '0', 1, '2024-02-08 04:10:56', 1, NULL, '通知');
INSERT INTO `sys_dict_data` VALUES (15, 2, '公告', '2', 'sys_notice_type', '', 'success', 'N', '0', 1, '2024-02-08 04:10:56', 1, NULL, '公告');
INSERT INTO `sys_dict_data` VALUES (18, 0, '其他', '0', 'sys_oper_type', '', 'info', 'N', '0', 1, '2024-02-08 04:10:56', 1, NULL, '其他操作');
INSERT INTO `sys_dict_data` VALUES (19, 1, '新增', '1', 'sys_oper_type', '', 'info', 'N', '0', 1, '2024-02-08 04:10:56', 1, NULL, '新增操作');
INSERT INTO `sys_dict_data` VALUES (20, 2, '修改', '2', 'sys_oper_type', '', 'info', 'N', '0', 1, '2024-02-08 04:10:56', 1, NULL, '修改操作');
INSERT INTO `sys_dict_data` VALUES (21, 3, '删除', '3', 'sys_oper_type', '', 'danger', 'N', '0', 1, '2024-02-08 04:10:56', 1, NULL, '删除操作');
INSERT INTO `sys_dict_data` VALUES (22, 4, '强退', '4', 'sys_oper_type', '', 'danger', 'N', '0', 1, '2024-02-08 04:10:56', 1, '2024-04-12 14:33:31', '强退操作');
INSERT INTO `sys_dict_data` VALUES (23, 5, '清空数据', '5', 'sys_oper_type', '', 'danger', 'N', '0', 1, '2024-02-08 04:10:56', 1, NULL, '清空操作');
INSERT INTO `sys_dict_data` VALUES (28, 1, '成功', '0', 'sys_common_status', '', 'primary', 'N', '0', 1, '2024-02-08 04:10:56', 1, NULL, '正常状态');
INSERT INTO `sys_dict_data` VALUES (29, 2, '失败', '1', 'sys_common_status', '', 'danger', 'N', '0', 1, '2024-02-08 04:10:56', 1, NULL, '停用状态');

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type`  (
                                  `dict_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典主键',
                                  `dict_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '字典名称',
                                  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '字典类型',
                                  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
                                  `create_by` bigint(20) NULL DEFAULT NULL COMMENT '创建者',
                                  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                                  `update_by` bigint(20) NULL DEFAULT NULL COMMENT '更新者',
                                  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                                  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
                                  PRIMARY KEY (`dict_id`) USING BTREE,
                                  UNIQUE INDEX `dict_type`(`dict_type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '字典类型表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES (1, '用户性别', 'sys_user_sex', '0', 1, '2024-02-08 04:10:55', 1, NULL, '用户性别列表');
INSERT INTO `sys_dict_type` VALUES (3, '系统开关', 'sys_normal_disable', '0', 1, '2024-02-08 04:10:55', 1, NULL, '系统开关列表');
INSERT INTO `sys_dict_type` VALUES (4, '任务状态', 'sys_job_status', '0', 1, '2024-02-08 04:10:55', 1, NULL, '任务状态列表');
INSERT INTO `sys_dict_type` VALUES (5, '任务分组', 'sys_job_group', '0', 1, '2024-02-08 04:10:55', 1, NULL, '任务分组列表');
INSERT INTO `sys_dict_type` VALUES (6, '系统是否', 'sys_yes_no', '0', 1, '2024-02-08 04:10:55', 1, NULL, '系统是否列表');
INSERT INTO `sys_dict_type` VALUES (7, '通知类型', 'sys_notice_type', '0', 1, '2024-02-08 04:10:55', 1, NULL, '通知类型列表');
INSERT INTO `sys_dict_type` VALUES (9, '操作类型', 'sys_oper_type', '0', 1, '2024-02-08 04:10:55', 1, NULL, '操作类型列表');
INSERT INTO `sys_dict_type` VALUES (10, '系统状态', 'sys_common_status', '0', 1, '2024-02-08 04:10:55', 1, NULL, '登录状态列表');

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job`  (
                            `job_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
                            `job_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '任务名称',
                            `job_params` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT NULL COMMENT '参数',
                            `invoke_target` varchar(500) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '调用目标字符串',
                            `cron_expression` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT '' COMMENT 'cron执行表达式',
                            `status` char(1) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT '0' COMMENT '状态（0正常 1暂停）',
                            `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT '' COMMENT '创建者',
                            `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                            `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT '' COMMENT '更新者',
                            `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                            PRIMARY KEY (`job_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 100 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci COMMENT = '定时任务调度表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_job
-- ----------------------------

-- ----------------------------
-- Table structure for sys_job_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_job_log`;
CREATE TABLE `sys_job_log`  (
                                `job_log_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务日志ID',
                                `job_id` bigint(20) NOT NULL COMMENT '任务ID',
                                `job_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '任务名称',
                                `job_group` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '任务组名',
                                `invoke_target` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '调用目标字符串',
                                `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '执行状态（0正常 1失败）',
                                `exception_info` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '异常信息',
                                `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                                `cost_time` bigint(20) NULL DEFAULT 0 COMMENT '耗时（毫秒）',
                                PRIMARY KEY (`job_log_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '定时任务调度日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_job_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_logininfor
-- ----------------------------
DROP TABLE IF EXISTS `sys_logininfor`;
CREATE TABLE `sys_logininfor`  (
                                   `info_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '访问ID',
                                   `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '用户账号',
                                   `ipaddr` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '登录IP地址',
                                   `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '登录地点',
                                   `browser` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '浏览器类型',
                                   `os` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '操作系统',
                                   `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0' COMMENT '登录状态（0成功 1失败）',
                                   `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '提示消息',
                                   `login_time` datetime(0) NULL DEFAULT NULL COMMENT '访问时间',
                                   PRIMARY KEY (`info_id`) USING BTREE,
                                   INDEX `idx_sys_logininfor_s`(`status`) USING BTREE,
                                   INDEX `idx_sys_logininfor_lt`(`login_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 352453849534959617 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '系统访问记录' ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for sys_notice
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice`;
CREATE TABLE `sys_notice`  (
                               `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '公告ID',
                               `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '公告标题',
                               `type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '公告类型（1通知 2公告）',
                               `txt` longblob NULL COMMENT '公告内容',
                               `dept_id` bigint(20) NULL DEFAULT NULL COMMENT '发件人所在部门',
                               `dept_ids` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '发送部门',
                               `create_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '创建者名称',
                               `create_by` bigint(20) NULL DEFAULT NULL COMMENT '创建者',
                               `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                               PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '通知公告表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_notice
-- ----------------------------

-- ----------------------------
-- Table structure for sys_notice_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice_user`;
CREATE TABLE `sys_notice_user`  (
                                    `user_id` bigint(20) NOT NULL,
                                    `notice_id` bigint(20) NOT NULL,
                                    `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                    PRIMARY KEY (`user_id`, `notice_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '消息通知关联用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_notice_user
-- ----------------------------

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log`  (
                                 `oper_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '日志主键',
                                 `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '模块标题',
                                 `business_type` int(11) NULL DEFAULT 0 COMMENT '业务类型（0其它 1新增 2修改 3删除）',
                                 `method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '方法名称',
                                 `request_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '请求方式',
                                 `oper_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '操作人员',
                                 `oper_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '请求URL',
                                 `oper_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '主机地址',
                                 `oper_param` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '请求参数',
                                 `json_result` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '返回参数',
                                 `status` int(11) NULL DEFAULT 0 COMMENT '操作状态（0正常 1异常）',
                                 `user_id` bigint(20) NULL DEFAULT NULL COMMENT '用户ID',
                                 `oper_time` datetime(0) NULL DEFAULT NULL COMMENT '操作时间',
                                 `cost_time` bigint(20) NULL DEFAULT 0 COMMENT '消耗时间',
                                 PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '操作日志记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_permission`;
CREATE TABLE `sys_permission`  (
                                   `permission_id` bigint(20) NOT NULL COMMENT '主键',
                                   `permission_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '权限名称',
                                   `parent_id` bigint(20) NULL DEFAULT NULL COMMENT '父ID',
                                   `permission` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '权限标识符',
                                   `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
                                   `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0' COMMENT '状态',
                                   `create_by` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
                                   `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                                   `update_by` bigint(20) NULL DEFAULT NULL COMMENT '更新人',
                                   `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                                   PRIMARY KEY (`permission_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_permission
-- ----------------------------
INSERT INTO `sys_permission` VALUES (1, '系统管理', 0, 'system', 1, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (2, '系统监控', 0, 'monitor', 2, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (3, '系统工具', 0, 'tool', 3, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (100, '用户管理', 1, 'system:user', 1, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (101, '角色管理', 1, 'system:role', 2, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (102, '权限管理', 1, 'system:permission', 3, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (103, '部门管理', 1, 'system:dept', 4, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (104, '岗位管理', 1, 'system:post', 5, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (105, '字典管理', 1, 'system:dict', 6, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (106, '参数设置', 1, 'system:config', 7, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (107, '通知公告', 1, 'system:notice', 8, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (108, '日志管理', 1, 'system:monitor', 9, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (109, '在线用户', 2, 'monitor:online', 1, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (110, '定时任务', 2, 'monitor:job', 2, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (112, '服务监控', 2, 'monitor:server', 3, '0', 1, '2025-02-28 13:38:05', 1, '2025-03-26 22:37:27');
INSERT INTO `sys_permission` VALUES (115, '表单构建', 3, 'tool:build', 1, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (116, '代码生成', 3, 'tool:gen', 2, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (117, '系统接口', 3, 'tool:swagger', 3, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (500, '操作日志', 108, 'system:monitor:operlog', 1, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (501, '登录日志', 108, 'system:monitor:logininfor', 2, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1000, '用户查询', 100, 'system:user:query', 1, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1001, '用户新增', 100, 'system:user:add', 2, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1002, '用户修改', 100, 'system:user:edit', 3, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1003, '用户删除', 100, 'system:user:remove', 4, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1004, '用户导出', 100, 'system:user:export', 5, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1005, '用户导入', 100, 'system:user:import', 6, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1006, '重置密码', 100, 'system:user:resetPwd', 7, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1007, '角色查询', 101, 'system:role:query', 1, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1008, '角色新增', 101, 'system:role:add', 2, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1009, '角色修改', 101, 'system:role:edit', 3, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1010, '角色删除', 101, 'system:role:remove', 4, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1011, '角色导出', 101, 'system:role:export', 5, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1012, '菜单查询', 102, 'system:permission:query', 1, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1013, '菜单新增', 102, 'system:permission:add', 2, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1014, '菜单修改', 102, 'system:permission:edit', 3, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1015, '菜单删除', 102, 'system:permission:remove', 4, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1016, '部门查询', 103, 'system:dept:query', 1, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1017, '部门新增', 103, 'system:dept:add', 2, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1018, '部门修改', 103, 'system:dept:edit', 3, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1019, '部门删除', 103, 'system:dept:remove', 4, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1020, '岗位查询', 104, 'system:post:query', 1, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1021, '岗位新增', 104, 'system:post:add', 2, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1022, '岗位修改', 104, 'system:post:edit', 3, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1023, '岗位删除', 104, 'system:post:remove', 4, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1024, '岗位导出', 104, 'system:post:export', 5, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1025, '字典查询', 105, 'system:dict:query', 1, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1026, '字典新增', 105, 'system:dict:add', 2, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1027, '字典修改', 105, 'system:dict:edit', 3, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1028, '字典删除', 105, 'system:dict:remove', 4, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1029, '字典导出', 105, 'system:dict:export', 5, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1030, '参数查询', 106, 'system:config:query', 1, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1031, '参数新增', 106, 'system:config:add', 2, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1032, '参数修改', 106, 'system:config:edit', 3, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1033, '参数删除', 106, 'system:config:remove', 4, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1034, '参数导出', 106, 'system:config:export', 5, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1035, '公告查询', 107, 'system:notice:query', 1, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1036, '公告新增', 107, 'system:notice:add', 2, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1037, '公告修改', 107, 'system:notice:edit', 3, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1038, '公告删除', 107, 'system:notice:remove', 4, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1039, '操作查询', 500, 'monitor:operlog:query', 1, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1040, '操作删除', 500, 'monitor:operlog:remove', 2, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1041, '日志导出', 500, 'monitor:operlog:export', 3, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1042, '登录查询', 501, 'monitor:logininfor:query', 1, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1043, '登录删除', 501, 'monitor:logininfor:remove', 2, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1044, '日志导出', 501, 'monitor:logininfor:export', 3, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1046, '在线查询', 109, 'monitor:online:query', 1, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1047, '批量强退', 109, 'monitor:online:batchLogout', 2, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1048, '单条强退', 109, 'monitor:online:forceLogout', 3, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1049, '任务查询', 110, 'monitor:job:query', 1, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1050, '任务新增', 110, 'monitor:job:add', 2, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1051, '任务修改', 110, 'monitor:job:edit', 3, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1052, '任务删除', 110, 'monitor:job:remove', 4, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1053, '状态修改', 110, 'monitor:job:changeStatus', 5, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1054, '任务导出', 110, 'monitor:job:export', 6, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1055, '生成查询', 116, 'tool:gen:query', 1, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1056, '生成修改', 116, 'tool:gen:edit', 2, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1057, '生成删除', 116, 'tool:gen:remove', 3, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1058, '导入代码', 116, 'tool:gen:import', 4, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1059, '预览代码', 116, 'tool:gen:preview', 5, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES (1060, '生成代码', 116, 'tool:gen:code', 6, '0', 1, '2025-02-28 13:38:05', 1, '2025-02-28 13:38:05');

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
                             `post_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
                             `post_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '岗位编码',
                             `post_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '岗位名称',
                             `post_sort` int(11) NOT NULL COMMENT '显示顺序',
                             `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '状态（0正常 1停用）',
                             `create_by` bigint(20) NULL DEFAULT NULL COMMENT '创建者',
                             `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                             `update_by` bigint(20) NULL DEFAULT NULL COMMENT '更新者',
                             `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                             `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
                             PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '岗位信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES (1, 'ceo', '董事长', 1, '0', 1, '2024-02-08 04:10:55', 1, NULL, '');
INSERT INTO `sys_post` VALUES (2, 'se', '项目经理', 2, '0', 1, '2024-02-08 04:10:55', 1, NULL, '');
INSERT INTO `sys_post` VALUES (3, 'hr', '人力资源', 3, '0', 1, '2024-02-08 04:10:55', 1, NULL, '');
INSERT INTO `sys_post` VALUES (4, 'user', '普通员工', 4, '0', 1, '2024-02-08 04:10:55', 1, '2024-02-25 10:40:41', '');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
                             `role_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
                             `role_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色名称',
                             `role_sort` int(11) NOT NULL COMMENT '显示顺序',
                             `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色状态（0正常 1停用）',
                             `del_flag` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
                             `create_by` bigint(20) NULL DEFAULT NULL COMMENT '创建者',
                             `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                             `update_by` bigint(20) NULL DEFAULT NULL COMMENT '更新者',
                             `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                             `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
                             PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 100 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, 'admin', 1, '0', '0', 1, '2024-02-08 04:10:55', 1, NULL, '超级管理员');
INSERT INTO `sys_role` VALUES (2, 'common', 2, '0', '0', 1, '2024-02-08 04:10:55', 1, '2024-03-04 14:16:58', '普通角色');

-- ----------------------------
-- Table structure for sys_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_permission`;
CREATE TABLE `sys_role_permission`  (
                                        `role_id` bigint(20) NOT NULL COMMENT '角色ID',
                                        `permission_id` bigint(20) NOT NULL COMMENT '权限ID',
                                        PRIMARY KEY (`role_id`, `permission_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色和权限关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_permission
-- ----------------------------
INSERT INTO `sys_role_permission` VALUES (2, 1);
INSERT INTO `sys_role_permission` VALUES (2, 2);
INSERT INTO `sys_role_permission` VALUES (2, 3);
INSERT INTO `sys_role_permission` VALUES (2, 4);
INSERT INTO `sys_role_permission` VALUES (2, 100);
INSERT INTO `sys_role_permission` VALUES (2, 101);
INSERT INTO `sys_role_permission` VALUES (2, 102);
INSERT INTO `sys_role_permission` VALUES (2, 103);
INSERT INTO `sys_role_permission` VALUES (2, 104);
INSERT INTO `sys_role_permission` VALUES (2, 105);
INSERT INTO `sys_role_permission` VALUES (2, 106);
INSERT INTO `sys_role_permission` VALUES (2, 107);
INSERT INTO `sys_role_permission` VALUES (2, 108);
INSERT INTO `sys_role_permission` VALUES (2, 109);
INSERT INTO `sys_role_permission` VALUES (2, 110);
INSERT INTO `sys_role_permission` VALUES (2, 111);
INSERT INTO `sys_role_permission` VALUES (2, 112);
INSERT INTO `sys_role_permission` VALUES (2, 113);
INSERT INTO `sys_role_permission` VALUES (2, 114);
INSERT INTO `sys_role_permission` VALUES (2, 115);
INSERT INTO `sys_role_permission` VALUES (2, 116);
INSERT INTO `sys_role_permission` VALUES (2, 500);
INSERT INTO `sys_role_permission` VALUES (2, 501);
INSERT INTO `sys_role_permission` VALUES (2, 1000);
INSERT INTO `sys_role_permission` VALUES (2, 1001);
INSERT INTO `sys_role_permission` VALUES (2, 1002);
INSERT INTO `sys_role_permission` VALUES (2, 1003);
INSERT INTO `sys_role_permission` VALUES (2, 1004);
INSERT INTO `sys_role_permission` VALUES (2, 1005);
INSERT INTO `sys_role_permission` VALUES (2, 1006);
INSERT INTO `sys_role_permission` VALUES (2, 1007);
INSERT INTO `sys_role_permission` VALUES (2, 1008);
INSERT INTO `sys_role_permission` VALUES (2, 1009);
INSERT INTO `sys_role_permission` VALUES (2, 1010);
INSERT INTO `sys_role_permission` VALUES (2, 1011);
INSERT INTO `sys_role_permission` VALUES (2, 1012);
INSERT INTO `sys_role_permission` VALUES (2, 1013);
INSERT INTO `sys_role_permission` VALUES (2, 1014);
INSERT INTO `sys_role_permission` VALUES (2, 1015);
INSERT INTO `sys_role_permission` VALUES (2, 1016);
INSERT INTO `sys_role_permission` VALUES (2, 1017);
INSERT INTO `sys_role_permission` VALUES (2, 1018);
INSERT INTO `sys_role_permission` VALUES (2, 1019);
INSERT INTO `sys_role_permission` VALUES (2, 1020);
INSERT INTO `sys_role_permission` VALUES (2, 1021);
INSERT INTO `sys_role_permission` VALUES (2, 1022);
INSERT INTO `sys_role_permission` VALUES (2, 1023);
INSERT INTO `sys_role_permission` VALUES (2, 1024);
INSERT INTO `sys_role_permission` VALUES (2, 1025);
INSERT INTO `sys_role_permission` VALUES (2, 1026);
INSERT INTO `sys_role_permission` VALUES (2, 1027);
INSERT INTO `sys_role_permission` VALUES (2, 1028);
INSERT INTO `sys_role_permission` VALUES (2, 1029);
INSERT INTO `sys_role_permission` VALUES (2, 1030);
INSERT INTO `sys_role_permission` VALUES (2, 1031);
INSERT INTO `sys_role_permission` VALUES (2, 1032);
INSERT INTO `sys_role_permission` VALUES (2, 1033);
INSERT INTO `sys_role_permission` VALUES (2, 1034);
INSERT INTO `sys_role_permission` VALUES (2, 1035);
INSERT INTO `sys_role_permission` VALUES (2, 1036);
INSERT INTO `sys_role_permission` VALUES (2, 1037);
INSERT INTO `sys_role_permission` VALUES (2, 1038);
INSERT INTO `sys_role_permission` VALUES (2, 1039);
INSERT INTO `sys_role_permission` VALUES (2, 1040);
INSERT INTO `sys_role_permission` VALUES (2, 1041);
INSERT INTO `sys_role_permission` VALUES (2, 1042);
INSERT INTO `sys_role_permission` VALUES (2, 1043);
INSERT INTO `sys_role_permission` VALUES (2, 1044);
INSERT INTO `sys_role_permission` VALUES (2, 1045);
INSERT INTO `sys_role_permission` VALUES (2, 1046);
INSERT INTO `sys_role_permission` VALUES (2, 1047);
INSERT INTO `sys_role_permission` VALUES (2, 1048);
INSERT INTO `sys_role_permission` VALUES (2, 1049);
INSERT INTO `sys_role_permission` VALUES (2, 1050);
INSERT INTO `sys_role_permission` VALUES (2, 1051);
INSERT INTO `sys_role_permission` VALUES (2, 1052);
INSERT INTO `sys_role_permission` VALUES (2, 1053);
INSERT INTO `sys_role_permission` VALUES (2, 1054);
INSERT INTO `sys_role_permission` VALUES (2, 1055);
INSERT INTO `sys_role_permission` VALUES (2, 1056);
INSERT INTO `sys_role_permission` VALUES (2, 1057);
INSERT INTO `sys_role_permission` VALUES (2, 1058);
INSERT INTO `sys_role_permission` VALUES (2, 1059);
INSERT INTO `sys_role_permission` VALUES (2, 1060);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
                             `user_id` bigint(20) NOT NULL COMMENT '用户ID',
                             `dept_id` bigint(20) NULL DEFAULT NULL COMMENT '部门ID',
                             `user_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户账号',
                             `nick_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户昵称',
                             `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '用户邮箱',
                             `phonenumber` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '手机号码',
                             `sex` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0' COMMENT '用户性别（0男 1女 2未知）',
                             `avatar` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '头像地址',
                             `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '密码',
                             `data_scope` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '5' COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限,无任何）权限',
                             `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0' COMMENT '帐号状态（0正常 1停用）',
                             `del_flag` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
                             `create_by` bigint(20) NULL DEFAULT NULL COMMENT '创建者',
                             `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                             `update_by` bigint(20) NULL DEFAULT NULL COMMENT '更新者',
                             `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                             `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '备注',
                             PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 105, 'admin', '白泽', 'bz@163.com', '15888888887', '0', '', '$2a$14$kOhr1gqDAGYara15BrSuLuI7uzOC7gfAyzJGiLsypO32mEb6O02LO', '1', '0', '0', 1, '2024-02-08 04:10:55', 1, '2024-03-15 09:31:36', '管理员');
INSERT INTO `sys_user` VALUES (2, 105, 'bz', '白泽', 'bz@qq.com', '15666666666', '1', '', '$2a$10$7JB720yubVSZvUI0rEqK/.VqGOZTH.ulu33dHOiBE8ByOhJIrdAu2', '2', '0', '0', 1, '2024-02-08 04:10:55', 1, '2025-04-06 22:07:27', '测试员');

-- ----------------------------
-- Table structure for sys_user_dept_scope
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_dept_scope`;
CREATE TABLE `sys_user_dept_scope`  (
                                        `user_id` bigint(20) NOT NULL COMMENT '用户ID',
                                        `dept_id` bigint(20) NOT NULL COMMENT '部门ID',
                                        PRIMARY KEY (`user_id`, `dept_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户和部门关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_dept_scope
-- ----------------------------
INSERT INTO `sys_user_dept_scope` VALUES (2, 101);
INSERT INTO `sys_user_dept_scope` VALUES (2, 102);
INSERT INTO `sys_user_dept_scope` VALUES (2, 103);
INSERT INTO `sys_user_dept_scope` VALUES (2, 104);
INSERT INTO `sys_user_dept_scope` VALUES (2, 108);
INSERT INTO `sys_user_dept_scope` VALUES (2, 109);

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post`  (
                                  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
                                  `post_id` bigint(20) NOT NULL COMMENT '岗位ID',
                                  PRIMARY KEY (`user_id`, `post_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户与岗位关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------
INSERT INTO `sys_user_post` VALUES (1, 1);
INSERT INTO `sys_user_post` VALUES (2, 2);

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
                                  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
                                  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
                                  PRIMARY KEY (`user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户和角色关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO `sys_user_role` VALUES (1, 1);
INSERT INTO `sys_user_role` VALUES (2, 2);

SET FOREIGN_KEY_CHECKS = 1;
