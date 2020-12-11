package main

import (
	"clash-subscription-updater/overider"
	"clash-subscription-updater/updater"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

func printVersion() {
	fmt.Print("<%VERSION%>")
}
func init() {
	viper.SetConfigName("clash-subscription-updater")
	viper.AddConfigPath("$HOME/.config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	defaultClashDir := os.Getenv("HOME") + "/.config/clash"
	pflag.StringP("clash-config-dir", "d", defaultClashDir, "config directory of clash")
	pflag.IntP("interval", "i", 60, "interval to fetch configuration (minutes)")
	pflag.BoolP("help", "h", false, "show this message")
	pflag.Bool("override", false, "override the existed config file")
	pflag.BoolP("version", "v", false, "show current version")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		if notFoundErr, ok := err.(viper.ConfigFileNotFoundError); ok {
			msg := notFoundErr.Error()
			s := strings.Index(msg, "[")
			e := strings.Index(msg[s+1:], " ")
			fpath := msg[s+1:s+e+1] + "/" + "clash-subscription-updater.yaml"
			if _, err := os.Create(fpath); err != nil {
				log.Fatal(err)
			}
			viper.WriteConfig()
		}
	}
	if url := pflag.Arg(0); url != "" {
		viper.Set("subscription-url", url)
	}
	if viper.GetBool("override") && !viper.GetBool("help") && !viper.GetBool("version") {
		viper.Set("override", false)
		viper.WriteConfig()
	}
}

func main() {
	if viper.GetBool("help") {
		pflag.PrintDefaults()
		return
	}
	if viper.GetBool("version") {
		fmt.Printf("version: ")
		printVersion()
		fmt.Println()
		return
	}
	url := viper.GetString("subscription-url")
	if url == "" {
		log.Fatal("subscription url is required")
		pflag.PrintDefaults()
	}
	configDir := viper.GetString("clash-config-dir")
	u := updater.NewHttpUpdater(url, configDir, viper.GetInt("interval"))
	pxs := viper.Get("proxies").([]interface{})
	proxies := make([]overider.Proxy, len(pxs))
	for i, p := range pxs {
		proxy := overider.Proxy{}
		mapstructure.Decode(p, &proxy)
		proxies[i] = proxy
	}
	u.SetProxies(proxies)
	rs := viper.GetStringSlice("rules")
	rules := make([]overider.Rule, len(rs))
	for i, r := range rs {
		rules[i] = overider.Rule(r)
	}
	u.SetRules(rules)

	var task = func() {
		err := u.Update()
		if err != nil {
			log.Printf("error fetch config %s", url)
		} else {
			log.Printf("Updated to %s | patch: proxies(+%d) rules(+%d)", configDir, len(proxies), len(rules))
		}
	}
	s := gocron.NewScheduler()
	s.Every(uint64(viper.GetInt("interval"))).Minute().Do(task)
	<-s.Start()
}
