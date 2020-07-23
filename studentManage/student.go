package studentManage

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"net/http"
	logs "yuncode/fabric/logs/yxlogs"
)

/**
 * @Author: WuNaiChi
 * @Date: 2020/6/22 15:57
 * @Desc:
 */

type StuController struct {
	DefaultController
}

// 学生信息上链
func (s *StuController) CommitStudentInfo() {
	customerId := s.Ctx.Input.Header("Customer-ID")
	chainCodeId := beego.AppConfig.String("defaultChaincodeID")
	var student = new(StudentInfo)
	err := json.Unmarshal(s.Ctx.Input.RequestBody, student)
	if err != nil {
		logs.Error("Failed to Unmarshal request. err: %v\", err")
		s.ReturnError(&Response{
			Code:        http.StatusBadRequest,
			Message:     InputParamError,
			ErrorDetail: err.Error(),
		})
		return
	}
	service, err := NewService(customerId, "", "")
	if err != nil {
		s.ReturnError(&Response{
			Code:        http.StatusBadRequest,
			Message:     GetClientFailed,
			ErrorDetail: err.Error(),
		})
		return
	}
	err = service.Commit(student, chainCodeId)
	if err != nil {
		s.ReturnError(&Response{
			Code:        http.StatusBadRequest,
			Message:     CommitStudentFromChainError,
			ErrorDetail: err.Error(),
		})
		return
	}
	s.ReturnOK(nil)
}

// 学生信息更新
func (s *StuController) UpdateStudentInfo() {
	customerId := s.Ctx.Input.Header("Customer-ID")
	chainCodeId := beego.AppConfig.String("defaultChaincodeID")
	var student = new(StudentInfo)
	err := json.Unmarshal(s.Ctx.Input.RequestBody, student)
	if err != nil {
		logs.Error("Failed to Unmarshal request. err: %v\", err")
		s.ReturnError(&Response{
			Code:        http.StatusBadRequest,
			Message:     InputParamError,
			ErrorDetail: err.Error(),
		})
		return
	}
	service, err := NewService(customerId, "", "")
	if err != nil {
		s.ReturnError(&Response{
			Code:        http.StatusBadRequest,
			Message:     GetClientFailed,
			ErrorDetail: err.Error(),
		})
		return
	}
	err = service.Update(student, chainCodeId)
	if err != nil {
		s.ReturnError(&Response{
			Code:        http.StatusBadRequest,
			Message:     UpdateStudentFromChainError,
			ErrorDetail: err.Error(),
		})
		return
	}
	s.ReturnOK(nil)
}

// 学生信息列表查询
func (s *StuController) QueryStudentInfoList() {
	customerId := s.Ctx.Input.Header("Customer-ID")
	chainCodeId := beego.AppConfig.String("defaultChaincodeID")
	var queryStudent = new(QueryStudentList)
	err := json.Unmarshal(s.Ctx.Input.RequestBody, queryStudent)
	if err != nil {
		logs.Error("Failed to Unmarshal request. err: %v\", err")
		s.ReturnError(&Response{
			Code:        http.StatusBadRequest,
			Message:     InputParamError,
			ErrorDetail: err.Error(),
		})
		return
	}
	service, err := NewService(customerId, "", "")
	if err != nil {
		s.ReturnError(&Response{
			Code:        http.StatusBadRequest,
			Message:     GetClientFailed,
			ErrorDetail: err.Error(),
		})
		return
	}
	respMember, err := service.QueryList(queryStudent, chainCodeId)
	if err != nil {
		logs.Error("[SDK_ERROR]Failed to query member list on the chain. err: %v", err)
		s.ReturnError(&Response{
			Code:        http.StatusInternalServerError,
			Message:     QueryStudentFromChainError,
			ErrorDetail: err.Error(),
		})
		return
	}
	logs.Info("[SDK_INFO]Successfully QueryListCustomer to the blockchain.")
	s.ReturnOK(respMember)
}

// 学生信息详情查询
func (s *StuController) QueryStudentInfo() {
	customerId := s.Ctx.Input.Header("Customer-ID")
	chainCodeId := beego.AppConfig.String("defaultChaincodeID")
	acctId := s.Ctx.Input.Param(":AcctId")
	service, err := NewService(customerId, "", "")
	if err != nil {
		s.ReturnError(&Response{
			Code:        http.StatusBadRequest,
			Message:     GetClientFailed,
			ErrorDetail: err.Error(),
		})
		return
	}
	logs.Info("[SDK_INFO]Start QueryInfoCustomer to the blockchain.", acctId, chainCodeId, customerId)
	respMember, err := service.QueryInfo(&QueryStudentInfo{AcctId: acctId}, chainCodeId)
	if err != nil {
		logs.Error("[SDK_ERROR]Failed to query member list on the chain. err: %v", err)
		s.ReturnError(&Response{
			Code:        http.StatusInternalServerError,
			Message:     QueryStudentFromChainError,
			ErrorDetail: err.Error(),
		})
		return
	}
	logs.Info("[SDK_INFO]Successfully QueryInfoCustomer to the blockchain.")
	s.ReturnOK(respMember)
}

// 学生信息删除
func (s *StuController) DeleteStudentInfo() {
	customerId := s.Ctx.Input.Header("Customer-ID")
	chainCodeId := beego.AppConfig.String("defaultChaincodeID")
	var deleteStudent = new(DeleteStudentInfo)
	err := json.Unmarshal(s.Ctx.Input.RequestBody, deleteStudent)
	if err != nil {
		logs.Error("Failed to Unmarshal request. err: %v\", err")
		s.ReturnError(&Response{
			Code:        http.StatusBadRequest,
			Message:     InputParamError,
			ErrorDetail: err.Error(),
		})
		return
	}
	service, err := NewService(customerId, "", "")
	if err != nil {
		s.ReturnError(&Response{
			Code:        http.StatusBadRequest,
			Message:     GetClientFailed,
			ErrorDetail: err.Error(),
		})
		return
	}
	err = service.DeleteInfo(deleteStudent, chainCodeId)
	if err != nil {
		logs.Error("[SDK_ERROR]Failed to query member list on the chain. err: %v", err)
		s.ReturnError(&Response{
			Code:        http.StatusInternalServerError,
			Message:     DeleteStudentFromChainError,
			ErrorDetail: err.Error(),
		})
		return
	}
	s.ReturnOK(nil)
}
