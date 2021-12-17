package internal

import (
	"os"
	"path"
	"time"

	"github.com/flipped-aurora/gf-vue-admin/library/global"
	logs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

// GetWriteSyncer zap logger中加入file-rotatelogs
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := logs.New(
		path.Join(global.Config.Zap.Director, "%Y-%m-%d.log"),
		logs.WithMaxAge(7*24*time.Hour),
		logs.WithRotationTime(24*time.Hour),
	)
	if global.Config.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
