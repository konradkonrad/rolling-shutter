package cmd

import (
	"log"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/shutter-network/shutter/shuttermint/cmd/shversion"
	"github.com/shutter-network/shutter/shuttermint/keyper"
)

// keyperCmd represents the keyper command.
var keyperCmd = &cobra.Command{
	Use:   "keyper",
	Short: "Run a Shutter keyper node",
	Long: `This command runs a keyper node. It will connect to both an Ethereum and a
Shuttermint node which have to be started separately in advance.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return keyperMain()
	},
}

func init() {
	keyperCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
}

func readKeyperConfig() (keyper.Config, error) {
	viper.SetEnvPrefix("KEYPER")
	viper.BindEnv("ShuttermintURL")
	viper.BindEnv("SigningKey")
	viper.BindEnv("ValidatorSeed")
	viper.BindEnv("EncryptionKey")
	viper.BindEnv("DKGPhaseLength")

	viper.SetDefault("ShuttermintURL", "http://localhost:26657")

	defer func() {
		if viper.ConfigFileUsed() != "" {
			log.Printf("Read config from %s", viper.ConfigFileUsed())
		}
	}()
	var err error
	config := keyper.Config{}

	viper.AddConfigPath("$HOME/.config/shutter")
	viper.SetConfigName("keyper")
	viper.SetConfigType("toml")
	viper.SetConfigFile(cfgFile)

	err = viper.ReadInConfig()

	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// Config file not found
		if cfgFile != "" {
			return config, err
		}
	} else if err != nil {
		return config, err // Config file was found but another error was produced
	}
	err = config.Unmarshal(viper.GetViper())

	if err != nil {
		return config, err
	}

	if !filepath.IsAbs(config.DBDir) {
		r := filepath.Dir(viper.ConfigFileUsed())
		dbdir, err := filepath.Abs(filepath.Join(r, config.DBDir))
		if err != nil {
			return config, err
		}
		config.DBDir = dbdir
	}

	return config, err
}

func keyperMain() error {
	kc, err := readKeyperConfig()
	if err != nil {
		return errors.WithMessage(err, "Please check your configuration")
	}

	log.Printf(
		"Starting keyper version %s with signing key %s, using %s for Shuttermint",
		shversion.Version(),
		kc.Address().Hex(),
		kc.ShuttermintURL,
	)
	return errors.Errorf("keyper command not implemented")
}
