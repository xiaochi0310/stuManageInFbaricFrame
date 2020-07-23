package studentManage

/**
 * @Author: WuNaiChi
 * @Date: 2020/6/22 15:58
 * @Desc:
 */

// 学生信息
type StudentInfo struct {
	AcctId string   `json:"acctId" validate:"required,max=10"` // 学号
	Name   string   `json:"name" validate:"required,max=5"`    // 姓名
	Sex    sexTag   `json:"sex" validate:"required"`           // 性别：map
	Grade  GradeTag `json:"grade" validate:"required"`         // 年级: map 123
	Hobby  string   `json:"hobby" validate:"omitempty,max=50"` //
}

// 学生查询信息
type QueryStudentList struct {
	PageNo   int    `json:"pageNo,omitempty"`   //从第1页开始
	PageSize int    `json:"pageSize,omitempty"` //每页条数
	Bookmark string `json:"bookmark,omitempty"` //couchdb中的书签号
	AcctId   string `json:"acctId"`
	Name     string `json:"name"` // 模糊查询
	Sex      string `json:"sex"`
	Grade    string `json:"grade"`
	Hobby    string `json:"hobby"` // 模糊查询
}

type QueryStudentInfo struct {
	AcctId string `json:"acctId"`
}

type DeleteStudentInfo struct {
	AcctId string `json:"acctId"`
}

// 查询列表的
type RspQueryStudentInfo struct {
	StudentInfo []StudentInfo `json:"studentInfo"`
	Bookmark    string        `json:"bookmark,omitempty"` //couchdb中的书签号
	Count       int32         `json:"count"`
}

type Response struct {
	Code        int         `json:"code,omitempty"`    // 返回码
	Message     string      `json:"message,omitempty"` // 概要信息
	ErrorDetail string      `json:"detail,omitempty"`  // 错误详细信息
	Data        interface{} `json:"data,omitempty"`    // 响应数据
}

// 性别
type sexTag string

const (
	man   sexTag = "man"
	woman sexTag = "woman"
)

var sexMap = map[sexTag]string{
	man:   "男",
	woman: "女",
}

// 年级
type GradeTag string

const (
	one   GradeTag = "1"
	twe   GradeTag = "2"
	three GradeTag = "3"
)

var GradeMap = map[GradeTag]string{
	one:   "一年级",
	twe:   "二年级",
	three: "三年级",
}

const (
	InputParamError             = "参数错误"
	GetClientFailed             = "获取客户端失败"
	QueryStudentFromChainError  = "查询学生信息失败"
	DeleteStudentFromChainError = "删除学生信息失败"
	CommitStudentFromChainError = "上链学生信息失败"
	UpdateStudentFromChainError = "更新学生信息失败"
)

const (
	CommitStudentInfo    = "createStudentInfo"
	UpdateStudentInfo    = "updateStudentInfo"
	QueryStudentInfoList = "getStudentInfoList"
	QueryStuInfo         = "queryStudentInfo"
	DeleteStuInfo        = "deleteStudentInfo"
)
