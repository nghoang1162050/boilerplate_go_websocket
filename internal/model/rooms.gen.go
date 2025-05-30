// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameRoom = "rooms"

// Room mapped from table <rooms>
type Room struct {
	ID          string    `gorm:"column:id;primaryKey" json:"id"`
	Name        string    `gorm:"column:name;not null" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	HostID      string    `gorm:"column:host_id;not null" json:"host_id"`
	IsClosed    []uint8   `gorm:"column:is_closed;default:b'0'" json:"is_closed"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	Host        User      `gorm:"foreignKey:id;references:host_id" json:"host"`
}

// TableName Room's table name
func (*Room) TableName() string {
	return TableNameRoom
}
