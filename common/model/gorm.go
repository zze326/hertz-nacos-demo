package model

import (
	"time"

	"gorm.io/gorm"
)

/**
 * @Author: zze
 * @Date: 2022/5/30 19:20
 * @Desc: gorm 基础模型
 */

// Model 由于 go swagger 不支持默认 gorm.Model 主键是 uint 类型，所以需要自定义
type Model struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
