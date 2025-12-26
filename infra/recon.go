package infra

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type HardwareReport struct {
	Platform    string
	Family      string
	Version     string
	PhysicalCPU int
	LogicalCPU  int
	CPUPercents []float64
	CPUInfo     string
	Memory      *mem.VirtualMemoryStat
}

func HardInfo() HardwareReport {
	platform, family, version := GetHost()
	physical, logical, percents := GetCpu()
	cpuInfo := InfoCpu()
	memory := GetMemory()

	return HardwareReport{
		Platform:    platform,
		Family:      family,
		Version:     version,
		PhysicalCPU: physical,
		LogicalCPU:  logical,
		CPUPercents: percents,
		CPUInfo:     cpuInfo,
		Memory:      memory,
	}
}

func GetMemory() *mem.VirtualMemoryStat {
	v, _ := mem.VirtualMemory()
	log.Debugf("Total: %v, Available: %v, UsedPercent:%f%%\n", v.Total, v.Available, v.UsedPercent)

	return v
}

func GetCpu() (int, int, []float64) {
	physicalCnt, _ := cpu.Counts(false)
	logicalCnt, _ := cpu.Counts(true)
	log.Debugf("Physical: %v, Logic: %v\n", physicalCnt, logicalCnt)

	perPercents, _ := cpu.Percent(3*time.Second, true)
	log.Debugf("Perpercent: %v", perPercents)

	return physicalCnt, logicalCnt, perPercents
}

func InfoCpu() string {
	infos, _ := cpu.Info()
	data, _ := json.MarshalIndent(infos, "", "  ")
	log.Debug(string(data))
	return string(data)
}

func GetHost() (string, string, string) {

	platform, family, version, _ := host.PlatformInformation()
	log.Debug(platform)
	log.Debug(family)
	log.Debug(version)

	return platform, family, version
}
