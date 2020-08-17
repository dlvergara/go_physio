package provider

type Provider interface {
	GetName() string
	GetData() string
}