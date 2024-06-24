package video

import "xkginweb/global"

type XkVideo struct {
	global.GVA_MODEL
	Title         string `validate:"required|minLen:7" message:"标题长度必须大于等于7位" gorm:"column:title;size:100;not null;default:'';comment:主题标题" json:"title"` // 主题标题
	Description   string `gorm:"column:description;size:200;not null;default:'';comment:主题缩略描述" json:"description"`
	Img           string `gorm:"column:img;not null;size:200;default:'';comment:封面图" json:"img"`
	Content       string `gorm:"column:content;not null;comment:主题内容;type:longtext;" json:"content"`                                      // 主题内容
	Tags          string `gorm:"column:tags;size:40;not null;default:'';comment:主题标签" json:"tags"`                                        // 主题标签
	Views         int32  `gorm:"column:views;;not null;default:0;comment:主题浏览次数;default:0" json:"views"`                                  // 主题浏览次数
	Collects      int32  `gorm:"column:collects;not null;default:0;comment:收藏数" json:"collects"`                                          // 收藏数
	Comments      int32  `gorm:"column:comments;not null;default:0;comment:评论数量" json:"comments"`                                         // 评论数量
	CategoryPid   string `gorm:"column:category_pid;default:0;not null;comment:主题分类ID" json:"categoryPid"`                                // 主题分类ID
	CategoryPname string `gorm:"column:category_pname;size:40;not null;default:'';comment:分类标题" json:"categoryPname"`                     // 分类标题
	CategoryCid   string `gorm:"column:category_cid;default:0;not null;comment:子分类ID" json:"categoryCid"`                                 // 主题分类ID
	CategoryCname string `gorm:"column:category_cname;size:40;not null;default:'';comment:子分类标题" json:"categoryCname"`                    // 分类标题
	Avatar        string `gorm:"column:avatar;not null;size:200;default:'';comment:作者头像" json:"avatar"`                                   // 作者头像
	Nickname      string `gorm:"column:nickname;not null;size:64;default:'';comment:昵称" json:"nickname"`                                  // 昵称
	Userid        int64  `gorm:"column:userid;not null;default:0;comment:用户" json:"userid"`                                               // 用户
	Coursetimer   string `gorm:"column:coursetimer;size:20;not null;default:'';comment:课程时长" json:"coursetimer"`                          // 课程时长
	Price         string `gorm:"column:price;not null;size:40;default:'0';comment:原始价格2499" json:"price"`                                 // 原始价格2499
	Realprice     string `gorm:"column:realprice;not null;size:40;default:'0';comment:真实价格1499" json:"realprice"`                         // 真实价格1499
	Coursetype    uint8  `gorm:"column:coursetype;not null;size:10;default:0;comment:课程类型 1 基础课  2 进阶课  4 面试课  3 实战课程" json:"coursetype"` // 课程类型 1 基础课  2 进阶课  4 面试课  3 实战课程
	Status        uint8  `gorm:"column:status;not null;default:1;comment:发布状态 1发布 0未发布" json:"status"`
	IsComment     uint8  `gorm:"column:is_comment;not null;default:1;comment:0不可以评论 1可以评论" json:"isComment"` // 0不可以评论 1可以评论
	IsNew         uint8  `gorm:"column:is_new;not null;default:0;comment:是否最新 1是 0否" json:"isNew"`
	IsHot         uint8  `gorm:"column:is_hot;not null;default:0;comment:是否最热 1是 0否"json:"isHot"`
	IsPush        uint8  `gorm:"column:is_push;not null;default:0;comment:是否推荐 1是0 否" json:"isPush"`
}

func (XkVideo) TableName() string {
	return "xk_video"
}
