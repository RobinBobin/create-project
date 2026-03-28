package rnpaper

import (
	"github.com/robinbobin/create-project/utils"
)

func Add() {
	utils.RunCmd("pnpm install react-native-paper")

	addBabelPlugin()
}
