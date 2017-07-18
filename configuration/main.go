package configuration

import (
	"context"

	"github.com/BixData/bixagent/common/key"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/spf13/viper"
)

type Configuration struct {
	BinaryAddress    string   `json:"binaryAddress"      xml:"binaryAddress,attr"`
	BinaryNetwork    string   `json:"binaryNetwork"      xml:"binaryNetwork,attr"`
	Debug            bool     `json:"debug"              xml:"debug,attr"`
	HttpAddress      string   `json:"httpAddress"        xml:"httpAddress,attr"`
	HttpNetwork      string   `json:"httpNetwork"        xml:"httpNetwork,attr"`
	Sidecar          bool     `json:"sidecar"            xml:"sidecar,attr"`
	SidecarAddress   string   `json:"sidecarAddress"     xml:"sidecarAddress,attr"`
	SidecarCommand   []string `json:"sidecarCommand"     xml:"sidecarCommand,attr"`
	SidecarDirectory string   `json:"sidecarDirectory"   xml:"sidecarDirectory,attr"`
	SidecarNetwork   string   `json:"sidecarNetwork"     xml:"sidecarNetwork,attr"`
}

func GetConfiguration() Configuration {
	return Configuration{
		BinaryAddress:    viper.GetString(key.BIXAGENT_CONFIG_BINARY_ADDRESS),
		BinaryNetwork:    viper.GetString(key.BIXAGENT_CONFIG_BINARY_NETWORK),
		Debug:            viper.GetBool(key.BIXAGENT_CONFIG_DEBUG),
		HttpAddress:      viper.GetString(key.BIXAGENT_CONFIG_HTTP_ADDRESS),
		HttpNetwork:      viper.GetString(key.BIXAGENT_CONFIG_HTTP_NETWORK),
		Sidecar:          viper.GetBool(key.BIXAGENT_CONFIG_SIDECAR),
		SidecarAddress:   viper.GetString(key.BIXAGENT_CONFIG_SIDECAR_ADDRESS),
		SidecarCommand:   viper.GetStringSlice(key.BIXAGENT_CONFIG_SIDECAR_COMMAND),
		SidecarDirectory: viper.GetString(key.BIXAGENT_CONFIG_SIDECAR_DIRECTORY),
		SidecarNetwork:   viper.GetString(key.BIXAGENT_CONFIG_SIDECAR_NETWORK),
	}
}

func Get(ctx *routing.Context) error {
	return ctx.Write(GetConfiguration())
}

// Function for the "command pattern".
func Ozzo(ctx context.Context, routeGroup *routing.RouteGroup) error {
	routeGroup.Use(content.TypeNegotiator(content.XML, content.JSON))
	routeGroup.Get("", Get)
	return nil
}
