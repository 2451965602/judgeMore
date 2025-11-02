package mysql

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserId    int64
	RoleId    string //实际上是我们业务过程中区分用户的主键
	UserName  string
	UserRole  string
	College   string
	Grade     string
	Major     string
	Email     string
	Status    int
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type Event struct {
	EventId        string `gorm:"primaryKey;autoIncrement:true;column:event_id"`
	UserId         string `gorm:"not null;column:user_id"`
	RecognizedId   string `gorm:"not null;column:recognized_id"`
	EventName      string `gorm:"size:200;not null;column:event_name"`
	EventOrganizer string `gorm:"size:200;not null;column:event_organizer"`
	EventLevel     string `gorm:"size:20;not null;column:event_level"`
	EventInfluence string `gorm:"size:10;not null;column:event_influence"`
	AwardLevel     string `gorm:"size:50;not null;column:award_level"`
	MaterialUrl    string `gorm:"size:500;not null;column:material_url"`
	MaterialStatus string `gorm:"size:20;not null;default:'待审核';column:material_status"`
	AutoExtracted  bool   `gorm:"not null;default:false;column:auto_extracted"`
	AwardAt        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
type RecognizedEvent struct {
	RecognizedEventId   string `gorm:"primaryKey;autoIncrement:true;column:recognized_event_id"`
	College             string `gorm:"size:255;not null;column:college"`
	RecognizedEventName string `gorm:"size:255;not null;column:recognized_event_name"`
	Organizer           string `gorm:"size:255;not null;column:organizer"`
	RecognizedEventTime string `gorm:"size:50;not null;column:recognized_event_time"`
	RelatedMajors       string `gorm:"size:255;column:related_majors"`
	ApplicableMajors    string `gorm:"size:255;column:applicable_majors"`
	RecognitionBasis    string `gorm:"size:255;column:recognition_basis"`
	RecognizedLevel     string `gorm:"size:50;not null;column:recognized_level"`
	IsActive            bool   `gorm:"default:true;column:is_active"`
	RuleId              int64  `gorm:"not null;column:rule_id"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}
