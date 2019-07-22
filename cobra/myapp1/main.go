//main.go
package main

import (
    "fmt"
    "github.com/go-example/cobra/myapp1/node"
    "github.com/spf13/cobra"
    "os"
)

var versionFlag bool

var mainCmd = &cobra.Command{
    Use: "command",
    Run: func(cmd *cobra.Command, args []string) {
        if versionFlag {
            fmt.Println("目前版本为1.0.0")
        } else {
            cmd.HelpFunc()(cmd, args)
        }
    },
}

func main() {
    mainFlags := mainCmd.PersistentFlags()
    mainFlags.BoolVarP(&versionFlag, "version", "v", false, "打印系统版本")

    mainCmd.AddCommand(node.Cmd())
    if mainCmd.Execute() != nil {
        os.Exit(1)
    }
}
