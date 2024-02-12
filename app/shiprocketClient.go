package app

import (
	"github.com/scalent-io/shiprocket/pkg"
)

type DependencyOptions struct {
	ShiprocketService pkg.ShiprockertService
}

type ShiprocketClient struct {
	Options DependencyOptions
	Token   string
}

func NewShiprocketClient(Options DependencyOptions) *ShiprocketClient {
	return &ShiprocketClient{
		Options: Options}
}
