package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/portworx/sched-ops/k8s/core"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/cmd/util"
)

const (
	url string = "https://pxlite.loggly.com/apiv2/search?q=%s"
)

var (
	secret          string
	secretNamespace string
	query           string
)

func main() {
	if err := NewCommand().Execute(); err != nil {
		os.Exit(1)
	}
}

func requestQuery(secret, namespace, query string) {
	val, err := core.Instance().GetSecret(secret, secretNamespace)
	if err != nil {
		err = &metering.ErrBillingAPIUnreachable{
			Reason: fmt.Sprintf("Error reading reporting secret %v:%v : %v",
				secret, secretNamespace, err),
		}
		return err
	}

	//https://pxlite.loggly.com/apiv2/search?q=json.Volumes&from=-10m&until=now
}

func retrieveData() {
	//https://<subdomain>.loggly.com/apiv2/events?rsid=<id>
}

func NewCommand() *cobra.Command {
	cmds := &cobra.Command{
		Use:   "lparse",
		Short: "loggly parsing tool.",
	}

	cmds.PersistentFlags().StringVarP(&secret, "secret", "s", "", "Name of the secret in which the loggly apitoken can be found")
	cmds.PersistentFlags().StringVarP(&secretNamespace, "secretNamespace", "n", "", "Namespace in which the loggly apitoken secret can be found")
	cmds.PersistentFlags().StringVarP(&query, "query", "q", "", "Query to send to Loggly")

	cmds.AddCommand(
		newParseCommand(),
	)
	cmds.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	err := flag.CommandLine.Parse([]string{})
	if err != nil {
		util.CheckErr(err)
		return nil
	}

	return cmds
}

func newParseCommand() *cobra.Command {
	parseCommand := &cobra.Command{
		Use:   "parse",
		Short: "Start sparsing",
		Run: func(c *cobra.Command, args []string) {
			resp, err = requestQuery(secret, namespace, query)
			if err != nil {
				logrus.Errorf("listvols failed with: %v", err)
				return
			}
		},
	}
	return parseCommand
}
