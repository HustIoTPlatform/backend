// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDeviceModelEvent = "device_model_events"

// DeviceModelEvent mapped from table <device_model_events>
type DeviceModelEvent struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`                                   // ID
	DeviceTemplateID string    `gorm:"column:device_template_id;not null" json:"device_template_id"` // Device template ID
	DataName         *string   `gorm:"column:data_name" json:"data_name"`                             // Data name
	DataIdentifier   string    `gorm:"column:data_identifier;not null" json:"data_identifier"`  // Data identifier
	Param            *string   `gorm:"column:params" json:"params"`                                  // Parameters
	Description      *string   `gorm:"column:description" json:"description"`                     // Description
	AdditionalInfo   *string   `gorm:"column:additional_info" json:"additional_info"`         // Additional info
	CreatedAt        time.Time `gorm:"column:created_at;not null" json:"created_at"`            // Creation time
	UpdatedAt        time.Time `gorm:"column:updated_at;not null" json:"updated_at"`              // Update time
	Remark           *string   `gorm:"column:remark" json:"remark"`                                    // Remark
	TenantID         string    `gorm:"column:tenant_id;not null" json:"tenant_id"`                                   // Tenant ID
}

// TableName DeviceModelEvent's table name
func (*DeviceModelEvent) TableName() string {
	return TableNameDeviceModelEvent
}
