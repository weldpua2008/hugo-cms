package config

import (
	"github.com/spf13/viper"
)

func LoadDefaultSettings() {
	viper.SetDefault("MetaDataFormat", "toml")
	viper.SetDefault("DisableRSS", false)
	viper.SetDefault("DisableSitemap", false)
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("StaticDir", "static")
	viper.SetDefault("PublishDir", "public")
	viper.SetDefault("DefaultLayout", "post")
	viper.SetDefault("Verbose", false)
	viper.SetDefault("IgnoreCache", false)
	viper.SetDefault("CanonifyURLs", false)
	viper.SetDefault("RelativeURLs", false)
	viper.SetDefault("DefaultExtension", "html")
	viper.SetDefault("PygmentsUseClasses", false)
	viper.SetDefault("Paginate", 10)
	viper.SetDefault("PaginatePath", "page")
}
