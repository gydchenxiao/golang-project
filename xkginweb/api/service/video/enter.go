package video

// 继承聚合---方便统一管理和调用
type ServiceGroup struct {
	VideoCategoryService
	VideoService
}
