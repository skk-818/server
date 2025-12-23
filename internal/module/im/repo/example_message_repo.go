package repo

import (
	"context"
	"server/internal/core/mysql"

	"gorm.io/gorm"
)

// 示例：IM 消息仓储层
type messageRepo struct {
	db *gorm.DB
}

// NewMessageRepo 使用 ImDB 创建消息仓储
func NewMessageRepo(imDB *mysql.ImDB) *messageRepo {
	return &messageRepo{
		db: imDB.DB,
	}
}

// 示例方法：创建消息
func (r *messageRepo) CreateMessage(ctx context.Context, message interface{}) error {
	return r.db.WithContext(ctx).Create(message).Error
}

// 示例方法：查询消息
func (r *messageRepo) FindMessage(ctx context.Context, id int64) (interface{}, error) {
	var message interface{}
	err := r.db.WithContext(ctx).First(&message, id).Error
	return message, err
}
