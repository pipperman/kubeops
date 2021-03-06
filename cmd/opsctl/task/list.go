package task

import (
	"fmt"
	"log"

	"github.com/pipperman/kubeops/app/opsctl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var taskListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		host := viper.GetString("server.host")
		port := viper.GetInt("server.port")
		c := opsctl.NewKubeOpsClient(host, port)
		rs, err := c.ListResult()
		if err != nil {
			log.Fatal(err)
		}
		for _, r := range rs {
			out := fmt.Sprintf("result_id: %s,  start_time: %s,   end_time: %s,   finished: %t,  success: %t",
				r.Id, r.StartTime, r.EndTime, r.Finished, r.Success)
			fmt.Println(out)
		}
	},
}

func init() {
}
