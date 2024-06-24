package request

// 接受参数
type XkVideoReq struct {
	PageNum     int    `form:"pageNum" json:"pageNum"`
	PageSize    int    `form:"pageSize" json:"pageSize"`
	CategoryId  int    `form:"categoryId" json:"categoryId"`
	CategoryCid int    `form:"categoryCid" json:"categoryCid"`
	Keyword     string `form:"keyword" json:"keyword"`
	StartTime   string `form:"startTime" json:"startTime"`
	EndTime     string `form:"endTime" json:"endTime"`
	Status      int8   `form:"status" json:"status"`
}

// 接受参数
type XkVideoSaveReq struct {
	Id            int64  `gorm:"column:id;not null;comment:ID;primary_key" json:"id"`
	Title         string `gorm:"column:title;size:100;not null;default:'';comment:主题标题" json:"title"`                                     // 主题标题
	Content       string `gorm:"column:content;not null;comment:主题内容;type:longtext;" json:"content"`                                      // 主题内容
	Tags          string `gorm:"column:tags;size:40;not null;default:'';comment:主题标签" json:"tags"`                                        // 主题标签
	Description   string `gorm:"column:description;size:200;not null;default:'';comment:主题缩略描述" json:"description"`                       // 主题缩略描述
	Categoryid    int64  `gorm:"column:categoryid;default:0;not null;comment:主题分类ID" json:"categoryid"`                                   // 主题分类ID
	Gotop         int8   `gorm:"column:gotop;;not null;default:0;comment: 0普通 1置顶;default:0" json:"gotop"`                                // 0普通 1置顶
	Views         int32  `gorm:"column:views;;not null;default:0;comment:主题浏览次数;default:0" json:"views"`                                  // 主题浏览次数
	Collects      int32  `gorm:"column:collects;not null;default:0;comment:收藏数" json:"collects"`                                          // 收藏数
	Comments      int32  `gorm:"column:comments;not null;default:0;comment:评论数量" json:"comments"`                                         // 评论数量
	Htmlcontent   string `gorm:"column:htmlcontent;;not null;comment:html内容';type:longtext" json:"htmlcontent"`                           // html内容
	Categorytitle string `gorm:"column:categorytitle;size:40;not null;default:'';comment:分类标题" json:"categorytitle"`                      // 分类标题
	Img           string `gorm:"column:img;not null;size:200;default:'';comment:封面图" json:"img"`                                          // 封面图
	Coursetimer   string `gorm:"column:coursetimer;size:20;not null;default:'';comment:课程时长" json:"coursetimer"`                          // 课程时长
	Price         string `gorm:"column:price;not null;size:40;default:'0';comment:原始价格2499" json:"price"`                                 // 原始价格2499
	Realprice     string `gorm:"column:realprice;not null;size:40;default:'0';comment:真实价格1499" json:"realprice"`                         // 真实价格1499
	Coursetype    uint8  `gorm:"column:coursetype;not null;size:10;default:0;comment:课程类型 1 基础课  2 进阶课  4 面试课  3 实战课程" json:"coursetype"` // 课程类型 1 基础课  2 进阶课  4 面试课  3 实战课程
	Sorted        uint8  `gorm:"column:sorted;not null;default:1;comment:排序" json:"sorted"`                                               // 排序
	Beginer       string `gorm:"column:beginer;not null;size:10;default:'初级';comment:难易程度 初级，中级，高级" json:"beginer"`
	Vip           uint8  `gorm:"column:vip;not null;default:0;comment:VIP访问" json:"vip"`
	Comment       uint8  `gorm:"column:comment;not null;default:1;comment:0不可以评论 1可以评论" json:"comment"`  // 0不可以评论 1可以评论
	Isdelete      uint8  `gorm:"column:isdelete;not null;default:0;comment:删除状态0未删除1删除" json:"isdelete"` // 删除状态0未删除1删除
	Status        uint8  `gorm:"column:status;not null;default:1;comment:发布状态 1发布 0未发布" json:"status"`
	Isnew         uint8  `gorm:"column:isnew;not null;default:0;comment:是否最新 1是 0否" json:"isnew"`
	Ishot         uint8  `gorm:"column:ishot;not null;default:0;comment:是否最热 1是 0否"json:"ishot"`
	Ispush        uint8  `gorm:"column:ispush;not null;default:0;comment:是否推荐 1是0 否" json:"ispush"`
}
