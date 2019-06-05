package command

import (
	"fmt"

	"github.com/oclaussen/dodo/pkg/stage"
	"github.com/oclaussen/dodo/pkg/types"
	"github.com/oclaussen/go-gimme/configfiles"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func NewStageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "stage",
		Short:            "Manage stages",
		TraverseChildren: true,
		SilenceUsage:     true,
	}

	cmd.AddCommand(NewUpCommand())
	cmd.AddCommand(NewDownCommand())
	return cmd
}

func NewUpCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "Create or start a stage",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := loadStageConfig(args[0])
			if err != nil {
				return err
			}
			target, err := stage.LoadStage(args[0], conf)
			if err != nil {
				return err
			}
			return target.Up()
		},
	}
}

func NewDownCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "Remove or pause a stage",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: do we actually need the config?
			conf, err := loadStageConfig(args[0])
			if err != nil {
				return err
			}
			target, err := stage.LoadStage(args[0], conf)
			if err != nil {
				return err
			}
			return target.Down(false)
		},
	}
}

func loadStageConfig(name string) (*types.Stage, error) {
	configFile, err := configfiles.GimmeConfigFiles(&configfiles.Options{
		Name:                      "dodo",
		Extensions:                []string{"yaml", "yml", "json"},
		IncludeWorkingDirectories: true,
		Filter: func(configFile *configfiles.ConfigFile) bool {
			return containsStage(configFile, name)
		},
	})
	if err != nil {
		return nil, err
	}

	var mapType map[interface{}]interface{}
	if err := yaml.Unmarshal(configFile.Content, &mapType); err != nil {
		return nil, err
	}

	config, err := types.DecodeGroup(configFile.Path, mapType)
	if err != nil {
		return nil, err
	}

	if result, ok := config.Stages[name]; ok {
		return &result, nil
	}

	return nil, fmt.Errorf("could not find stage %s in file %s", name, configFile.Path)
}

func containsStage(configFile *configfiles.ConfigFile, name string) bool {
	var mapType map[interface{}]interface{}
	if err := yaml.Unmarshal(configFile.Content, &mapType); err != nil {
		return false
	}

	config, err := types.DecodeNames(configFile.Path, "", mapType)
	if err != nil {
		return false
	}

	_, ok := config.Stages[name]
	return ok
}