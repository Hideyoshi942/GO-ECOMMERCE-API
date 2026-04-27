package models

const TableNameGoCrmUser = "go_crm_user"

// GoCrmUser Account
type GoCrmUser struct {
	UsrID int32 `gorm:"column:usr_id;primaryKey;autoIncrement;comment:Account ID" json:"usr_id"`

	UsrEmail    string `gorm:"column:usr_email;type:varchar(255);uniqueIndex;not null;comment:Email" json:"usr_email"`
	UsrPhone    string `gorm:"column:usr_phone;type:varchar(20);uniqueIndex;not null;comment:Phone" json:"usr_phone"`
	UsrUsername string `gorm:"column:usr_username;type:varchar(100);uniqueIndex;not null;comment:Username" json:"usr_username"`

	UsrPassword string `gorm:"column:usr_password;type:varchar(255);not null;comment:Hashed Password" json:"-"`

	UsrCreatedAt int64 `gorm:"column:usr_created_at;autoCreateTime;comment:Created Time" json:"usr_created_at"`
	UsrUpdatedAt int64 `gorm:"column:usr_updated_at;autoUpdateTime;comment:Updated Time" json:"usr_updated_at"`

	UsrCreateIPAt    string `gorm:"column:usr_create_ip_at;type:varchar(45);comment:Created IP" json:"usr_create_ip_at"`
	UsrLastLoginAt   int64  `gorm:"column:usr_last_login_at;comment:Last Login Time" json:"usr_last_login_at"`
	UsrLastLoginIPAt string `gorm:"column:usr_last_login_ip_at;type:varchar(45);comment:Last Login IP" json:"usr_last_login_ip_at"`

	UsrLoginTimes int32 `gorm:"column:usr_login_times;default:0;comment:Login Times" json:"usr_login_times"`

	UsrStatus int8 `gorm:"column:usr_status;default:1;comment:1:active,0:inactive" json:"usr_status"`
}
