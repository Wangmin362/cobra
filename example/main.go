package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "secgator",
	Short: "root命令的Short解释", // 命令解释
	Long:  "root命令的Long解释",
	// 注意：如果必须要求子命令，那么跟命令不能添加Run，这样的设计也显得非常的合理。应为直接执行根命令时无效的，因此也不需要处理根命令。
	// cobra会自动为我们打印命令的用法，但是如果指定了Run,或者RunE，就说明根命令时有效的，此时cobra并不会为我们自动处理
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Run call..")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.HelpFunc()(cmd, args)
		return nil
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true, // 是否禁用命令补全命令
	},
}

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "root命令的Short解释", // 命令解释
	Long:  "root命令的Long解释",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Run call..")
	},
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "获取帮助信息——Short", // 命令解释
	Long:  "获取帮助信息——Long",
	Run: func(cmd *cobra.Command, args []string) {
		cmd, _, e := cmd.Root().Find(args)
		if cmd == nil || e != nil {
			cmd.Printf("Unknown help topic %#q\n", args)
			cobra.CheckErr(cmd.Root().Usage())
		} else {
			cmd.InitDefaultHelpFlag()    // make possible 'help' flag to be shown
			cmd.InitDefaultVersionFlag() // make possible 'version' flag to be shown
			cobra.CheckErr(cmd.Help())
		}
	},
}

func main() {

	rootCmd.AddCommand(viewCmd)

	template := `用法:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [子命令]{{end}}{{if gt (len .Aliases) 0}}

别名:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}{{$cmds := .Commands}}{{if eq (len .Groups) 0}}

子命令:{{range $cmds}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{else}}{{range $group := .Groups}}

{{.Title}}{{range $cmds}}{{if (and (eq .GroupID $group.ID) (or .IsAvailableCommand (eq .Name "help")))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if not .AllChildCommandsHaveGroup}}

Additional Commands:{{range $cmds}}{{if (and (eq .GroupID "") (or .IsAvailableCommand (eq .Name "help")))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

子命令参数:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

全局参数:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

使用 "{{.CommandPath}} [子命令] --help" 获取命令的帮助信息.{{end}}
`

	rootCmd.SetHelpTemplate(template)
	rootCmd.SetHelpCommand(helpCmd)

	cobra.OnInitialize(func() {
		fmt.Println("OnInitialize 1")
	}, func() {
		fmt.Println("OnInitialize 2")
	})

	cobra.OnFinalize(func() {
		fmt.Println("OnFinalize 1")
	}, func() {
		fmt.Println("OnFinalize 2")
	})

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}
