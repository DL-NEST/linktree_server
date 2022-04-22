package DB

import (
	"linktree_server/utils/logger"
)

//// 是否需要迁移
//func needMigration() bool {
//	return DB.Where("name = ?", "db_version_"+conf.RequiredDBVersion).First(&setting).Error != nil
//}

func migration() {

	logger.Log().Info("开始初始化数据库")
	// 自动迁移(表的创建)
	err := DB.AutoMigrate(&User{})
	if err != nil {
		return
	}

}


