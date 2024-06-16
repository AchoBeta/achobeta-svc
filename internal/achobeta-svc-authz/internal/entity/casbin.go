package entity

// CasbinRule casbin rule
type CasbinRule struct {
	ID    uint   `json:"id" gorm:"primarykey"`
	PType string `json:"p_type" gorm:"column:ptype;size:64"`
	V0    string `json:"v0" gorm:"column:v0;size:64"`
	V1    string `json:"v1" gorm:"column:v1;size:64"`
	V2    string `json:"v2" gorm:"column:v2;size:64"`
	V3    string `json:"v3" gorm:"column:v3;size:64"`
	V4    string `json:"v4" gorm:"column:v4;size:64"`
	V5    string `json:"v5" gorm:"column:v5;size:64"`
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}
