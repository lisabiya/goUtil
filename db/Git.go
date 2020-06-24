package db

type Model struct {
	ID int `gorm:"primary_key" form:"id" json:"id"`
}

type Git struct {
	Model
	GitUrl      string `json:"giturl" gorm:"not null;unique"`
	Name        string `json:"name"`
	Nickname    string `json:"nickname"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Dept        string `json:"dept"`
}

func (Git) TableName() string {
	return "t_git"
}

/**
 * 根据 @entity 新建数据
 */
func (entity *Git) Create() (err error) {
	err = GetDB().Table("t_git").Create(entity).Error
	return err
}
