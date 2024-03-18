package cmd

import "github.com/spf13/cobra"

// ----------------------------------------------------------------------------
// INIT
// ----------------------------------------------------------------------------

func init() {
	rootCmd.AddCommand(componentCmd)
	componentCmd.PersistentFlags().StringP("search", "s", "", "specify filtering text")
	// 	componentCmd.AddCommand(godivaRuoliServiziCmd)

	// godivaRuoliServiziCmd.Flags().String("anagId", "", "uuid oppure id anagrafico del servizio")
	// godivaRuoliServiziCmd.Flags().Int("nodoId", 0, "id nodo dell'applicazione")
	// godivaRuoliServiziCmd.Flags().Bool("all", false, "mostra i ruoli per tutte le anagrafiche di tipo servizio")
}

var componentCmd = &cobra.Command{
	Use:   "component",
	Short: "Component Management",
	Long:  `Manage the component of the build ssytem`,

	Run: func(cmd *cobra.Command, args []string) {
		search, _ := cmd.Flags().GetString("search")
		println("search: ", search)
	},
}
