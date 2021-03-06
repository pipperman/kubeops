package main

import (
	"flag"
	"log"
	"os"

	kconf "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/pipperman/kubeops/app/inventory"
	"github.com/pipperman/kubeops/app/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "pipperman.kubeops.inventory"
	// Version is the version of the compiled software.
	Version string
	// configPath is the config flag.
	configPath string
	// baseDir is the config flag.
	baseDir string
	// ansibleConfDir
	ansibleConfDir string
	// ansibleTemplateFilePath
	ansibleTemplateFilePath string
	// ansibleVariablesName
	ansibleVariablesName string
)

func init() {
	flag.StringVar(&configPath, "conf", "./conf", "config path, eg: -conf ./conf")
	flag.StringVar(&baseDir, "basedir", "./dist/etc/kubeops", "base director, eg: -basedir /etc/errors")
	flag.StringVar(&ansibleConfDir, "ansibleConfDir", "/etc/ansible", "config path, eg: -ansibleConfDir /etc/ansible")
	flag.StringVar(&ansibleTemplateFilePath, "ansibleTemplateFilePath", "./dist/etc/kubeops/", "base director, eg: -ansibleTemplateFilePath /etc/errors/")
	flag.StringVar(&ansibleVariablesName, "variablesName", "variable.yml", "variable name, eg: -variablesName variable.yml")
	rootCmd.Flags().Bool("list", false, "")
}

var rootCmd = &cobra.Command{
	Use:   "inventory",
	Short: "A inventory provider for kubeops",
	Run: func(cmd *cobra.Command, args []string) {
		host := viper.GetString("server.host")
		port := viper.GetInt("server.port")
		provider := inventory.NewKubeOpsInventoryProvider(host, port)
		list, err := cmd.Flags().GetBool("list")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		if list {
			result, err := provider.ListHandler()
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			log.Println(result)
			os.Exit(0)
		}
		log.Printf("Host: %s\n Port: %d", host, port)
	},
}

func main() {
	c := kconf.New(
		kconf.WithSource(
			file.NewSource(configPath),
		),
		kconf.WithDecoder(func(kv *kconf.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)
	var bc config.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
