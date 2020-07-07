// Code generated by "mapstructure-to-hcl2 -type IAPConfig"; DO NOT EDIT.
package googlecompute

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatIAPConfig is an auto-generated flat version of IAPConfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatIAPConfig struct {
	IAP                    *bool   `mapstructure:"use_iap" required:"false" cty:"use_iap" hcl:"use_iap"`
	IAPLocalhostPort       *int    `mapstructure:"iap_localhost_port" cty:"iap_localhost_port" hcl:"iap_localhost_port"`
	IAPHashBang            *string `mapstructure:"iap_hashbang" required:"false" cty:"iap_hashbang" hcl:"iap_hashbang"`
	IAPExt                 *string `mapstructure:"iap_ext" required:"false" cty:"iap_ext" hcl:"iap_ext"`
	IAPTunnelLaunchTimeout *int    `mapstructure:"iap_tunnel_launch_timeout" required:"false" cty:"iap_tunnel_launch_timeout" hcl:"iap_tunnel_launch_timeout"`
}

// FlatMapstructure returns a new FlatIAPConfig.
// FlatIAPConfig is an auto-generated flat version of IAPConfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*IAPConfig) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatIAPConfig)
}

// HCL2Spec returns the hcl spec of a IAPConfig.
// This spec is used by HCL to read the fields of IAPConfig.
// The decoded values from this spec will then be applied to a FlatIAPConfig.
func (*FlatIAPConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"use_iap":                   &hcldec.AttrSpec{Name: "use_iap", Type: cty.Bool, Required: false},
		"iap_localhost_port":        &hcldec.AttrSpec{Name: "iap_localhost_port", Type: cty.Number, Required: false},
		"iap_hashbang":              &hcldec.AttrSpec{Name: "iap_hashbang", Type: cty.String, Required: false},
		"iap_ext":                   &hcldec.AttrSpec{Name: "iap_ext", Type: cty.String, Required: false},
		"iap_tunnel_launch_timeout": &hcldec.AttrSpec{Name: "iap_tunnel_launch_timeout", Type: cty.Number, Required: false},
	}
	return s
}
