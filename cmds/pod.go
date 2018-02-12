package cmds

import (
	"github.com/spf13/cobra"
	"github.com/appscode/go/term"
	. "github.com/sanjid133/pod-controller/pkg"
)

func NewCmdPod() *cobra.Command  {
	var kubeconfig string
	var master string
	cmd := &cobra.Command{
		Use: "run",
		Short: "run the controller",
		DisableAutoGenTag:true,
		Run: func(cmd *cobra.Command, args []string) {
			ctl, err := InitializeClient(kubeconfig, master)
			if err != nil {
				term.Println(err)
			}
			ctl.Run()

		},
	}
	cmd.Flags().StringVar(&kubeconfig, "kubeconfig", "", "kubeconfig file")
	cmd.Flags().StringVar(&master, "master", "", "master url")
	return cmd
}
