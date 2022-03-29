// Copyright (c) 2015-2022 CloudJ Technology Co., Ltd.
// THIS FILE GENERATED BY 'gen-lang' DO NOT EDIT

package e

var langs = []string{"zh-cn", "en-us"}

var errorMsgs = map[int]map[string]string{
	InternalError: {
		"en-US": "internal error",
		"zh-CN": "未知错误",
	},
	ObjectAlreadyExists: {
		"en-US": "object already exists",
		"zh-CN": "对象已存在",
	},
	ObjectNotExists: {
		"en-US": "object not exists",
		"zh-CN": "对象不存在",
	},
	ObjectNotExistsOrNoPerm: {
		"en-US": "object not exists or no permisson to access object",
		"zh-CN": "对象不存在或者无权限",
	},
	ObjectDisabled: {
		"en-US": "object disabled",
		"zh-CN": "对象已禁用",
	},
	JSONParseError: {
		"en-US": "invalid json",
		"zh-CN": "JSON 数据解析出错",
	},
	URLParseError: {
		"en-US": "invalid url",
		"zh-CN": "URL解析出错",
	},
	NotImplement: {
		"en-US": "not implemented",
		"zh-CN": "暂未实现",
	},
	DBError: {
		"en-US": "database error",
		"zh-CN": "数据库错误",
	},
	DBAttrValidateErr: {
		"en-US": "database attribute invalidate",
		"zh-CN": "字段验证错误",
	},
	BadOrgId: {
		"en-US": "invalid organization id",
		"zh-CN": "组织 ID 错误",
	},
	BadProjectId: {
		"en-US": "invalid project id",
		"zh-CN": "项目 ID 错误",
	},
	BadTemplateId: {
		"en-US": "invalid templated id",
		"zh-CN": "模板 ID 错误",
	},
	BadEnvId: {
		"en-US": "invalid environment id",
		"zh-CN": "环境 ID 错误",
	},
	BadParam: {
		"en-US": "invalid or missing parameter",
		"zh-CN": "无效参数",
	},
	TemplateNameRepeat: {
		"en-US": "duplicated template name",
		"zh-CN": "云模版名称重复",
	},
	TemplateWorkdirError: {
		"en-US": "invalid template working directory",
		"zh-CN": "工作目录校验失败",
	},
	BadRequest: {
		"en-US": "bad request",
		"zh-CN": "无效请求",
	},
	InvalidPipeline: {
		"en-US": "invalid pipeline definition",
		"zh-CN": "pipeline 格式错误",
	},
	InvalidPipelineVersion: {
		"en-US": "invalid pipeline version",
		"zh-CN": "不支持的 pipeline 版本",
	},
	InvalidExportVersion: {
		"en-US": "data format version not supported",
		"zh-CN": "不支持的导出数据版本",
	},
	DataTooLong: {
		"en-US": "request data too large",
		"zh-CN": "内容过长",
	},
	NameTooLong: {
		"en-US": "name too long",
		"zh-CN": "名称过长",
	},
	RemarkTooLong: {
		"en-US": "remark too long",
		"zh-CN": "备注过长",
	},
	TagTooLong: {
		"en-US": "tag too long",
		"zh-CN": "标签过长",
	},
	TagTooMuch: {
		"en-US": "number of tags exceeds system limit",
		"zh-CN": "标签过多",
	},
	IOError: {
		"en-US": "unexpected io error",
		"zh-CN": "io 错误",
	},
	TooManyRetries: {
		"en-US": "too many retries",
		"zh-CN": "达到最大重试次数",
	},
	EncryptError: {
		"en-US": "data encryption error",
		"zh-CN": "数据加密错误",
	},
	DecryptError: {
		"en-US": "data dencryption error",
		"zh-CN": "数据解密错误",
	},
	MailServerError: {
		"en-US": "failed to send email",
		"zh-CN": "邮件服务错误",
	},
	InvalidAccessKeyId: {
		"en-US": "invalid access key id",
		"zh-CN": "AccessKeyId错误",
	},
	InvalidAccessKeySecret: {
		"en-US": "invalid access key secret",
		"zh-CN": "AccessKeySecret错误",
	},
	ForbiddenAccessKey: {
		"en-US": "access denied",
		"zh-CN": "AccessKey权限不足",
	},
	InvalidToken: {
		"en-US": "invalid token",
		"zh-CN": "凭证无效",
	},
	InvalidTokenScope: {
		"en-US": "invalid token scope",
		"zh-CN": "凭证 scope 不匹配",
	},
	InvalidOrgId: {
		"en-US": "invalid organization id",
		"zh-CN": "无效的组织",
	},
	TokenExpired: {
		"en-US": "token expired",
		"zh-CN": "凭证已过期",
	},
	ColValidateError: {
		"en-US": "database column invalid",
		"zh-CN": "字段校验错误",
	},
	InvalidPassword: {
		"en-US": "invalid account or password",
		"zh-CN": "无效的邮箱或密码",
	},
	InvalidColumn: {
		"en-US": "invalid column",
		"zh-CN": "无效的字段名",
	},
	InvalidOperation: {
		"en-US": "invalid operation",
		"zh-CN": "无效操作",
	},
	UserAlreadyExists: {
		"en-US": "user already exists",
		"zh-CN": "用户已存在",
	},
	UserNotExists: {
		"en-US": "user does not exists",
		"zh-CN": "用户不存在",
	},
	UserEmailDuplicate: {
		"en-US": "email address already exists",
		"zh-CN": "用户邮箱已存在",
	},
	UserEmailDuplicateInactive: {
		"en-US": "email address inactive",
		"zh-CN": "无效的用户邮箱",
	},
	UserInvalidStatus: {
		"en-US": "invalid user status",
		"zh-CN": "无效的用户状态",
	},
	UserInactive: {
		"en-US": "user is inactive",
		"zh-CN": "用户未激活",
	},
	UserDisabled: {
		"en-US": "user is disabled",
		"zh-CN": "用户已禁用",
	},
	InvalidPasswordFormat: {
		"en-US": "password does not meet complexity requirement",
		"zh-CN": "密码格式错误",
	},
	UserActivated: {
		"en-US": "user already activated",
		"zh-CN": "账号已激活",
	},
	InvalidRoleName: {
		"en-US": "invalid role name",
		"zh-CN": "无效角色名",
	},
	RoleNameDuplicate: {
		"en-US": "role name already exists",
		"zh-CN": "角色名重复",
	},
	PermissionDeny: {
		"en-US": "access denied",
		"zh-CN": "无权限",
	},
	PermDenyApproval: {
		"en-US": "access denied",
		"zh-CN": "无审批权限",
	},
	ValidateError: {
		"en-US": "login validation fail",
		"zh-CN": "验证失败",
	},
	OrganizationAlreadyExists: {
		"en-US": "organization already exists",
		"zh-CN": "组织已存在",
	},
	OrganizationNotExists: {
		"en-US": "organization not exists",
		"zh-CN": "组织不存在",
	},
	OrganizationDisabled: {
		"en-US": "organization disabled",
		"zh-CN": "组织被禁用",
	},
	OrganizationInvalidStatus: {
		"en-US": "organization invalid status",
		"zh-CN": "无效的组织状态",
	},
	InvalidOrganizationId: {
		"en-US": "invalid organization id",
		"zh-CN": "无效的组织ID",
	},
	NameDuplicate: {
		"en-US": "user already exists",
		"zh-CN": "名称重复",
	},
	TaskStepNotExists: {
		"en-US": "task step does not exists",
		"zh-CN": "步骤不存在",
	},
	InvalidProjectId: {
		"en-US": "invalid project id",
		"zh-CN": "无效的项目id",
	},
	TaskNotHaveStep: {
		"en-US": "task have no steps",
		"zh-CN": "任务无步骤",
	},
	TaskAborting: {
		"en-US": "task is aborting",
		"zh-CN": "任务正在中止",
	},
	TaskAborted: {
		"en-US": "task aborted",
		"zh-CN": "任务已中止",
	},
	TaskCannotAbort: {
		"en-US": "task cannot abort",
		"zh-CN": "任务当前无法中止",
	},
	TemplateAlreadyExists: {
		"en-US": "template already exists",
		"zh-CN": "模板名称重复",
	},
	HCLParseError: {
		"en-US": "hcl parse error",
		"zh-CN": "模板语法解析错误",
	},
	VariableAlreadyExists: {
		"en-US": "variable already exists",
		"zh-CN": "变量已存在",
	},
	VariableAliasDuplicate: {
		"en-US": "variable alias name already exists",
		"zh-CN": "变量别名重复",
	},
	VariableScopeConflict: {
		"en-US": "variable scope conflict",
		"zh-CN": "变量作用域冲突",
	},
	InvalidVarName: {
		"en-US": "invalid variable name",
		"zh-CN": "无效变量名",
	},
	EmptyVarName: {
		"en-US": "variable name is empty",
		"zh-CN": "变量名不可为空",
	},
	EmptyVarValue: {
		"en-US": "variable value is empty",
		"zh-CN": "变量值不可为空",
	},
	ProjectAlreadyExists: {
		"en-US": "project already exists",
		"zh-CN": "项目已存在",
	},
	ProjectNotExists: {
		"en-US": "project not exists",
		"zh-CN": "项目不存在",
	},
	ProjectAliasDuplicate: {
		"en-US": "project name already exists",
		"zh-CN": "项目名称重复",
	},
	ProjectUserAlreadyExists: {
		"en-US": "project user already exists",
		"zh-CN": "项目用户已经存在",
	},
	ProjectUserAliasDuplicate: {
		"en-US": "project name already exists",
		"zh-CN": "项目名称重复",
	},
	TokenAlreadyExists: {
		"en-US": "token already exists",
		"zh-CN": "Token已经存在",
	},
	TokenNotExists: {
		"en-US": "token does not exists",
		"zh-CN": "Token不存在",
	},
	TokenAliasDuplicate: {
		"en-US": "token alias already exists",
		"zh-CN": "Token别名重复",
	},
	TemplateNotExists: {
		"en-US": "template not exists",
		"zh-CN": "模板不存在",
	},
	TemplateDisabled: {
		"en-US": "template disabled",
		"zh-CN": "模板不可用",
	},
	TemplateActiveEnvExists: {
		"en-US": "can not delete template which have active environment",
		"zh-CN": "模板存在活跃环境",
	},
	ConsulConnError: {
		"en-US": "connect to consul error",
		"zh-CN": "consul链接失败",
	},
	EnvAlreadyExists: {
		"en-US": "environment already exists",
		"zh-CN": "环境已经存在",
	},
	EnvNotExists: {
		"en-US": "environment not exists",
		"zh-CN": "环境不存在",
	},
	EnvAliasDuplicate: {
		"en-US": "environment alias already exists",
		"zh-CN": "环境别名重复",
	},
	EnvArchived: {
		"en-US": "environment had been archived",
		"zh-CN": "环境已归档，不允许操作",
	},
	EnvDeploying: {
		"en-US": "deployment in progress",
		"zh-CN": "环境正在部署中，请不要重复发起",
	},
	EnvCheckAutoApproval: {
		"en-US": "auto approval is required",
		"zh-CN": "配置自动纠漂移、推送到分支时重新部署时，必须配置自动审批",
	},
	EnvLockFailedTaskActive: {
		"en-US": "environment lock failed. Active tasks in the environment",
		"zh-CN": "环境锁定失败，环境下有活跃任务",
	},
	TaskAlreadyExists: {
		"en-US": "task already exists",
		"zh-CN": "任务已经存在",
	},
	TaskNotExists: {
		"en-US": "task does not exists",
		"zh-CN": "任务不存在",
	},
	VcsError: {
		"en-US": "unknown VCS error",
		"zh-CN": "VCS仓库错误",
	},
	VcsAddressError: {
		"en-US": "invalid VCS address",
		"zh-CN": "VCS地址错误",
	},
	VcsInvalidToken: {
		"en-US": "VCS token invalid",
		"zh-CN": "VCS令牌无效",
	},
	VcsConnectError: {
		"en-US": "connect to VCS failed",
		"zh-CN": "VCS服务连接失败",
	},
	VcsConnectTimeOut: {
		"en-US": "VCS connection time out",
		"zh-CN": "VCS服务连接超时",
	},
	VcsNotExists: {
		"en-US": "repository does not exist",
		"zh-CN": "VCS仓库不存在",
	},
	VcsDeleteError: {
		"en-US": "VCS deletion failed",
		"zh-CN": "VCS存在相关依赖云模版，无法删除",
	},
	ImportError: {
		"en-US": "failed to import",
		"zh-CN": "导入出错",
	},
	ImportIdDuplicate: {
		"en-US": "import id was duplicated",
		"zh-CN": "id 重复",
	},
	ImportUpdateOrgId: {
		"en-US": "import failed",
		"zh-CN": "同 id 的数据己属于另一组织，无法使用“覆盖”方案(不允许更改组织 id)",
	},
	TaskApproveNotPending: {
		"en-US": "task is not awaiting for approve",
		"zh-CN": "作业状态非待审批，不允许操作",
	},
	KeyAlreadyExists: {
		"en-US": "key already exists",
		"zh-CN": "管理密钥已存在",
	},
	KeyNotExist: {
		"en-US": "key does not exist",
		"zh-CN": "管理密钥不存在",
	},
	KeyAliasDuplicate: {
		"en-US": "key alias exists",
		"zh-CN": "管理密钥名称重复",
	},
	KeyDecryptFail: {
		"en-US": "key decryption fail",
		"zh-CN": "管理密钥解析失败",
	},
	EnvCannotArchiveActive: {
		"en-US": " 无法归档",
		"zh-CN": "环境当前状态活跃",
	},
	InvalidTfVersion: {
		"en-US": "invalid terraform version",
		"zh-CN": "自动选择版本失败",
	},
	PolicyAlreadyExist: {
		"en-US": "policy already exist",
		"zh-CN": "策略已存在",
	},
	PolicyNotExist: {
		"en-US": "policy does not exist",
		"zh-CN": "策略不存在",
	},
	PolicyGroupAlreadyExist: {
		"en-US": "policy group already exist",
		"zh-CN": "策略组已存在",
	},
	PolicyGroupNotExist: {
		"en-US": "policy group does not exist",
		"zh-CN": "策略组不存在",
	},
	PolicyBelongedToAnotherGroup: {
		"en-US": "policy belonged to another group",
		"zh-CN": "策略属于其他策略组",
	},
	PolicyResultAlreadyExist: {
		"en-US": "policy result already exist",
		"zh-CN": "结果已存在",
	},
	PolicyResultNotExist: {
		"en-US": "policy result does not exist",
		"zh-CN": "结果不存在",
	},
	PolicyErrorParseTemplate: {
		"en-US": "template parse error",
		"zh-CN": "模板解析错误",
	},
	PolicyRegoMissingComment: {
		"en-US": "missing comment header in rego file",
		"zh-CN": "Rego脚本头缺失",
	},
	PolicySuppressNotExist: {
		"en-US": "suppress status does not exist",
		"zh-CN": "屏蔽记录不存在",
	},
	PolicySuppressAlreadyExist: {
		"en-US": "suppress status already exist",
		"zh-CN": "屏蔽记录已存在",
	},
	PolicyRelNotExist: {
		"en-US": "policy relation does not exist",
		"zh-CN": "策略关联关系不存在",
	},
	PolicyRelAlreadyExist: {
		"en-US": "policy relation already exist",
		"zh-CN": "策略关联关系已存在",
	},
	PolicyScanNotEnabled: {
		"en-US": "policy scan is not enabled",
		"zh-CN": "扫描未启用",
	},
	CronExpressError: {
		"en-US": "invalid cron expression",
		"zh-CN": "cron定时任务表达式错误",
	},
	CronTaskFailed: {
		"en-US": "cron task failed",
		"zh-CN": "cron定时任务执行失败",
	},
	PolicyMetaInvalid: {
		"en-US": "invalid meta json definition",
		"zh-CN": "策略元数据解析无效",
	},
	PolicyRegoInvalid: {
		"en-US": "invalid rego file",
		"zh-CN": "rego 脚本解析无效",
	},
	SystemConfigNotExist: {
		"en-US": "system config does not exist",
		"zh-CN": "当前配置不存在",
	},
	TemplateKeyIdNotSet: {
		"en-US": "ssh keypair is not setup",
		"zh-CN": "SSH 密钥未配置",
	},
	PolicyGroupDirError: {
		"en-US": "policy not found in the repository",
		"zh-CN": "仓库在当前目录找不到策略文件",
	},
	LdapConnectFailed: {
		"en-US": "ldap servers connect failed",
		"zh-CN": "ldap链接服务端链接失败",
	},
	LdapUpdateFailed: {
		"en-US": "LDAP accounts cannot be modified on the IAC server",
		"zh-CN": "ldap账户不支持在iac服务端进行修改",
	},
}
