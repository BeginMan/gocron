// +build gocron
// 调度中心

package main

import (
    "github.com/urfave/cli"
    "os"
    "gocron/cmd"
)

const AppVersion = "1.0"

func main() {
    app := cli.NewApp()
    app.Name = "gocron"
    app.Usage = "gocron service"
    app.Version = AppVersion
    s := make([]cli.Command,10,10)
    s[0] = cmd.CmdWeb
    app.Commands = s

    app.Flags = append(app.Flags, []cli.Flag{}...)
    app.Run(os.Args)
}
