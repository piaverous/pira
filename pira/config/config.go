package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/piaverous/pira/pira/types"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// TODO(busser): document fields.
type Config struct {
	Jira   JiraConfig `mapstructure:"jira"`
	DryRun bool       `mapstructure:"dry_run"`
}

type JiraConfig struct {
	Board             JiraBoardConfig         `mapstructure:"board"`
	ProjectKey        string                  `mapstructure:"project_key"`
	Token             string                  `mapstructure:"token"`
	Url               string                  `mapstructure:"url"`
	User              string                  `mapstructure:"user"`
	RequestMaxResults string                  `mapstructure:"request_max_results"`
	CustomFields      []types.JiraCustomField `mapstructure:"custom_fields"`
	SprintConfig      JiraSprintConfig        `mapstructure:"sprint_config"`
}

type JiraSprintConfig struct {
	StoryPointFieldId string                   `mapstructure:"story_point_field_id"`
	TicketStatuses    []JiraSprintTicketStatus `mapstructure:"ticket_statuses"`
}

type JiraSprintTicketStatus struct {
	Name     string   `mapstructure:"name"`
	Statuses []string `mapstructure:"statuses"`
}

type JiraBoardConfig struct {
	Id   string `mapstructure:"id"`
	Name string `mapstructure:"name"`
	Type string `mapstructure:"type"`
}

func (c *Config) Load(flags *pflag.FlagSet) error {
	v := viper.New()

	// TODO: Implement this feature.
	// pira looks for configuration files called config.yaml, config.json,
	// config.toml, config.hcl, etc.
	v.SetConfigName("config")

	// pira looks for configuration files in the common configuration
	// directories.
	v.AddConfigPath("/etc/pira")
	v.AddConfigPath("$HOME/.local/.pira")
	v.AddConfigPath("$HOME/.pira")

	// Viper logs the configuration file it uses, if any.
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}
	}

	// configFile := v.ConfigFileUsed()
	// if configFile != "" {
	// 	fmt.Fprintf(os.Stderr, "Found config file at %s\n", configFile)
	// }

	// pira can be configured with environment variables that start with
	// pira_.
	v.SetEnvPrefix("pira")
	v.AutomaticEnv()

	// Options with dashes in flag names have underscores when set inside a
	// configuration file or with environment variables.
	flags.SetNormalizeFunc(func(fs *pflag.FlagSet, name string) pflag.NormalizedName {
		name = strings.ReplaceAll(name, "-", "_")
		return pflag.NormalizedName(name)
	})
	v.BindPFlags(flags)

	// Nested configuration options set with environment variables use an
	// underscore as a separator.
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	bindEnvironmentVariables(v, *c)

	return v.Unmarshal(c)
}

// bindEnvironmentVariables inspects iface's structure and recursively binds its
// fields to environment variables. This is a workaround to a limitation of
// Viper, found here:
// https://github.com/spf13/viper/issues/188#issuecomment-399884438
func bindEnvironmentVariables(v *viper.Viper, iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		val := ifv.Field(i)
		typ := ift.Field(i)
		tv, ok := typ.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch val.Kind() {
		case reflect.Struct:
			bindEnvironmentVariables(v, val.Interface(), append(parts, tv)...)
		default:
			v.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
