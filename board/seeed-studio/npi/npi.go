// NPi i.MX6ULL support for tamago/arm
// https://github.com/f-secure-foundry/tamago
//
// Copyright (c) F-Secure Corporation
// https://foundry.f-secure.com
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

// Package mx6ullevk provides hardware initialization, automatically on import,
// for the NXP MCIMX6ULL-EVK evaluation board.
//
// This package is only meant to be used with `GOOS=tamago GOARCH=arm` as
// supported by the TamaGo framework for bare metal Go on ARM SoCs, see
// https://github.com/f-secure-foundry/tamago.
package npi

import (
	_ "unsafe"

	"github.com/f-secure-foundry/tamago/soc/imx6"
	_ "github.com/f-secure-foundry/tamago/soc/imx6/imx6ul"
)

// Init takes care of the lower level SoC initialization triggered early in
// runtime setup.
//
//go:linkname Init runtime.hwinit
func Init() {
	imx6.Init()

	// initialize console
	imx6.UART1.Init()
}
