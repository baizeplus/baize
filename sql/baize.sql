
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for gen_table
-- ----------------------------
DROP TABLE IF EXISTS `gen_table`;
CREATE TABLE `gen_table`  (
                              `table_id` char(16) NOT NULL COMMENT '编号',
                              `table_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '表名称',
                              `parent_menu_id` char(16) NULL DEFAULT NULL COMMENT '父菜单ID',
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
                              `create_by` char(16) NULL DEFAULT NULL COMMENT '创建者',
                              `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                              `update_by` char(16) NULL DEFAULT NULL COMMENT '更新者',
                              `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                              `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
                              PRIMARY KEY (`table_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '代码生成业务表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gen_table
-- ----------------------------

-- ----------------------------
-- Table structure for gen_table_column
-- ----------------------------
DROP TABLE IF EXISTS `gen_table_column`;
CREATE TABLE `gen_table_column`  (
                                     `column_id` char(16) NOT NULL COMMENT '编号',
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
                                     `create_by` char(16) NULL DEFAULT NULL COMMENT '创建者',
                                     `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                                     `update_by` char(16) NULL DEFAULT NULL COMMENT '更新者',
                                     `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                                     `html_field` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
                                     PRIMARY KEY (`column_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '代码生成业务表字段' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gen_table_column
-- ----------------------------

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
                               `config_id` char(16) NOT NULL COMMENT '参数主键',
                               `config_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '参数名称',
                               `config_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '参数键名',
                               `config_value` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '参数键值',
                               `config_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'N' COMMENT '系统内置（Y是 N否）',
                               `create_by` char(16) NULL DEFAULT NULL COMMENT '创建者',
                               `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                               `update_by` char(16) NULL DEFAULT NULL COMMENT '更新者',
                               `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                               `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
                               PRIMARY KEY (`config_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '参数配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES ('0000000000000001', '主框架页-默认皮肤样式名称', 'sys.index.skinName', 'skin-blue', 'Y', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow');
INSERT INTO `sys_config` VALUES ('0000000000000002', '用户管理-账号初始密码', 'sys.user.initPassword', '123456', 'Y', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '初始化密码 123456');
INSERT INTO `sys_config` VALUES ('0000000000000003', '主框架页-侧边栏主题', 'sys.index.sideTheme', 'theme-dark', 'Y', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '深色主题theme-dark，浅色主题theme-light');
INSERT INTO `sys_config` VALUES ('0000000000000004', '账号自助-验证码开关', 'sys.account.captchaEnabled', 'false', 'Y', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', '2024-03-09 13:17:45', '是否开启验证码功能（true开启，false关闭）');
INSERT INTO `sys_config` VALUES ('0000000000000005', '账号自助-是否开启用户注册功能', 'sys.account.registerUser', 'true', 'Y', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', '2024-03-09 13:24:50', '是否开启注册用户功能（true开启，false关闭）');

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
                             `dept_id` char(16) NOT NULL COMMENT '部门id',
                             `parent_id` char(16) NULL DEFAULT '0000000000000000' COMMENT '父部门id',
                             `ancestors` varchar(320) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '祖级列表',
                             `dept_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '部门名称',
                             `order_num` int(11) NULL DEFAULT 0 COMMENT '显示顺序',
                             `leader` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '负责人',
                             `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '联系电话',
                             `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '邮箱',
                             `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
                             `del_flag` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
                             `create_by` char(16) NULL DEFAULT NULL COMMENT '创建者',
                             `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                             `update_by` char(16) NULL DEFAULT NULL COMMENT '更新者',
                             `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                             PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '部门表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES ('0000000000000100', '0000000000000000', '0000000000000000', '白泽科技', 0, '白泽', '15888888888', 'bz@qq.com', '0', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL);
INSERT INTO `sys_dept` VALUES ('0000000000000101', '0000000000000100', '0000000000000000,0000000000000100', '深圳总公司', 1, '白泽', '15888888888', 'bz@qq.com', '0', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL);
INSERT INTO `sys_dept` VALUES ('0000000000000102', '0000000000000100', '0000000000000000,0000000000000100', '长沙分公司', 2, '白泽', '15888888888', 'bz@qq.com', '0', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL);
INSERT INTO `sys_dept` VALUES ('0000000000000103', '0000000000000101', '0000000000000000,0000000000000100,0000000000000101', '研发部门', 1, '白泽', '15888888888', 'bz@qq.com', '0', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL);
INSERT INTO `sys_dept` VALUES ('0000000000000104', '0000000000000101', '0000000000000000,0000000000000100,0000000000000101', '市场部门', 2, '白泽', '15888888888', 'bz@qq.com', '0', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL);
INSERT INTO `sys_dept` VALUES ('0000000000000105', '0000000000000101', '0000000000000000,0000000000000100,0000000000000101', '测试部门', 3, '白泽', '15888888888', 'bz@qq.com', '0', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL);
INSERT INTO `sys_dept` VALUES ('0000000000000106', '0000000000000101', '0000000000000000,0000000000000100,0000000000000101', '财务部门', 4, '白泽', '15888888888', 'bz@qq.com', '0', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL);
INSERT INTO `sys_dept` VALUES ('0000000000000107', '0000000000000101', '0000000000000000,0000000000000100,0000000000000101', '运维部门', 5, '白泽', '15888888888', 'bz@qq.com', '0', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL);
INSERT INTO `sys_dept` VALUES ('0000000000000108', '0000000000000102', '0000000000000000,0000000000000100,0000000000000102', '市场部门', 1, '白泽', '15888888888', 'bz@qq.com', '0', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL);
INSERT INTO `sys_dept` VALUES ('0000000000000109', '0000000000000102', '0000000000000000,0000000000000100,0000000000000102', '财务部门', 2, '白泽', '15888888888', 'bz@qq.com', '0', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
                                  `dict_code` char(16) NOT NULL COMMENT '字典编码',
                                  `dict_sort` int(11) NULL DEFAULT 0 COMMENT '字典排序',
                                  `dict_label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '字典标签',
                                  `dict_value` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '字典键值',
                                  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '字典类型',
                                  `css_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '样式属性（其他样式扩展）',
                                  `list_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '表格回显样式',
                                  `is_default` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'N' COMMENT '是否默认（Y是 N否）',
                                  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
                                  `create_by` char(16) NULL DEFAULT NULL COMMENT '创建者',
                                  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                                  `update_by` char(16) NULL DEFAULT NULL COMMENT '更新者',
                                  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                                  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '备注',
                                  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '字典数据表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES ('0000000000000001', 1, '男', '0', 'sys_user_sex', '', '', 'Y', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '性别男');
INSERT INTO `sys_dict_data` VALUES ('0000000000000002', 2, '女', '1', 'sys_user_sex', '', '', 'N', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '性别女');
INSERT INTO `sys_dict_data` VALUES ('0000000000000003', 3, '未知', '2', 'sys_user_sex', '', '', 'N', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '性别未知');
INSERT INTO `sys_dict_data` VALUES ('0000000000000006', 1, '正常', '0', 'sys_normal_disable', '', 'primary', 'Y', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '正常状态');
INSERT INTO `sys_dict_data` VALUES ('0000000000000007', 2, '停用', '1', 'sys_normal_disable', '', 'danger', 'N', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '停用状态');
INSERT INTO `sys_dict_data` VALUES ('0000000000000008', 1, '正常', '0', 'sys_job_status', '', 'primary', 'Y', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '正常状态');
INSERT INTO `sys_dict_data` VALUES ('0000000000000009', 2, '暂停', '1', 'sys_job_status', '', 'danger', 'N', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '停用状态');
INSERT INTO `sys_dict_data` VALUES ('0000000000000010', 1, '默认', 'DEFAULT', 'sys_job_group', '', '', 'Y', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '默认分组');
INSERT INTO `sys_dict_data` VALUES ('0000000000000011', 2, '系统', 'SYSTEM', 'sys_job_group', '', '', 'N', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '系统分组');
INSERT INTO `sys_dict_data` VALUES ('0000000000000012', 1, '是', 'Y', 'sys_yes_no', '', 'primary', 'Y', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '系统默认是');
INSERT INTO `sys_dict_data` VALUES ('0000000000000013', 2, '否', 'N', 'sys_yes_no', '', 'danger', 'N', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '系统默认否');
INSERT INTO `sys_dict_data` VALUES ('0000000000000014', 1, '通知', '1', 'sys_notice_type', '', 'warning', 'Y', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '通知');
INSERT INTO `sys_dict_data` VALUES ('0000000000000015', 2, '公告', '2', 'sys_notice_type', '', 'success', 'N', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '公告');
INSERT INTO `sys_dict_data` VALUES ('0000000000000018', 0, '其他', '0', 'sys_oper_type', '', 'info', 'N', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '其他操作');
INSERT INTO `sys_dict_data` VALUES ('0000000000000019', 1, '新增', '1', 'sys_oper_type', '', 'info', 'N', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '新增操作');
INSERT INTO `sys_dict_data` VALUES ('0000000000000020', 2, '修改', '2', 'sys_oper_type', '', 'info', 'N', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '修改操作');
INSERT INTO `sys_dict_data` VALUES ('0000000000000021', 3, '删除', '3', 'sys_oper_type', '', 'danger', 'N', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '删除操作');
INSERT INTO `sys_dict_data` VALUES ('0000000000000022', 4, '强退', '4', 'sys_oper_type', '', 'danger', 'N', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', '2024-04-12 14:33:31', '强退操作');
INSERT INTO `sys_dict_data` VALUES ('0000000000000023', 5, '清空数据', '5', 'sys_oper_type', '', 'danger', 'N', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '清空操作');
INSERT INTO `sys_dict_data` VALUES ('0000000000000028', 1, '成功', '0', 'sys_common_status', '', 'primary', 'N', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '正常状态');
INSERT INTO `sys_dict_data` VALUES ('0000000000000029', 2, '失败', '1', 'sys_common_status', '', 'danger', 'N', '0', '0000000000000001', '2024-02-08 04:10:56', '0000000000000001', NULL, '停用状态');

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type`  (
                                  `dict_id` char(16) NOT NULL COMMENT '字典主键',
                                  `dict_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '字典名称',
                                  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '字典类型',
                                  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
                                  `create_by` char(16) NULL DEFAULT NULL COMMENT '创建者',
                                  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                                  `update_by` char(16) NULL DEFAULT NULL COMMENT '更新者',
                                  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                                  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
                                  PRIMARY KEY (`dict_id`) USING BTREE,
                                  UNIQUE INDEX `dict_type`(`dict_type`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '字典类型表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES ('0000000000000001', '用户性别', 'sys_user_sex', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL, '用户性别列表');
INSERT INTO `sys_dict_type` VALUES ('0000000000000003', '系统开关', 'sys_normal_disable', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL, '系统开关列表');
INSERT INTO `sys_dict_type` VALUES ('0000000000000004', '任务状态', 'sys_job_status', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL, '任务状态列表');
INSERT INTO `sys_dict_type` VALUES ('0000000000000005', '任务分组', 'sys_job_group', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL, '任务分组列表');
INSERT INTO `sys_dict_type` VALUES ('0000000000000006', '系统是否', 'sys_yes_no', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL, '系统是否列表');
INSERT INTO `sys_dict_type` VALUES ('0000000000000007', '通知类型', 'sys_notice_type', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL, '通知类型列表');
INSERT INTO `sys_dict_type` VALUES ('0000000000000009', '操作类型', 'sys_oper_type', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL, '操作类型列表');
INSERT INTO `sys_dict_type` VALUES ('0000000000000010', '系统状态', 'sys_common_status', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL, '登录状态列表');

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job` (
                           `job_id` char(16) COLLATE utf8_unicode_ci NOT NULL COMMENT '任务ID',
                           `job_name` varchar(64) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '任务名称',
                           `job_params` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '参数',
                           `invoke_target` varchar(500) COLLATE utf8_unicode_ci NOT NULL COMMENT '调用目标字符串',
                           `cron_expression` varchar(255) COLLATE utf8_unicode_ci DEFAULT '' COMMENT 'cron执行表达式',
                           `status` char(1) COLLATE utf8_unicode_ci DEFAULT '0' COMMENT '状态（0正常 1暂停）',
                           `create_by` char(16) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '创建者',
                           `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                           `update_by` char(16) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '更新者',
                           `update_time` datetime DEFAULT NULL COMMENT '更新时间',
                           PRIMARY KEY (`job_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='定时任务调度表';

-- ----------------------------
-- Records of sys_job
-- ----------------------------
BEGIN;
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `invoke_target`, `cron_expression`, `status`, `create_by`, `create_time`, `update_by`, `update_time`) VALUES ('67RBSK3TD4M01000', '无参调用', '', 'NoParams', '* * * * *', '0', '1', '2025-09-05 10:37:16', '0000000000000001', '2025-09-06 16:06:27');
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `invoke_target`, `cron_expression`, `status`, `create_by`, `create_time`, `update_by`, `update_time`) VALUES ('67RBSK3TD4M01001', '有参测试', 'baize,18', 'Params', '* * * * *', '1', '1', '2025-09-05 10:37:48', '0000000000000001', '2025-09-06 17:57:59');
COMMIT;

-- ----------------------------
-- Table structure for sys_job_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_job_log`;
CREATE TABLE `sys_job_log` (
                               `job_log_id` char(16) NOT NULL COMMENT '任务日志ID',
                               `job_id` char(16) NOT NULL COMMENT '任务ID',
                               `job_name` varchar(64) NOT NULL COMMENT '任务名称',
                               `job_params` varchar(255) NOT NULL COMMENT '任务组名',
                               `invoke_target` varchar(500) NOT NULL COMMENT '调用目标字符串',
                               `status` char(1) DEFAULT '0' COMMENT '执行状态（0正常 1失败）',
                               `exception_info` varchar(2000) DEFAULT '' COMMENT '异常信息',
                               `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                               `cost_time` bigint(20) DEFAULT '0' COMMENT '耗时（毫秒）',
                               PRIMARY KEY (`job_log_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='定时任务调度日志表';

-- ----------------------------
-- Records of sys_job_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_logininfor
-- ----------------------------
DROP TABLE IF EXISTS `sys_logininfor`;
CREATE TABLE `sys_logininfor`  (
                                   `info_id` char(16) NOT NULL COMMENT '访问ID',
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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '系统访问记录' ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for sys_notice
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice`;
CREATE TABLE `sys_notice`  (
                               `id` char(16) NOT NULL COMMENT '公告ID',
                               `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '公告标题',
                               `type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '公告类型（1通知 2公告）',
                               `txt` longblob NULL COMMENT '公告内容',
                               `dept_id` char(16) NULL DEFAULT NULL COMMENT '发件人所在部门',
                               `dept_ids` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '发送部门',
                               `create_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '创建者名称',
                               `create_by` char(16) NULL DEFAULT NULL COMMENT '创建者',
                               `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                               PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '通知公告表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_notice
-- ----------------------------

-- ----------------------------
-- Table structure for sys_notice_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice_user`;
CREATE TABLE `sys_notice_user`  (
                                    `user_id` char(16) NOT NULL,
                                    `notice_id` char(16) NOT NULL,
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
                                 `oper_id` char(16) NOT NULL COMMENT '日志主键',
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
                                 `user_id` char(16) NULL DEFAULT NULL COMMENT '用户ID',
                                 `oper_time` datetime(0) NULL DEFAULT NULL COMMENT '操作时间',
                                 `cost_time` bigint(20) NULL DEFAULT 0 COMMENT '消耗时间',
                                 PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '操作日志记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_permission`;
CREATE TABLE `sys_permission`  (
                                   `permission_id` char(16) NOT NULL COMMENT '主键',
                                   `permission_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '权限名称',
                                   `parent_id` char(16) NULL DEFAULT NULL COMMENT '父ID',
                                   `permission` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '权限标识符',
                                   `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
                                   `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0' COMMENT '状态',
                                   `create_by` char(16) NULL DEFAULT NULL COMMENT '创建人',
                                   `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                                   `update_by` char(16) NULL DEFAULT NULL COMMENT '更新人',
                                   `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                                   PRIMARY KEY (`permission_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_permission
-- ----------------------------
INSERT INTO `sys_permission` VALUES ('0000000000000001', '系统管理', '0000000000000000', 'system', 1, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000000002', '系统监控', '0000000000000000', 'monitor', 2, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000000003', '系统工具', '0000000000000000', 'tool', 3, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000000100', '用户管理', '0000000000000001', 'system:user', 1, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000000101', '角色管理', '0000000000000001', 'system:role', 2, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000000102', '权限管理', '0000000000000001', 'system:permission', 3, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000000103', '部门管理', '0000000000000001', 'system:dept', 4, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000000104', '岗位管理', '0000000000000001', 'system:post', 5, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000000105', '字典管理', '0000000000000001', 'system:dict', 6, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000000106', '参数设置', '0000000000000001', 'system:config', 7, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000000107', '通知公告', '0000000000000001', 'system:notice', 8, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000000108', '日志管理', '0000000000000001', 'system:monitor', 9, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000000109', '在线用户', '0000000000000002', 'monitor:online', 1, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000000110', '定时任务', '0000000000000002', 'monitor:job', 2, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000000112', '服务监控', '0000000000000002', 'monitor:server', 3, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-03-26 22:37:27');
INSERT INTO `sys_permission` VALUES ('0000000000000115', '表单构建', '0000000000000003', 'tool:build', 1, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000000116', '代码生成', '0000000000000003', 'tool:gen', 2, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000000117', '系统接口', '0000000000000003', 'tool:swagger', 3, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000000500', '操作日志', '0000000000000108', 'system:monitor:operlog', 1, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000000501', '登录日志', '0000000000000108', 'system:monitor:logininfor', 2, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001000', '用户查询', '0000000000000100', 'system:user:query', 1, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001001', '用户新增', '0000000000000100', 'system:user:add', 2, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001002', '用户修改', '0000000000000100', 'system:user:edit', 3, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001003', '用户删除', '0000000000000100', 'system:user:remove', 4, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001004', '用户导出', '0000000000000100', 'system:user:export', 5, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001005', '用户导入', '0000000000000100', 'system:user:import', 6, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001006', '重置密码', '0000000000000100', 'system:user:resetPwd', 7, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001007', '角色查询', '0000000000000101', 'system:role:query', 1, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001008', '角色新增', '0000000000000101', 'system:role:add', 2, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001009', '角色修改', '0000000000000101', 'system:role:edit', 3, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001010', '角色删除', '0000000000000101', 'system:role:remove', 4, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001011', '角色导出', '0000000000000101', 'system:role:export', 5, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001012', '菜单查询', '0000000000000102', 'system:permission:query', 1, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001013', '菜单新增', '0000000000000102', 'system:permission:add', 2, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001014', '菜单修改', '0000000000000102', 'system:permission:edit', 3, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001015', '菜单删除', '0000000000000102', 'system:permission:remove', 4, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001016', '部门查询', '0000000000000103', 'system:dept:query', 1, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001017', '部门新增', '0000000000000103', 'system:dept:add', 2, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001018', '部门修改', '0000000000000103', 'system:dept:edit', 3, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001019', '部门删除', '0000000000000103', 'system:dept:remove', 4, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001020', '岗位查询', '0000000000000104', 'system:post:query', 1, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001021', '岗位新增', '0000000000000104', 'system:post:add', 2, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001022', '岗位修改', '0000000000000104', 'system:post:edit', 3, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001023', '岗位删除', '0000000000000104', 'system:post:remove', 4, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001024', '岗位导出', '0000000000000104', 'system:post:export', 5, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001025', '字典查询', '0000000000000105', 'system:dict:query', 1, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001026', '字典新增', '0000000000000105', 'system:dict:add', 2, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001027', '字典修改', '0000000000000105', 'system:dict:edit', 3, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001028', '字典删除', '0000000000000105', 'system:dict:remove', 4, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001029', '字典导出', '0000000000000105', 'system:dict:export', 5, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001030', '参数查询', '0000000000000106', 'system:config:query', 1, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001031', '参数新增', '0000000000000106', 'system:config:add', 2, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001032', '参数修改', '0000000000000106', 'system:config:edit', 3, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001033', '参数删除', '0000000000000106', 'system:config:remove', 4, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001034', '参数导出', '0000000000000106', 'system:config:export', 5, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001035', '公告查询', '0000000000000107', 'system:notice:query', 1, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001036', '公告新增', '0000000000000107', 'system:notice:add', 2, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001037', '公告修改', '0000000000000107', 'system:notice:edit', 3, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001038', '公告删除', '0000000000000107', 'system:notice:remove', 4, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001039', '操作查询', '0000000000000500', 'monitor:operlog:query', 1, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001040', '操作删除', '0000000000000500', 'monitor:operlog:remove', 2, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001041', '日志导出', '0000000000000500', 'monitor:operlog:export', 3, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001042', '登录查询', '0000000000000501', 'monitor:logininfor:query', 1, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001043', '登录删除', '0000000000000501', 'monitor:logininfor:remove', 2, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001044', '日志导出', '0000000000000501', 'monitor:logininfor:export', 3, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001046', '在线查询', '0000000000000109', 'monitor:online:query', 1, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001047', '批量强退', '0000000000000109', 'monitor:online:batchLogout', 2, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001048', '单条强退', '0000000000000109', 'monitor:online:forceLogout', 3, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001049', '任务查询', '0000000000000110', 'monitor:job:query', 1, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001050', '任务新增', '0000000000000110', 'monitor:job:add', 2, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001051', '任务修改', '0000000000000110', 'monitor:job:edit', 3, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001052', '任务删除', '0000000000000110', 'monitor:job:remove', 4, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001053', '状态修改', '0000000000000110', 'monitor:job:changeStatus', 5, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001054', '任务导出', '0000000000000110', 'monitor:job:export', 6, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001055', '生成查询', '0000000000000116', 'tool:gen:query', 1, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001056', '生成修改', '0000000000000116', 'tool:gen:edit', 2, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001057', '生成删除', '0000000000000116', 'tool:gen:remove', 3, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001058', '导入代码', '0000000000000116', 'tool:gen:import', 4, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001059', '预览代码', '0000000000000116', 'tool:gen:preview', 5, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');
INSERT INTO `sys_permission` VALUES ('0000000000001060', '生成代码', '0000000000000116', 'tool:gen:code', 6, '0', '0000000000000001', '2025-02-28 13:38:05', '0000000000000001', '2025-02-28 13:38:05');

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
                             `post_id` char(16) NOT NULL COMMENT '岗位ID',
                             `post_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '岗位编码',
                             `post_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '岗位名称',
                             `post_sort` int(11) NOT NULL COMMENT '显示顺序',
                             `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '状态（0正常 1停用）',
                             `create_by` char(16) NULL DEFAULT NULL COMMENT '创建者',
                             `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                             `update_by` char(16) NULL DEFAULT NULL COMMENT '更新者',
                             `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                             `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
                             PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '岗位信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES ('0000000000000001', 'ceo', '董事长', 1, '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL, '');
INSERT INTO `sys_post` VALUES ('0000000000000002', 'se', '项目经理', 2, '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL, '');
INSERT INTO `sys_post` VALUES ('0000000000000003', 'hr', '人力资源', 3, '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL, '');
INSERT INTO `sys_post` VALUES ('0000000000000004', 'user', '普通员工', 4, '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', '2024-02-25 10:40:41', '');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
                             `role_id` char(16) NOT NULL COMMENT '角色ID',
                             `role_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色名称',
                             `role_sort` int(11) NOT NULL COMMENT '显示顺序',
                             `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色状态（0正常 1停用）',
                             `del_flag` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
                             `create_by` char(16) NULL DEFAULT NULL COMMENT '创建者',
                             `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                             `update_by` char(16) NULL DEFAULT NULL COMMENT '更新者',
                             `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                             `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
                             PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES ('0000000000000001', 'admin', 1, '0', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', NULL, '超级管理员');
INSERT INTO `sys_role` VALUES ('0000000000000002', 'common', 2, '0', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', '2024-03-04 14:16:58', '普通角色');

-- ----------------------------
-- Table structure for sys_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_permission`;
CREATE TABLE `sys_role_permission`  (
                                        `role_id` char(16) NOT NULL COMMENT '角色ID',
                                        `permission_id` char(16) NOT NULL COMMENT '权限ID',
                                        PRIMARY KEY (`role_id`, `permission_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色和权限关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_permission
-- ----------------------------
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000001');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000002');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000003');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000004');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000100');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000101');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000102');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000103');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000104');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000105');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000106');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000107');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000108');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000109');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000110');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000111');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000112');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000113');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000114');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000115');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000116');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000500');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000000501');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001000');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001001');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001002');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001003');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001004');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001005');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001006');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001007');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001008');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001009');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001010');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001011');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001012');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001013');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001014');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001015');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001016');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001017');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001018');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001019');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001020');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001021');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001022');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001023');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001024');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001025');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001026');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001027');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001028');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001029');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001030');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001031');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001032');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001033');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001034');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001035');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001036');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001037');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001038');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001039');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001040');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001041');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001042');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001043');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001044');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001045');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001046');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001047');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001048');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001049');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001050');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001051');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001052');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001053');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001054');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001055');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001056');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001057');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001058');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001059');
INSERT INTO `sys_role_permission` VALUES ('0000000000000002', '0000000000001060');

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
                             `user_id` char(16) NOT NULL COMMENT '用户ID',
                             `dept_id` char(16) NULL DEFAULT NULL COMMENT '部门ID',
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
                             `create_by` char(16) NULL DEFAULT NULL COMMENT '创建者',
                             `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                             `update_by` char(16) NULL DEFAULT NULL COMMENT '更新者',
                             `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                             `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '备注',
                             PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('0000000000000001', '0000000000000105', 'admin', '白泽', 'bz@163.com', '15888888887', '0', '', '$2a$14$kOhr1gqDAGYara15BrSuLuI7uzOC7gfAyzJGiLsypO32mEb6O02LO', '1', '0', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', '2024-03-15 09:31:36', '管理员');
INSERT INTO `sys_user` VALUES ('0000000000000002', '0000000000000105', 'bz', '白泽', 'bz@qq.com', '15666666666', '1', '', '$2a$10$7JB720yubVSZvUI0rEqK/.VqGOZTH.ulu33dHOiBE8ByOhJIrdAu2', '2', '0', '0', '0000000000000001', '2024-02-08 04:10:55', '0000000000000001', '2025-04-06 22:07:27', '测试员');

-- ----------------------------
-- Table structure for sys_user_dept_scope
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_dept_scope`;
CREATE TABLE `sys_user_dept_scope`  (
                                        `user_id` char(16) NOT NULL COMMENT '用户ID',
                                        `dept_id` char(16) NOT NULL COMMENT '部门ID',
                                        PRIMARY KEY (`user_id`, `dept_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户和部门关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_dept_scope
-- ----------------------------
INSERT INTO `sys_user_dept_scope` VALUES ('0000000000000002', '0000000000000101');
INSERT INTO `sys_user_dept_scope` VALUES ('0000000000000002', '0000000000000102');
INSERT INTO `sys_user_dept_scope` VALUES ('0000000000000002', '0000000000000103');
INSERT INTO `sys_user_dept_scope` VALUES ('0000000000000002', '0000000000000104');
INSERT INTO `sys_user_dept_scope` VALUES ('0000000000000002', '0000000000000108');
INSERT INTO `sys_user_dept_scope` VALUES ('0000000000000002', '0000000000000109');

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post`  (
                                  `user_id` char(16) NOT NULL COMMENT '用户ID',
                                  `post_id` char(16) NOT NULL COMMENT '岗位ID',
                                  PRIMARY KEY (`user_id`, `post_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户与岗位关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------
INSERT INTO `sys_user_post` VALUES ('0000000000000001', '0000000000000001');
INSERT INTO `sys_user_post` VALUES ('0000000000000002', '0000000000000002');

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
                                  `user_id` char(16) NOT NULL COMMENT '用户ID',
                                  `role_id` char(16) NOT NULL COMMENT '角色ID',
                                  PRIMARY KEY (`user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户和角色关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO `sys_user_role` VALUES ('0000000000000001', '0000000000000001');
INSERT INTO `sys_user_role` VALUES ('0000000000000002', '0000000000000002');

SET FOREIGN_KEY_CHECKS = 1;
