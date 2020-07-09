package db

import (
	"strconv"
	"time"
)

/*********************审批单原始信息解析*****************************/
type BpmsProcessResponSe struct {
	ErrCode         int                 `json:"errcode"`
	ProcessInstance BpmsProcessInstance `json:"process_instance"`
	RequestId       string              `json:"request_id"`
}

type BpmsProcessInstance struct {
	AttachedProcessInstanceIds []string              `json:"attached_process_instance_ids"`
	BizAction                  string                `json:"biz_action"`
	BusinessId                 string                `json:"business_id"`
	CreateTime                 string                `json:"create_time"`
	FinishTime                 string                `json:"finish_time"`
	ComponentValues            []BpmsComponentValues `json:"form_component_values"`
	OperationRecords           []interface{}         `json:"operation_records"`
	OriginatorDeptId           string                `json:"originator_dept_id"`
	OriginatorDeptName         string                `json:"originator_dept_name"`
	OriginatorUserId           string                `json:"originator_userid"`
	Result                     string                `json:"result"`
	Status                     string                `json:"status"`
	Tasks                      []interface{}         `json:"tasks"`
	Title                      string                `json:"title"`
}

type BpmsComponentValues struct {
	ComponentType string `json:"component_type"`
	ExtValue      string `json:"ext_value"`
	Id            string `json:"id"`
	Name          string `json:"name"`
	Value         string `json:"value"`
}

type BpmsComponentValueTable []struct {
	ComponentName string `json:"componentName"`
	ComponentType string `json:"componentType"`
	Props         struct {
		BizAlias       string        `json:"bizAlias"`
		HolidayOptions []interface{} `json:"holidayOptions"`
		ID             string        `json:"id"`
		Label          string        `json:"label"`
		Placeholder    string        `json:"placeholder"`
		Required       bool          `json:"required"`
	} `json:"props,omitempty"`
	Value    string `json:"value"`
	ExtValue string `json:"extValue,omitempty"`
}

type Employee struct {
	EmployeeName string `json:"employee_name"`
	Department   string `json:"department"`
	Position     string `json:"position"`
	MobileNumber string `json:"mobile_number"`
	EmployeeType string `json:"employee_type"`
	EntryDate    string `json:"entry_date"`
}

type RowValue struct {
	Detail []RowValueDetail `json:"rowValue"`
}

type RowValueDetail struct {
	Label       string           `json:"label"`
	ExtendValue []RowExtendValue `json:"extendValue"`
	Value       string           `json:"value"`
}

type RowExtendValue struct {
	EmplId string `json:"emplId"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type ProjectValue struct {
	Detail []ProjectRowValueDetail `json:"rowValue"`
}

type ProjectRowValueDetail struct {
	Label string   `json:"label"`
	Value []string `json:"value"`
}

/*********************审批单基础类型*****************************/
type TransportIml interface {
	Input(name string, value BpmsComponentValues)
	TransFromApproval(processInstanceId string, instance BpmsProcessInstance)
}

type BaseBpms struct {
	Model
	//基本头
	ProcessInstanceId  string `json:"process_instance_id" gorm:"COMMENT:'审批id';unique_index:ProcessInstanceId;"`
	OriginatorDeptId   string `json:"originator_dept_id" gorm:"COMMENT:'申请部门Id'"`
	OriginatorDeptName string `json:"originator_dept_name" gorm:"COMMENT:'申请部门'"`
	OriginatorUserid   string `json:"originator_userid" gorm:"COMMENT:'申请用户userid'"`
	Result             string `json:"result" gorm:"COMMENT:'结果'"`
	Status             string `json:"status" gorm:"COMMENT:'状态'"`
	Title              string `json:"title" gorm:"COMMENT:'标题'"`
	//创建时间
	CreateTime      string `json:"create_time" gorm:"COMMENT:'创建时间'"`
	FinishTime      string `json:"finish_time" gorm:"COMMENT:'结束时间'"`
	CreateTimestamp int64  `json:"create_timestamp" gorm:"COMMENT:'创建时间时间戳'"`
	FinishTimestamp int64  `json:"finish_timestamp" gorm:"COMMENT:'结束时间时间戳'"`
}

//从审批单中解析出基础数据
func (base *BaseBpms) TransBaseFromApproval(processInstanceId string, instance BpmsProcessInstance) {
	base.ProcessInstanceId = processInstanceId
	base.OriginatorDeptId = instance.OriginatorDeptId
	base.OriginatorDeptName = instance.OriginatorDeptName
	base.OriginatorUserid = instance.OriginatorUserId
	base.Result = instance.Result
	base.Status = instance.Status
	base.Title = instance.Title
	base.CreateTime = instance.CreateTime
	base.FinishTime = instance.FinishTime
	base.CreateTimestamp = DatetimeTransToTimestamp(instance.CreateTime, "2006-01-02 15:04:05")
	base.FinishTimestamp = DatetimeTransToTimestamp(instance.FinishTime, "2006-01-02 15:04:05")
}

func TranFormDataByName(forms []BpmsComponentValues, inteface TransportIml) {
	for _, val := range forms {
		inteface.Input(val.Name, val)
	}
	return
}

func DatetimeTransToTimestamp(datetime, timeLayout string) int64 {
	if datetime == "" {
		return -1
	}
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	return tmp.Unix()
}

func parseToInt64(val string) int64 {
	if val == "" || val == "null" {
		return 0
	}
	parseInt, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return parseInt
}

func parseToInt(val string) int {
	if val == "" || val == "null" {
		return 0
	}
	r, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return r
}

func parseToFloat64(val string) float64 {
	if val == "" || val == "null" {
		return 0
	}
	r, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0
	}
	return r
}
