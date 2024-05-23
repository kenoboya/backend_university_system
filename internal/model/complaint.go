package model

import "time"

type Complaint struct {
	ComplaintID     uint64    `db:"complaint_id" json:"complaint_id"`
	ReportingUserID uint64    `db:"reporting_user_id" json:"reporting_user_id"`
	ReportedUserID  uint64    `db:"reported_user_id" json:"reported_user_id"`
	Cause           string    `db:"cause" json:"cause"`
	Time            time.Time `db:"time" json:"time"`
	Response        *string   `db:"response" json:"response"`
}
type CreateComplaintInput struct {
	ReportingUserID uint64 `db:"reporting_user_id" json:"reporting_user_id"`
	ReportedUserID  uint64 `db:"reported_user_id" json:"reported_user_id"`
	Cause           string `db:"cause" json:"cause"`
}
type ResponseComplaintInput struct {
	ComplaintID uint64 `db:"complaint_id" json:"complaint_id"`
	Response    string `db:"response" json:"response"`
	AdminID     uint64 `db:"admin_id" json:"admin_id"`
	UserID      uint64 `db:"user_id" json:"user_id"`
	Blocked     bool   `json:"blocked"`
}
