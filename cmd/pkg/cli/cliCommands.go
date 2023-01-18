package cli

import (
	"blockchainEtherium/cmd/pkg/blockchain"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var rootCmd = &cobra.Command{
	Use:   "blockchain",
	Short: "blockchain CLI",
	Long: `CLI
   
Работа внутри терминала`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "при выполнении CLI произошка ошибка '%s'", err)
		os.Exit(1)
	}
}

var readWalletCmd = &cobra.Command{
	Use:     "readWallet",
	Aliases: []string{"read"},
	Short:   "Read wallet",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		blockchain.ReadWalletSmartContract(args[0])
	},
}

var createWalletCmd = &cobra.Command{
	Use:     "createWallet",
	Aliases: []string{"create"},
	Short:   "Create wallet",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		argInt64, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			panic(err)
		}

		blockchain.CreateWalletSmartContract(args[0], argInt64)
	},
}

var sendBalanceWalletCmd = &cobra.Command{
	Use:     "sendBalance",
	Aliases: []string{"send"},
	Short:   "Send Balance wallet",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		argInt64, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			panic(err)
		}

		blockchain.SendMoneySnartContract(args[0], args[1], argInt64)
	},
}

var deployContractWalletCmd = &cobra.Command{
	Use:     "deployContract",
	Aliases: []string{"deploy"},
	Short:   "Deploy Contract",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		blockchain.DeploySmartContract()
	},
}

var checkContractCmd = &cobra.Command{
	Use:     "checkContract",
	Aliases: []string{"check"},
	Short:   "check Contract",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		blockchain.CheckContract(args[0])
	},
}

func init() {
	rootCmd.AddCommand(readWalletCmd)
	rootCmd.AddCommand(createWalletCmd)
	rootCmd.AddCommand(sendBalanceWalletCmd)
	rootCmd.AddCommand(deployContractWalletCmd)
	rootCmd.AddCommand(checkContractCmd)
}
