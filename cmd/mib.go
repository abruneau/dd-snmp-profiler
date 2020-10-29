package cmd

import (
	"github.com/abruneau/dd-snmp-profiler/parsers"
	"github.com/spf13/cobra"
)

// mibCmd represents the mib command
var (
	// Used for flags.
	path   string
	module string

	mibCmd = &cobra.Command{
		Use:   "mib",
		Short: "Generate profile from MIB file",
		Long: `mib (dd-snmp-profiler mib) will generate an new snmp profile
	for Datadog agent based on the provided MIB file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return parsers.ParseMIBModule(path, module, outputFile, debug)
		},
	}
)

func init() {
	rootCmd.AddCommand(mibCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	mibCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "Path to add")
	mibCmd.PersistentFlags().StringVarP(&module, "module", "m", "", "Module to load")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mibCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
