package top

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	kcmd "k8s.io/kubernetes/pkg/kubectl/cmd"
	"k8s.io/kubernetes/pkg/kubectl/cmd/templates"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

const (
	TopRecommendedName = "top"
)

var topLong = templates.LongDesc(`
	Show usage statistics of resources on the server

	This command analyzes resources managed by the platform and presents current
	usage statistics.`)

func NewCommandTop(name, fullName string, f cmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	// Parent command to which all subcommands are added.
	cmds := &cobra.Command{
		Use:   name,
		Short: "Show usage statistics of resources on the server",
		Long:  topLong,
		Run:   cmdutil.DefaultSubCommandRun(streams.ErrOut),
	}

	cmdTopNode := kcmd.NewCmdTopNode(f, nil, streams)
	cmdTopPod := kcmd.NewCmdTopPod(f, nil, streams)

	cmds.AddCommand(NewCmdTopImages(f, fullName, TopImagesRecommendedName, streams))
	cmds.AddCommand(NewCmdTopImageStreams(f, fullName, TopImageStreamsRecommendedName, streams))
	cmdTopNode.Long = templates.LongDesc(cmdTopNode.Long)
	cmdTopNode.Example = templates.Examples(cmdTopNode.Example)
	cmdTopPod.Long = templates.LongDesc(cmdTopPod.Long)
	cmdTopPod.Example = templates.Examples(cmdTopPod.Example)
	cmds.AddCommand(cmdTopNode)
	cmds.AddCommand(cmdTopPod)
	return cmds
}
