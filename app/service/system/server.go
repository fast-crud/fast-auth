package system

import (
	"runtime"
	"time"

	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

var Server = new(_server)

type _server struct{}

type ServerRes struct {
	Os   Os   `json:"os"`
	Cpu  Cpu  `json:"cpu"`
	Rrm  Rrm  `json:"ram"`
	Disk Disk `json:"disk"`
}

type Os struct {
	GOOS         string `json:"goos"`
	NumCPU       int    `json:"numCpu"`
	Compiler     string `json:"compiler"`
	GoVersion    string `json:"goVersion"`
	NumGoroutine int    `json:"numGoroutine"`
}

type Cpu struct {
	Cpus  []float64 `json:"cpus"`
	Cores int       `json:"cores"`
}

type Rrm struct {
	UsedMB      int `json:"usedMb"`
	TotalMB     int `json:"totalMb"`
	UsedPercent int `json:"usedPercent"`
}

type Disk struct {
	UsedMB      int `json:"usedMb"`
	UsedGB      int `json:"usedGb"`
	TotalMB     int `json:"totalMb"`
	TotalGB     int `json:"totalGb"`
	UsedPercent int `json:"usedPercent"`
}

// GetServerInfo 获取服务器信息
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_server) GetServerInfo() (data *ServerRes, err error) {
	server := ServerRes{Os: s.InitOs()}
	if server.Cpu, err = s.InitCpu(); err != nil {
		return nil, errors.Wrap(err, "获取CPU信息失败!")
	}
	if server.Rrm, err = s.InitRAM(); err != nil {
		return nil, errors.Wrap(err, "获取ARM信息失败!")
	}
	if server.Disk, err = s.InitDisk(); err != nil {
		return nil, errors.Wrap(err, "获取硬盘信息失败!")
	}

	return &server, nil
}

// InitOs 获取系统信息 组装数据为 response.Os
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_server) InitOs() Os {
	return Os{
		GOOS:         runtime.GOOS,
		NumCPU:       runtime.NumCPU(),
		Compiler:     runtime.Compiler,
		GoVersion:    runtime.Version(),
		NumGoroutine: runtime.NumGoroutine(),
	}
}

// InitCpu 获取CPU信息 组装数据为 response.Cpu
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_server) InitCpu() (Cpu, error) {
	var _cpu Cpu
	cores, err := cpu.Counts(false)
	if err != nil {
		return _cpu, err
	}
	_cpu.Cores = cores
	_cpu.Cpus, err = cpu.Percent(time.Duration(200)*time.Millisecond, true)
	if err != nil {
		return _cpu, err
	}
	return _cpu, nil
}

// InitRAM 获取ARM信息 组装数据为 Rrm
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_server) InitRAM() (Rrm, error) {
	var arm Rrm
	virtualMemoryStat, err := mem.VirtualMemory()
	if err != nil {
		return arm, err
	}
	arm.UsedMB = int(virtualMemoryStat.Used) / MB
	arm.TotalMB = int(virtualMemoryStat.Total) / MB
	arm.UsedPercent = int(virtualMemoryStat.UsedPercent)
	return arm, nil
}

// InitDisk 获取硬盘信息 组装数据为 Disk
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_server) InitDisk() (Disk, error) {
	var _disk Disk
	usageStat, err := disk.Usage("/")
	if err != nil {
		return _disk, err
	}
	_disk.UsedMB = int(usageStat.Used) / MB
	_disk.UsedGB = int(usageStat.Used) / GB
	_disk.TotalMB = int(usageStat.Total) / MB
	_disk.TotalGB = int(usageStat.Total) / GB
	_disk.UsedPercent = int(usageStat.UsedPercent)
	return _disk, nil
}
