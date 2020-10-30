// NPi i.MX6ULL support for tamago/arm
// https://github.com/f-secure-foundry/tamago
//
// Copyright (c) F-Secure Corporation
// https://foundry.f-secure.com
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

package npi

import (
	_ "unsafe"

	"github.com/f-secure-foundry/tamago/soc/imx6"
)

// On the NPi i.MX6ULL the serial console is UART1, therefore standard
// output is redirected there.

//go:linkname printk runtime.printk
func printk(c byte) {
	imx6.UART1.Tx(c)
}
