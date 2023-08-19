package service

type FavoriteService interface {
	// FavoriteAction 点赞操作
	FavoriteAction(userId int64, videoId int64, actionType int32) error

	// GetFavoriteList 获取当前用户点赞列表
	GetFavoriteList(userId int64) ([]Video, error)
	////获取视频列表
	//GetVideo(videoId []int64, likeCnt int64) ([]Video, error)

	// IsFavoritedByUser 当前用户是否点赞该视频
	IsFavoritedByUser(userId int64, videoId int64) (bool, error)
	// GetUserFavoriteCount 获取用户点赞数量
	GetUserFavoriteCount(userId int64) (int64, error)
	// GetVideoFavoritedCount 获取视频点赞数
	GetVideoFavoritedCount(videoId int64) (int64, error)

	// GetUserFavoritedCount 计算用户被点赞的视频获赞总数
	GetUserFavoritedCount(userId int64) (int64, error)
}
