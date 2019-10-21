package model

import "time"

//noinspection SpellCheckingInspection
type Record struct {
	ID            int64     `gorm:"id"`
	SharingKey    int64     `gorm:"sharing_key"`
	PtActivityId  string    `gorm:"pt_activity_id"`
	ShareMemberNo string    `gorm:"share_member_no"`
	Phone         string    `gorm:"phone"`
	Frequency     int32     `gorm:"frequency"`
	CreateTime    time.Time `gorm:"create_time"`
}

func (Record)TableName() string {
	return "pt_poster_scan_record"
}
