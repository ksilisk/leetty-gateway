package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"leetty-gateway/internal/logger"
	"os"
)

const pathEnvVariable = "LEETTY_GATEWAY_CONFIG_PATH"
const applicationProfileEnvVariable = "LEETTY_GATEWAY_APP_PROFILE"

type commandLineArgs struct {
	profile string
}

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	KafkaBrokers []string `yaml:"kafka-brokers"`
	Mapping      []struct {
		Endpoint   string `yaml:"endpoint"`
		KafkaTopic string `yaml:"kafka-topic"`
		Partition  int    `yaml:"partition"`
	} `yaml:"mappings"`
}

func ParseConfig() (conf *Config, error error) {
	var args = parseArgs()
	logger.Logger.Info("parsing configuration file")
	var data, err = os.ReadFile(getConfigPath(getProfile(args)))
	if err != nil {
		return nil, err
	}
	var config = Config{}
	var expandedData = os.ExpandEnv(string(data))
	err = yaml.Unmarshal([]byte(expandedData), &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func getConfigPath(profile string) string {
	var path, present = os.LookupEnv(pathEnvVariable)
	if !present {
		return "configs/config_" + profile + ".yml"
	}
	logger.Logger.Info("Found " + path + " environment variable")
	return path
}

func getProfile(args *commandLineArgs) string {
	var profile, present = os.LookupEnv(applicationProfileEnvVariable)
	if !present {
		profile = args.profile
	}
	logger.Logger.Info("Using Profile: '" + profile + "'")
	return profile
}

func parseArgs() *commandLineArgs {
	var profile = flag.String("profile", "dev", "the environment to use")
	var help = flag.Bool("help", false, "display this help and exit")
	flag.Parse()
	if *help == true {
		flag.Usage()
		os.Exit(0)
	}
	return &commandLineArgs{*profile}
}
