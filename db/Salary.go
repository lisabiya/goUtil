package db

type Salary struct {
	Model
	Name           string `json:"name" gorm:"not null;unique"` //姓名
	Department     string `json:"department"`                  //部门
	SocialSecurity string `json:"social_security"`             //社保
	ProvidentFund  string `json:"provident_fund"`              //公积金
	Salary         string `json:"salary"`                      //实发工资
	SalaryTime     string `json:"salary_time"`                 //工资日期
}

func (Salary) TableName() string {
	return "t_salary"
}

/**
 * 根据 @entity 新建数据
 */
func (entity *Salary) Create() (err error) {
	err = GetDB().Table("t_salary").Create(entity).Error
	return err
}
