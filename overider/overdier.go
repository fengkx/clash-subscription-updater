package overider

type Rule string
type Overrider interface {
	OverrideProxy(original []Proxy, patch []Proxy) []Proxy
	OverrideRule(original []Rule, patch []Rule) []Rule
}

type ClashOverrider struct{}

func (o ClashOverrider) OverrideRule(original []Rule, patch []Rule) []Rule {
	panic("implement me")
}

func (o ClashOverrider) OverrideProxy(original []Proxy, patch []Proxy) []Proxy {
	panic("implement me")
}
