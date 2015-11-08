package config

import (
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

// inspireted by https://github.com/spf13/hugo/blob/3a412543f617cf9fa460061aa5a33db43675c9f9/commands/hugo.go

//Its a root command. Every other command attached to HugoCMSCmd is a child command to it.
var HugoCMSCmd = &cobra.Command{
	Use:   "hugocms",
	Short: "hugocms for building your site",
	Long: `hugo is the main command, used to start your CMS site.
Complete documentation is available at https://github.com/weldpua2008/hugo-cms/ .`,
	Run: func(cmd *cobra.Command, args []string) {
		InitializeConfig()
		build()
	},
}

var hugoCMSCmdV *cobra.Command

//Flags that are to be added to commands.
var BuildWatch, IgnoreCache, Draft, Future, UglyURLs, CanonifyURLs, Verbose, Logging, VerboseLog, DisableRSS, DisableSitemap, PluralizeListTitles, PreserveTaxonomyNames, NoTimes bool
var Source, CacheDir, Destination, Theme, BaseURL, CfgFile, LogFile, Editor string

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

// InitializeConfig initializes a config file with sensible default configuration flags.
func InitializeConfig() {
	viper.SetConfigFile(CfgFile)
	if Source == "" {
		viper.AddConfigPath(".")
	} else {
		viper.AddConfigPath(Source)
	}
	err := viper.ReadInConfig()
	if err != nil {
		jww.ERROR.Println("Unable to locate Config file.")
	}
	LoadDefaultSettings()

	if hugoCMSCmdV.PersistentFlags().Lookup("disableRSS").Changed {
		viper.Set("DisableRSS", DisableRSS)
	}

	if hugoCMSCmdV.PersistentFlags().Lookup("disableSitemap").Changed {
		viper.Set("DisableSitemap", DisableSitemap)
	}

	if hugoCMSCmdV.PersistentFlags().Lookup("verbose").Changed {
		viper.Set("Verbose", Verbose)
	}

	if hugoCMSCmdV.PersistentFlags().Lookup("pluralizeListTitles").Changed {
		viper.Set("PluralizeListTitles", PluralizeListTitles)
	}

	if hugoCMSCmdV.PersistentFlags().Lookup("preserveTaxonomyNames").Changed {
		viper.Set("PreserveTaxonomyNames", PreserveTaxonomyNames)
	}

	if hugoCMSCmdV.PersistentFlags().Lookup("editor").Changed {
		viper.Set("NewContentEditor", Editor)
	}

	if hugoCMSCmdV.PersistentFlags().Lookup("logFile").Changed {
		viper.Set("LogFile", LogFile)
	}
	if BaseURL != "" {
		if !strings.HasSuffix(BaseURL, "/") {
			BaseURL = BaseURL + "/"
		}
		viper.Set("BaseURL", BaseURL)
	}

	if !viper.GetBool("RelativeURLs") && viper.GetString("BaseURL") == "" {
		jww.ERROR.Println("No 'baseurl' set in configuration or as a flag. Features like page menus will not work without one.")
	}

	if Theme != "" {
		viper.Set("theme", Theme)
	}

	if Destination != "" {
		viper.Set("PublishDir", Destination)
	}

	if Source != "" {
		viper.Set("WorkingDir", Source)
	} else {
		dir, _ := os.Getwd()
		viper.Set("WorkingDir", dir)
	}

	if hugoCMSCmdV.PersistentFlags().Lookup("ignoreCache").Changed {
		viper.Set("IgnoreCache", IgnoreCache)
	}

	if CacheDir != "" {
		if helpers.FilePathSeparator != CacheDir[len(CacheDir)-1:] {
			CacheDir = CacheDir + helpers.FilePathSeparator
		}
		isDir, err := helpers.DirExists(CacheDir, hugofs.SourceFs)
		utils.CheckErr(err)
		if isDir == false {
			mkdir(CacheDir)
		}
		viper.Set("CacheDir", CacheDir)
	} else {
		viper.Set("CacheDir", helpers.GetTempDir("hugo_cache", hugofs.SourceFs))
	}

	if VerboseLog || Logging || (viper.IsSet("LogFile") && viper.GetString("LogFile") != "") {
		if viper.IsSet("LogFile") && viper.GetString("LogFile") != "" {
			jww.SetLogFile(viper.GetString("LogFile"))
		} else {
			jww.UseTempLogFile("hugo")
		}
	} else {
		jww.DiscardLogging()
	}

	if viper.GetBool("verbose") {
		jww.SetStdoutThreshold(jww.LevelInfo)
	}

	if VerboseLog {
		jww.SetLogThreshold(jww.LevelInfo)
	}

}
