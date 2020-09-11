package models

import (

	"time"
)

type Notification struct {
	ID        int64        `gorm:"column:id;primary_key" json:"id"`
	Link      string     `gorm:"column:link" json:"link"`
	Sender    uint     `gorm:"column:sender" json:"sender"`
	Title     string     `gorm:"column:title" json:"title"`
	Body      string     `gorm:"column:body" json:"body"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

//  TableName sets the insert table name for this struct type
func (n *Notification) TableName() string {
	return "notifications"
}

type NotificationStatus struct {
	ID         int64        `gorm:"column:id;primary_key" json:"id"`
	NotifID    int64        `gorm:"column:notif_id" json:"notif_id"`
	ReceiverID int64        `gorm:"column:receiver_id" json:"receiver_id"`
	Seen       int        `gorm:"column:seen" json:"seen"`
	SeenTime   time.Time `gorm:"column:seen_time" json:"seen_time"`
}
type UserNotifactions struct {
	ID        int64        `gorm:"column:id" json:"id"`
	Link      string     `gorm:"column:link" json:"link"`
	Sender    string     `gorm:"column:sender" json:"sender"`
	Title     string     `gorm:"column:title" json:"title"`
	Body      string     `gorm:"column:body" json:"body"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	SID         int64        `gorm:"column:sid" json:"sid"`
	ReceiverID int64        `gorm:"column:receiver_id" json:"receiver_id"`
	Seen       int        `gorm:"column:seen" json:"seen"`
	SeenTime   *time.Time `gorm:"column:seen_time" json:"seen_time"`
	FirstName      string     `gorm:"column:first_name" json:"first_name"`
	Login      string     `gorm:"column:login" json:"login"`
}
//  TableName sets the insert table name for this struct type
func (n *NotificationStatus) TableName() string {
	return "notification_status"
}

type Payload struct {
	RegistrationIds []string `json:"registration_ids"`
	Data            FCMData     `json:"data"`
	Notification            FCMNotification     `json:"notification"`
}
type FCMData struct {
	Title      string `json:"title"`
	Body      string     `gorm:"column:body" json:"body"`
	Sound      string `json:"sound"`
	Icon  string `json:"icon"`
	ClickAction      string     `gorm:"column:click_action" json:"click_action"`
	Link      string     `gorm:"column:link" json:"link"`
	FirstName      string     `gorm:"column:first_name" json:"first_name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	ID        int64        `gorm:"column:id" json:"id"`
}
type FCMNotification struct {
	Title      string `json:"title"`
	Body      string     `gorm:"column:body" json:"body"`
	ClickAction      string     `gorm:"column:click_action" json:"click_action"`
	Icon  string `json:"icon"`
}

type NotificationData struct {
	Users []int
	Roles []int
	Data            FCMData     `json:"data"`
	Notification            FCMNotification     `json:"notification"`
}
type NotificationTarget struct {
	ID        int64        `gorm:"column:id;primary_key" json:"id"`
	Link      string     `gorm:"column:link" json:"link"`
	Title     string     `gorm:"column:title" json:"title"`
	Body      string     `gorm:"column:body" json:"body"`
	SchemaId      int     `gorm:"column:schema_id" json:"schema_id"`
	TargetRole      int     `gorm:"column:target_role" json:"target_role"`
	Condition      string     `gorm:"column:condition" json:"condition"`
	TargetActions      string     `gorm:"column:target_actions" json:"target_actions"`
}

//  TableName sets the insert table name for this struct type
func (n *NotificationTarget) TableName() string {
	return "notification_targets"
}