package kbimg

import (
	"fmt"

	"github.com/spf13/cobra"
)

var applyConfig bool
var cfgFile string = "S:/Go/GoCode/KubeOS/scripts/kbimg.yaml"

// ApplyUserConfiguration 创建一个 *cobra.Command 对象. 之后，可以使用 Command 对象的 Execute 方法来启动应用程序.
func ApplyUserConfiguration() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "kbimg",
		// 命令的详细描述
		Long: `Generate and execute shell script based on kbimg.yaml`,

		// 命令出错时，不打印帮助信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}

	// 在此处定义标志和配置设置

	// 定义config 命令，调用 --config 之后，applyConfig == true，开始读取配置并生成文件
	cmd.PersistentFlags().BoolVar(&applyConfig, "config", false, "Apply your configuration")
	return cmd
}

// run 函数是实际的功能入口函数.
func run() error {
	// 根据 applyConfig 的值判断是否应用用户配置
	if applyConfig == true {
		fmt.Println("Apply your configuration")
		initConfig()
	} else {
		fmt.Println("Hello kbimg!")
	}
	return nil
}
