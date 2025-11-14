package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "core-api",
	Short: "this api for compro",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Run(startCmd, nil)
	},
}



func Execute(){
	cobra.CheckErr(rootCmd.Execute())
}


func init() {
	cobra.OnInitialize(InitConfig)


	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file is default (.env)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func InitConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigFile(".env")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
