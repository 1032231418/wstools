package main

import (
	"fmt"
	"os"

	"github.com/czxichen/wstools/command"
	"github.com/spf13/cobra"
)

var (
	Author  string = "dijielin"
	EMail   string = "dijielin@qq.com"
	BuildTS string
	GitHash string
)

var version = &cobra.Command{
	Use:   "version",
	Short: "查看软件基础信息",
	Run: func(*cobra.Command, []string) {
		fmt.Printf("Author:\t%s\r\nEmail:\t%s\r\nTime:\t%s\r\nVersion:%s\r\n", Author, EMail, BuildTS, GitHash)
	},
}

var rootCMD = &cobra.Command{
	Use: os.Args[0],
}

func main() {
	command.HelpFunc(rootCMD)
	rootCMD.AddCommand(version, command.Compress, command.Md5sum, command.Net, command.Deploy,
		command.Find, command.Compare, command.Ftp, command.RSA, command.Tail, command.Watchdog,
		command.Http, command.Mail, command.Replace, command.SysInfo, command.Ssh, command.Fsnotify)
	rootCMD.Execute()
}
