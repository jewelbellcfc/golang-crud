/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"GOLAND/database"

	"log"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: Update}

func init() {
	updateCmd.Flags().String("id","","id want update")
	updateCmd.MarkFlagRequired("id")
	updateCmd.Flags().String("fullname","","new name")
	updateCmd.Flags().String("gender","","new gender")
	rootCmd.AddCommand(updateCmd)
}
func Update(cmd*cobra.Command,grgs []string){
	id:= cmd.Flag("id").Value.String()
	fullname:= cmd.Flag("fullname").Value.String()
	gender:= cmd.Flag("gender").Value.String()
	if fullname==""||gender=="" {
		log.Fatal("full name and gender is required")
	}
	if gender =="updating"|| gender=="male"||gender=="female"{
		database.UpdatePerson(id,fullname,gender)
	} else{
		log.Fatalf("provide valid gender -male,felame or updating")
	}
}
