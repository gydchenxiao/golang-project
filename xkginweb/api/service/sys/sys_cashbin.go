package sys

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"strconv"
	"strings"
	"sync"
	"xkginweb/global"

	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// Casbin info structure
type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateCasbin
//@description: 更新casbin权限
//@param: authorityId string, casbinInfos []request.CasbinInfo
//@return: error

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService)

func (casbinService *CasbinService) UpdateCasbin(AuthorityID uint, casbinInfos []CasbinInfo) error {
	authorityId := strconv.Itoa(int(AuthorityID))
	casbinService.ClearCasbin(0, authorityId)
	rules := [][]string{}
	for _, v := range casbinInfos {
		rules = append(rules, []string{authorityId, v.Path, v.Method})
	}
	e := casbinService.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	err := e.InvalidateCache()
	if err != nil {
		return err
	}
	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateCasbinApi
//@description: API更新随动
//@param: oldPath string, newPath string, oldMethod string, newMethod string
//@return: error

func (casbinService *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.KSD_DB.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	e := casbinService.Casbin()
	err = e.InvalidateCache()
	if err != nil {
		return err
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPolicyPathByAuthorityId
//@description: 获取权限列表
//@param: authorityId string
//@return: pathMaps []request.CasbinInfo

func (casbinService *CasbinService) GetPolicyPathByAuthorityId(AuthorityID uint) (pathMaps []CasbinInfo) {
	e := casbinService.Casbin()
	authorityId := strconv.Itoa(int(AuthorityID))
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ClearCasbin
//@description: 清除匹配的权限
//@param: v int, p ...string
//@return: bool

func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Casbin
//@description: 持久化到数据库  引入自定义规则
//@return: *casbin.Enforcer

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

func (casbinService *CasbinService) CasbinFile() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(global.KSD_DB)
		if err != nil {
			global.Log.Error("适配数据库失败请检查casbin表是否为InnoDB引擎!", zap.Error(err))
			return
		}
		m, err := model.NewModelFromFile("./conf/rbac_models.conf")
		if err != nil {
			global.Log.Error("字符串加载模型失败!", zap.Error(err))
			return
		}
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, a)
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		_ = syncedCachedEnforcer.LoadPolicy()

		// 设置用户root的角色为superadmin
		syncedCachedEnforcer.AddRoleForUser("root", "superadmin")
		// 添加自定义函数
		syncedCachedEnforcer.AddFunction("checkSuperAdmin", func(arguments ...interface{}) (interface{}, error) {
			// 获取用户名
			username := arguments[0].(string)
			// 检查用户名的角色是否为superadmin
			return syncedCachedEnforcer.HasRoleForUser(username, "superadmin")
		})
		// 添加自定义函数
		// equals(r.sub, p.sub)
		syncedCachedEnforcer.AddFunction("equals", func(arguments ...interface{}) (interface{}, error) {
			// 获取用户名
			args1 := arguments[0].(string)
			args2 := arguments[1].(string)
			// 检查用户名的角色是否为superadmin
			return strings.EqualFold(args1, args2), nil
		})

	})
	return syncedCachedEnforcer
}

func (casbinService *CasbinService) Casbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(global.KSD_DB)
		if err != nil {
			global.Log.Error("适配数据库失败请检查casbin表是否为InnoDB引擎!", zap.Error(err))
			return
		}

		// 比较规则
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			global.Log.Error("字符串加载模型失败!", zap.Error(err))
			return
		}
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, a)
		syncedCachedEnforcer.SetExpireTime(60 * 60) // 设置超时时间
		_ = syncedCachedEnforcer.LoadPolicy()       // 加载 cashbin_rule 数据库表的数据到内存
	})
	return syncedCachedEnforcer
}
