package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
	"time"

	"17vpn/internal/pritunl"
)

var connectCmd = &cobra.Command{
	Use:   "c",
	Short: "connect to specific profile",
	Run: func(cmd *cobra.Command, args []string) {
		if err := initConfig(); err != nil {
			color.Red(err.Error())
			return
		}
		p := pritunl.New()
		var targetProfile pritunl.Profile
		profiles := p.Profiles()
		conns := p.Connections()

		if err := list(profiles, conns); err != nil {
			color.Yellow(err.Error())
			return
		}
		id := strings.ToUpper(args[0])
		for i, profile := range profiles {
			if strconv.Itoa(i+1) == id || strings.ToUpper(id) == profile.Server {
				targetProfile = profile
				break
			}
		}
		if targetProfile == (pritunl.Profile{}) {
			color.Red("Profile not exists!")
			return
		}
		color.Yellow("Connecting %s...", targetProfile.Server)
		p.Connect(targetProfile.ID, password())
		timeout := time.NewTimer(30 * time.Second)

	Loop:
		for {
			select {
			case <-timeout.C:
				color.Red("Connect %s timeout!", targetProfile.Server)
				break Loop
			default:
				status := p.Connections()[targetProfile.ID].Status
				switch status {
				case "connected":
					color.Green("Connect %s completed!", targetProfile.Server)
					break Loop
				case "":
					color.Red("Connect %s failed!", targetProfile.Server)
					break Loop
				}
				time.Sleep(500 * time.Millisecond)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}
