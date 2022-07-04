package db_entity

import (
	"time"

	// "github.com/jinzhu/gorm"
	"github.com/google/uuid"
)

// カラムは必ず大文字始まり
type Sensor struct {
	Sensor_id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"sensor_id"`
	Sensor_name string `json:"sensor_name"`
	User_id string  `json:"user_id"`
	Is_active bool `json:"is_active"`
	Created_at time.Time `json:"created_at" gorm:"primary_key"`
}