/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/slaysd/ju/pkg/mail"
	"github.com/slaysd/ju/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
)

// notifyCmd represents the notify command
var notifyCmd = &cobra.Command{
	Use:   "notify",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires at least one arg")
		}

		if host, port, username, password := viper.Get("smtp.host"), viper.Get("smtp.port"), viper.Get("smtp.username"), viper.Get("smtp.password"); host == nil || port == nil || username == nil || password == nil {
			return fmt.Errorf("requires smtp config")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		var (
			title     string
			message   string
			reference string
		)

		shell := exec.Command(args[0], args[1:]...)
		shell.Stdout = os.Stdout
		shell.Stderr = os.Stderr

		title = "[jsutil] Notify command '" + args[0] + "' "
		message = "Command: " + util.ToString(args)
		if err := shell.Run(); err == nil {
			title += "Success"
			reference = ""
		} else {
			title += "Fail"
			reference = err.Error()
		}

		sender := mail.MailSender{
			Host:     viper.GetString("smtp.host"),
			Port:     viper.GetString("smtp.port"),
			Username: viper.GetString("smtp.username"),
			Password: viper.GetString("smtp.password"),
		}
		sender.Send(title, message, reference)
	},
}

var notifyConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			host     string
			port     int
			username string
			password string
		)

		pre_host := viper.GetString("smtp.host")
		pre_port := viper.GetInt("smtp.port")
		pre_username := viper.GetString("smtp.username")
		pre_password := viper.GetString("smtp.password")

		fmt.Printf("SMTP Host (%s): ", pre_host)
		if cnt, _ := fmt.Scanln(&host); cnt > 0 {
			viper.Set("smtp.host", host)
		}
		fmt.Printf("SMTP Port (%d): ", pre_port)
		if cnt, _ := fmt.Scanln(&port); cnt > 0 {
			viper.Set("smtp.port", port)
		}
		fmt.Printf("SMTP Username (%s): ", pre_username)
		if cnt, _ := fmt.Scanln(&username); cnt > 0 {
			viper.Set("smtp.username", username)
		}
		fmt.Printf("SMTP Password (%s): ", pre_password)
		if cnt, _ := fmt.Scanln(&password); cnt > 0 {
			viper.Set("smtp.password", password)
		}

		if err := viper.WriteConfig(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	notifyCmd.AddCommand(notifyConfigCmd)
	rootCmd.AddCommand(notifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// notifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
