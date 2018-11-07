// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"AgendaGo/service"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// deleteMeetingCmd represents the deleteMeeting command
var deleteMeetingCmd = &cobra.Command{
	Use:   "delete meeting",
	Short: "delete the meeting",
	Long:  `input is the title,then according to the input to delete the meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		err := service.DeleteMeetingByTitle(title)
		if err == nil {
			fmt.Printf("delete the %s meeting successfully\n", title)
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteMeetingCmd)
	// deleteMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	deleteMeetingCmd.Flags().StringP("title", "t", "", "/")
}
