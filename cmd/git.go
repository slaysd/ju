/*
Copyright © 2020 Jeeseung Han <jeeseung.han@gmail.com>

*/
package cmd

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/slaysd/ju/pkg/util"
	"github.com/spf13/cobra"
)

// gitCmd represents the git command
var gitCmd = &cobra.Command{
	Use:       "git",
	Short:     "Git toolbox",
	Long:      `Git toolbox like open web browser`,
	ValidArgs: []string{"open"},
	Args:      cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if output, err := exec.Command("git", "config", "--get", "remote.origin.url").Output(); err != nil {
			fmt.Println(err)
		} else {
			r, _ := regexp.Compile(`^(?:https:\/\/|git@)([\w\.]+)[\:\/]{1}([\w-\/]+)(?:.git)?$`)
			parse := r.FindStringSubmatch(strings.Replace(string(output), "\n", "", 1))
			if len(parse) == 3 {
				url := "https://" + parse[1] + "/" + parse[2]
				fmt.Println(url)
				util.OpenBrowser(url)
			} else {
				fmt.Println(parse, err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(gitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
