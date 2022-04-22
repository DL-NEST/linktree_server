package service

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"strings"
	"time"
)

type SysInfo struct {
	CpuInfo  CpuInfo             `json:"cpuInfo"`
	MemInfo  MemInfo             `json:"memInfo"`
	DiskNode map[string]DiskNode `json:"diskNode"`
	HostInfo HostInfo            `json:"hostInfo"`
	NetNode  map[string]NetNode  `json:"netNode"`
}

type CpuInfo struct {
	// cpu厂商
	CpuVendor string `json:"CpuVendor"`
	// cpu型号
	CpuName string
	// 核心数
	CpuCore int
	// 逻辑线程
	CpuHt int
	// cpu频率
	CpuHz float64
}

type MemInfo struct {
	//	总内存
	Total int
	// 使用内存
	Used int
	// 剩余内存
	Free int
	// 使用占比
	UsedPercent float64
}

type DiskNode struct {
	//	总磁盘大小
	Total int
	// 使用磁盘大小
	Used int
	// 剩余磁盘大小
	Free int
	// 使用占比
	UsedPercent float64
}

type HostInfo struct {
	// 主机用户名称
	HostName string
	// BootTime
	BootTime string
	// 运行时间
	UpTime string
	// 操作系统
	OS string
	// 操作平台
	PlatForm string
	// 系统版本号
	PlatFormVersion string
	// 系统架构
	KernelArch string
	// 系统标识
	HostID string
}

type NetNode struct {
	// 字节发送量
	BytesSent int
	// 字节接收量
	BytesRect int
	// 包发送量
	PacketsSent int
	// 包接收量
	PacketsRect int
}

func GetCpuInfo() CpuInfo {
	core, _ := cpu.Info()
	coreInfo := core[0]
	cores, _ := cpu.Counts(false)
	Ht, _ := cpu.Counts(true)
	cpuInfo := CpuInfo{
		CpuVendor: coreInfo.VendorID,
		CpuName:   strings.Replace(coreInfo.ModelName, " ", "", -1),
		CpuCore:   cores,
		CpuHt:     Ht,
		CpuHz:     coreInfo.Mhz / 1000,
	}
	return cpuInfo
}

func GetMemInfo() MemInfo {
	ramInfo, _ := mem.VirtualMemory()
	memInfo := MemInfo{
		Total:       int(ramInfo.Total),
		Used:        int(ramInfo.Used),
		Free:        int(ramInfo.Free),
		UsedPercent: ramInfo.UsedPercent,
	}
	return memInfo
}

func GetDiskInfo() map[string]DiskNode {
	diskInfo := make(map[string]DiskNode)
	diskList, _ := disk.Partitions(true)
	for i := range diskList {
		diskIO, _ := disk.Usage(diskList[i].Device)
		diskInfo[diskList[i].Device] = DiskNode{
			Total:       int(diskIO.Total / 1024 / 1024 / 1024),
			Used:        int(diskIO.Used / 1024 / 1024 / 1024),
			Free:        int(diskIO.Free / 1024 / 1024 / 1024),
			UsedPercent: diskIO.UsedPercent,
		}
	}
	return diskInfo
}

func GetHostInfo() HostInfo {
	hostInfo, _ := host.Info()
	hosts := HostInfo{
		HostName:        hostInfo.Hostname,
		BootTime:        time.Unix(int64(hostInfo.BootTime), 0).Format("2006-01-02 15:04:05"),
		UpTime:          time.Unix(int64(hostInfo.Uptime), 0).Format("15:04:05"),
		OS:              hostInfo.OS,
		PlatForm:        hostInfo.Platform,
		PlatFormVersion: hostInfo.PlatformVersion,
		KernelArch:      hostInfo.KernelArch,
		HostID:          hostInfo.HostID,
	}
	return hosts
}

func GetNetInfo() map[string]NetNode {
	netInfo, _ := net.IOCounters(true)
	NetInfo := make(map[string]NetNode)
	for i := range netInfo {
		netIo := netInfo[i]
		NetInfo[netIo.Name] = NetNode{
			BytesSent:   int(netIo.BytesSent),
			BytesRect:   int(netIo.BytesRecv),
			PacketsSent: int(netIo.PacketsSent),
			PacketsRect: int(netIo.PacketsRecv),
		}
	}
	return NetInfo
}

func GetSysInfo() SysInfo {
	return SysInfo{
		CpuInfo:  GetCpuInfo(),
		MemInfo:  GetMemInfo(),
		DiskNode: GetDiskInfo(),
		HostInfo: GetHostInfo(),
		NetNode:  GetNetInfo(),
	}
}

//func GetMqttList(MqState *emqx.MqState) map[string]emqx.Sub {
//	return MqState.LinkList
//}
