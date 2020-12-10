package updater

import (
	"clash-subscription-updater/overider"
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
)

type HttpUpdater struct {
	url             string
	interval        int
	dir             string
	overrideRules   []overider.Rule
	overrideProxies []overider.Proxy
}

func NewHttpUpdater(url string, dir string, interval int) HttpUpdater {
	return HttpUpdater{url: url, interval: interval, dir: dir}
}

func (u *HttpUpdater) SetRules(rules []overider.Rule) {
	u.overrideRules = rules
}

func (u *HttpUpdater) SetProxies(proxies []overider.Proxy) {
	u.overrideProxies = proxies
}

func (u *HttpUpdater) Update() error {
	resp, err := http.Get(u.url)
	if err != nil {
		return err
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var c map[string]interface{}
	yaml.Unmarshal(buf, &c)
	proxies, ok := c["proxies"].([]interface{})
	if !ok {
		return errors.New("error fetch upstream subscription url")
	}
	patchedProxies := make([]interface{}, 0, len(proxies)+len(u.overrideProxies))
	for _, p := range u.overrideProxies {
		patchedProxies = append(patchedProxies, p)
	}
	for _, p := range proxies {
		patchedProxies = append(patchedProxies, p)
	}

	rules := c["rules"].([]interface{})
	patchedRules := make([]interface{}, 0, len(rules)+len(u.overrideRules))
	for _, r := range u.overrideRules {
		patchedRules = append(patchedRules, r)
	}
	for _, r := range rules {
		patchedRules = append(patchedRules, r)
	}

	savedPath := u.dir + "/config.yaml"
	_, err = os.Stat(savedPath)
	if os.IsNotExist(err) {
		os.Create(savedPath)
	}
	f, err := os.Open(savedPath)
	defer f.Close()
	if err != nil {
		return err
	}
	c["proxies"] = patchedProxies
	c["rules"] = patchedRules
	out, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	ioutil.WriteFile(savedPath, out, 0644)
	return nil
}
