package config

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func Set(args []string) error {
	var key, value string
	if len(args) == 1 && !strings.Contains(args[0], "=") {
		parts := strings.SplitN(args[0], "=", 2)
		key = parts[0]
		value = parts[1]
	} else if len(args) == 2 {
		key = args[0]
		value = args[1]
	} else if len(args) == 3 && args[1] == "=" {
		key = args[0]
		value = args[2]
	} else {
		return fmt.Errorf("please use either 'set <item>=<value>' or 'set <item> <value>' format")
	}
	config, err := loadConfig()
	if err != nil {
		return fmt.Errorf("config file is damaged, please try repair it or use 'ktctl config unset --all'")
	}
	err = setConfigValue(config, key, value)
	if err != nil {
		return fmt.Errorf("%s, please check available config items with 'ktctl config show --all'", err)
	}
	return saveConfig(config)
}

func SetHandle(cmd *cobra.Command) {
	cmd.ValidArgsFunction = setConfigValidator
}

func setConfigValue(config map[string]map[string]string, key string, value string) error {
	group, item, err := parseConfigItem(key)
	if err != nil {
		return err
	}
	if _, exist := config[group]; exist {
		config[group][item] = value
	} else {
		config[group] = map[string]string{item: value}
	}
	return nil
}

func setConfigValidator(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var items []string
	if len(args) == 0 {
		travelConfigItem(func(groupName string, itemName string) {
			items = append(items, groupName+"."+itemName)
		})
	}
	return items, cobra.ShellCompDirectiveNoFileComp
}
