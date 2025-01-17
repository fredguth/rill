package project

import (
	"fmt"

	"github.com/rilldata/rill/cli/pkg/cmdutil"
	adminv1 "github.com/rilldata/rill/proto/gen/rill/admin/v1"
	"github.com/spf13/cobra"
)

func ResetCmd(ch *cmdutil.Helper) *cobra.Command {
	var project, path string
	var force bool
	cfg := ch.Config

	resetCmd := &cobra.Command{
		Use:               "reset [<project-name>]",
		Args:              cobra.MaximumNArgs(1),
		Short:             "Re-deploy project",
		Long:              "Create a new deployment for the project (and tear down the current one)",
		PersistentPreRunE: cmdutil.CheckChain(cmdutil.CheckAuth(cfg), cmdutil.CheckOrganization(cfg)),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			client, err := cmdutil.Client(cfg)
			if err != nil {
				return err
			}
			defer client.Close()

			if len(args) > 0 {
				project = args[0]
			}

			if !cmd.Flags().Changed("project") && len(args) == 0 && cfg.Interactive {
				var err error
				project, err = inferProjectName(ctx, client, cfg.Org, path)
				if err != nil {
					return err
				}
			}

			if !force {
				msg := "This will create a new deployment, causing downtime as data sources are reloaded from scratch. If you just need to refresh data, use `rill project refresh`. Do you want to continue?"
				if !cmdutil.ConfirmPrompt(msg, "", false) {
					return nil
				}
			}

			_, err = client.TriggerRedeploy(ctx, &adminv1.TriggerRedeployRequest{Organization: cfg.Org, Project: project})
			if err != nil {
				return err
			}

			fmt.Printf("Triggered project reset. To see status, run `rill project status --project %s`.\n", project)

			return nil
		},
	}

	resetCmd.Flags().SortFlags = false
	resetCmd.Flags().StringVar(&project, "project", "", "Project name")
	resetCmd.Flags().StringVar(&path, "path", ".", "Project directory")
	resetCmd.Flags().BoolVar(&force, "force", false, "Force reset even if project is already deployed")
	return resetCmd
}
