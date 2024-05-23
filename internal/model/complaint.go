package model

import "time"

type Complaint struct {
	ComplaintID     int64     `db:"complaint_id" json:"complaint_id"`
	ReportingUserID int64     `db:"reporting_user_id" json:"reporting_user_id"`
	ReportedUserID  int64     `db:"reported_user_id" json:"reported_user_id"`
	Cause           string    `db:"cause" json:"cause"`
	Time            time.Time `db:"time" json:"time"`
	Response        *string   `db:"response" json:"response"`
}
type CreateComplaintInput struct {
	ReportingUserID int64  `db:"reporting_user_id" json:"reporting_user_id"`
	ReportedUserID  int64  `db:"reported_user_id" json:"reported_user_id"`
	Cause           string `db:"cause" json:"cause"`
}
type ResponseComplaintInput struct {
	ComplaintID int64  `db:"complaint_id" json:"complaint_id"`
	Response    string `db:"response" json:"response"`
	AdminID     int64  `db:"admin_id" json:"admin_id"`
	UserID      int64  `db:"user_id" json:"user_id"`
	Blocked     bool   `json:"blocked"`
}
