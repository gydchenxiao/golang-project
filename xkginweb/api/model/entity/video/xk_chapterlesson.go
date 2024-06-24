package video

import "xkginweb/global"

// 视频章节
type XkVideoChapterLesson struct {
	global.GVA_MODEL
	Title           string `gorm:"size:100;not null;default:'';comment:章节标题" json:"title"` // 主题标题
	Description     string `gorm:"size:200;not null;default:'';comment:章节描述" json:"description"`
	Type            uint8  `gorm:"not null;default:1;comment:1视频 2源码 3 文件" json:"type"`
	VideoId         string `gorm:"size:60;not null;default:'';comment:播放视频ID" json:"description"`
	VideoInfo       string `gorm:"size:1000;not null;default:'';comment:视频的相关信息" json:"description"`
	TotalTimer      string `gorm:"size:20;not null;default:'';comment:总时长" json:"totalTimer"`
	Status          uint8  `gorm:"not null;default:1;comment:发布状态 1发布 0未发布" json:"status"`
	IsFree          uint8  `gorm:";not null;default:1;comment:是否最新 1不免费 0免费" json:"isFree"`
	CreatorUserId   int    `gorm:"not null;default:0;comment:创作者ID" json:"creatorUserId"`
	CreatorUserName int    `gorm:"not null;default:0;comment:创作者名字" json:"creatorUserName"`
	ParentId        int    `gorm:"not null;default:0;comment:0代表章 非0代表节" json:"parentId"`
}

func (XkVideoChapterLesson) TableName() string {
	return "xk_video_chapterlesson"
}
