package main

import (
	"github.com/scalent-io/shiprocket/pkg"
)

type DependencyOptions struct {
	ShiprocketService pkg.ShiprockertService
	Token             string
}

type ShiprocketClient struct {
	Options DependencyOptions
}

func NewShiprocketClient(Options DependencyOptions) *ShiprocketClient {
	return &ShiprocketClient{
		Options: Options}
}
