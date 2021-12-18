package system

import (
	"fmt"
	"strings"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/fast-crud/fast-auth/library/global"
)

var (
	once           sync.Once
	Casbin         = new(_casbin)
	syncedEnforcer *casbin.SyncedEnforcer
)

type _casbin struct{}

// Casbin 持久化到数据库  引入自定义规则
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_casbin) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := adapter.NewAdapterByDB(global.Db)
		fmt.Println(global.Config.Casbin.ModelPath)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(global.Config.Casbin.ModelPath, a)
		syncedEnforcer.AddFunction("ParamsMatch", s.ParamsMatchFunc)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}

// Clear 清除匹配的权限
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_casbin) Clear(v int, p ...string) bool {
	e := s.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

// ParamsMatch 自定义规则函数
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_casbin) ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0] // 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

// ParamsMatchFunc 自定义规则函数
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_casbin) ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	return s.ParamsMatch(name1, name2), nil
}
