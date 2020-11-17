/*
Copyright © 2020 Jeeseung Han <jeeseung.han@gmail.com>

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
	"os"
	"os/exec"

	"github.com/go-playground/validator"
	"github.com/slaysd/ju/pkg/mail"
	"github.com/slaysd/ju/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	validate *validator.Validate
	notify   notifier
)

type notifier interface {
	Send(title string, message string, reference string) error
}

// notifyCmd represents the notify command
var notifyCmd = &cobra.Command{
	Use:   "notify",
	Short: "Notify status of shell command",
	Long:  `Notify status of shell command via email. You must be set SMTP configuration (type 'ju notify config')`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("Requires at least one arg")
		}

		if host, port, username, password := viper.Get("smtp.host"), viper.Get("smtp.port"), viper.Get("smtp.username"), viper.Get("smtp.password"); host == nil || port == nil || username == nil || password == nil {
			return fmt.Errorf("Requires SMTP config, please set SMTP config via 'ju notifiy config'")
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

		notify = &mail.MailSender{
			Host:     viper.GetString("smtp.host"),
			Port:     viper.GetString("smtp.port"),
			Username: viper.GetString("smtp.username"),
			Password: viper.GetString("smtp.password"),
		}
		notify.Send(title, message, reference)
	},
}

var notifyConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Config notify sender configuration",
	Long:  `Config notify sender configuration. You should to know smtp host, port, username, password`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			host     string
			port     int
			username string
			password string
		)

		preHost := viper.GetString("smtp.host")
		prePort := viper.GetInt("smtp.port")
		preUsername := viper.GetString("smtp.username")
		prePassword := viper.GetString("smtp.password")

		validate = validator.New()

		fmt.Printf("SMTP Host (%s): ", preHost)
		if cnt, _ := fmt.Scanln(&host); cnt > 0 {
			if err := validate.Var(host, "required,hostname"); err != nil {
				fmt.Println(err.Error())
				return
			}
			viper.Set("smtp.host", host)
		}

		fmt.Printf("SMTP Port (%d): ", prePort)
		if cnt, _ := fmt.Scanln(&port); cnt > 0 {
			if err := validate.Var(port, "required"); err != nil {
				fmt.Println(err.Error())
				return
			}
			viper.Set("smtp.port", port)
		}
		fmt.Printf("SMTP Username (%s): ", preUsername)
		if cnt, _ := fmt.Scanln(&username); cnt > 0 {
			if err := validate.Var(username, "required,email"); err != nil {
				fmt.Println(err.Error())
				return
			}
			viper.Set("smtp.username", username)
		}
		fmt.Printf("SMTP Password (%s): ", prePassword)
		if cnt, _ := fmt.Scanln(&password); cnt > 0 {
			if err := validate.Var(password, "required"); err != nil {
				fmt.Println(err.Error())
				return
			}
			viper.Set("smtp.password", password)
		}

		if err := viper.WriteConfig(); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Done.")
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
