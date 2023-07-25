package kbimg

import (
	"fmt"
	"os"

	_ "github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Configs struct {
	Option    OptionConfig   `mapstructure:"option_config"`
	Partitons PartitonConfig `mapstructure:"storage config"`
	Disks     []DiskConfig   `mapstructure:"disk_config"`
	Files     []FileConfig   `mapstructure:"file_config"`
	Users     []UserConfig   `mapstructure:"user_config"`
	Host      HostConfig     `mapstructure:"host_config"`
	Grub      GrubConfig     `mapstructure:"grub_config"`
	Systemd   ServiceConfig  `mapstructure:"systemd_service_config"`
}

// 镜像制作方式
type OptionConfig struct {
	Image string `mapstructure:"image"`
	P     string `mapstructure:"p"`
	V     string `mapstructure:"v"`
	B     string `mapstructure:"b"`
	E     string `mapstructure:"e"`
	D     string `mapstructure:"d"`
}

// 存储配置
// 分区划分
type PartitonConfig struct {
	BootSize    int `mapstructure:"boot_size"`
	RootASize   int `mapstructure:"rootA_size"`
	RootBSize   int `mapstructure:"rootB_size"`
	PersistSize int `mapstructure:"persist_size"`
}

// 磁盘指定与挂载
type DiskConfig struct {
	Disk  string `mapstructure:"disk"`
	Mount string `mapstructure:"mount"`
}

// 文件配置
// 创建新文件/文件夹
type FileConfig struct {
	Type string `mapstructure:"type"`
	Name string `mapstructure:"name"`
	Path string `mapstructure:"path"`
}

// 用户配置
// 配置新用户/用户组
type UserConfig struct {
	Name     string   `mapstructure:"name"`
	Password string   `mapstructure:"password"`
	Groups   []string `mapstructure:"groups"`
}

// 主机名称
type HostConfig struct {
	HostName string `mapstructure:"hostname"`
}

// 设置 grub 密码
type GrubConfig struct {
	Password string `mapstructure:"password"`
}

// systemd service 配置
type ServiceConfig struct {
	Name  string `mapstructure:"name"`
	Start bool   `mapstructure:"start"`
}

func initConfig() {
	viper.SetConfigFile("S:/Go/GoCode/KubeOS/scripts/kbimg.yaml")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	//打印 viper 当前使用的配置文件
	fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())

	var config Configs

	// 解析配置文件
	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("Failed to decode YAML:", err)
		return
	} else {
		fmt.Println("Success to decode YAML")
		printConfig(config)
	}

	// 在这里可以使用 config 结构体生成对应的 shell 文件内容
	// 例如：config.Host.HostName 获取 host_config 的 hostname 字段值
}

func printConfig(config Configs) {
	fmt.Println("Option Config:")
	fmt.Println("Image:", config.Option.Image)
	fmt.Println("P:", config.Option.P)
	fmt.Println("V:", config.Option.V)
	fmt.Println("B:", config.Option.B)
	fmt.Println("E:", config.Option.E)
	fmt.Println("D:", config.Option.D)
	fmt.Println()

	fmt.Println("Partition Config:")
	fmt.Println("Boot Size:", config.Partitons.BootSize)
	fmt.Println("RootA Size:", config.Partitons.RootASize)
	fmt.Println("RootB Size:", config.Partitons.RootBSize)
	fmt.Println("Persist Size:", config.Partitons.PersistSize)
	fmt.Println()

	fmt.Println("Disk Config:")
	for _, disk := range config.Disks {
		fmt.Println("Disk:", disk.Disk)
		fmt.Println("Mount:", disk.Mount)
	}
	fmt.Println()

	fmt.Println("File Config:")
	for _, file := range config.Files {
		fmt.Println("Type:", file.Type)
		fmt.Println("Name:", file.Name)
		fmt.Println("Path:", file.Path)
	}
	fmt.Println()

	fmt.Println("User Config:")
	for _, user := range config.Users {
		fmt.Println("Name:", user.Name)
		fmt.Println("Password:", user.Password)
		fmt.Println("Groups:", user.Groups)
	}
	fmt.Println()

	fmt.Println("Host Config:")
	fmt.Println("HostName:", config.Host.HostName)
	fmt.Println()

	fmt.Println("Grub Config:")
	fmt.Println("Password:", config.Grub.Password)
	fmt.Println()

	fmt.Println("Systemd Service Config:")
	fmt.Println("Name:", config.Systemd.Name)
	fmt.Println("Start:", config.Systemd.Start)
}
