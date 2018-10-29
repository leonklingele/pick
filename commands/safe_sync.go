package commands

import (
	"github.com/bndw/pick/backends"
	fileBackend "github.com/bndw/pick/backends/file"
	"github.com/bndw/pick/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func init() {
	// TODO: This command is deprecated and will be removed soon.
	rootCmd.AddCommand(&cobra.Command{
		Use:   "sync [other-pick-safe]",
		Short: "Sync current safe with another pick safe",
		Long:  "The sync command is used to sync the current pick safe with another pick safe.",
		Run: func(cmd *cobra.Command, args []string) {
			runMovedCommand(Sync, cmd, args, "safe sync")
		},
		Hidden: true,
	})
}

func Sync(args []string, flags *pflag.FlagSet) error {
	otherSafePath, err := parseSyncArgs(args)
	if err != nil {
		return err
	}

	safeLoader := newSafeLoader(true)

	safe, err := safeLoader.Load()
	if err != nil {
		return err
	}
	safeLoader.RememberPassword()

	otherStorageConfig := config.Storage
	// TODO(leon): :(
	otherStorageConfig.Settings["path"] = otherSafePath
	otherBackendClient, err := backends.NewWithType(fileBackend.ClientName, &otherStorageConfig)
	if err != nil {
		return err
	}
	otherSafe, err := safeLoader.LoadWithBackendClient(otherBackendClient)
	if err != nil {
		return err
	}

	if err := safe.SyncWith(otherSafe); err != nil {
		return err
	}
	return nil
}

func parseSyncArgs(args []string) (otherSafePath string, err error) {
	if len(args) != 1 {
		err = errors.ErrInvalidCommandUsage
		return
	}

	otherSafePath = args[0]

	return
}
