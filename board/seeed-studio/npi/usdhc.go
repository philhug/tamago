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
	"github.com/f-secure-foundry/tamago/soc/imx6"
	"github.com/f-secure-foundry/tamago/soc/imx6/usdhc"
)

// SD1 is the base board full size SD instance
var SD1 = usdhc.USDHC1

// SD2 is the CPU board microSD instance
var SD2 = usdhc.USDHC2

// SD1/SD2 configuration constants.
//
// On the NPi i.MX6ULL the following uSDHC interfaces are connected:
//   * uSDHC1: base board full size SD slot (SD1)
//   * uSDHC2: CPU board eMMC (SD2)
const (
	IOMUXC_SW_MUX_CTL_PAD_CSI_DATA04 = 0x020e01f4
	IOMUXC_SW_PAD_CTL_PAD_CSI_DATA04 = 0x020e0480
	IOMUXC_USDHC1_WP_SELECT_INPUT    = 0x020e066c

	USDHC1_WP_MODE   = 8
	DAISY_CSI_DATA04 = 0b10

	IOMUXC_SW_MUX_CTL_PAD_CSI_PIXCLK = 0x020e01d8
	IOMUXC_SW_PAD_CTL_PAD_CSI_PIXCLK = 0x020e0464
	IOMUXC_USDHC2_WP_SELECT_INPUT    = 0x020e069c

	USDHC2_WP_MODE   = 1
	DAISY_CSI_PIXCLK = 0b10

	SD1_BUS_WIDTH = 4
	SD2_BUS_WIDTH = 4
)

func init() {
	var err error

	// There are no write-protect lines on uSD cards. The write-protect
	// line on the full size slot is not connected. Therefore the
	// respective SoC pads must be selected on pulled down unconnected pads
	// to ensure the driver never see write protection enabled.
	ctl := uint32((1 << imx6.SW_PAD_CTL_PUE) | (1 << imx6.SW_PAD_CTL_PKE))

	// SD write protect (USDHC1_WP)
	wpSD1, err := imx6.NewPad(IOMUXC_SW_MUX_CTL_PAD_CSI_DATA04,
		IOMUXC_SW_PAD_CTL_PAD_CSI_DATA04,
		IOMUXC_USDHC1_WP_SELECT_INPUT)

	if err != nil {
		panic(err)
	}

	// SD2 write protect (USDHC2_WP)
	wpSD2, err := imx6.NewPad(IOMUXC_SW_MUX_CTL_PAD_CSI_PIXCLK,
		IOMUXC_SW_PAD_CTL_PAD_CSI_PIXCLK,
		IOMUXC_USDHC2_WP_SELECT_INPUT)

	if err != nil {
		panic(err)
	}

	if !imx6.Native {
		return
	}

	wpSD1.Mode(USDHC1_WP_MODE)
	wpSD1.Select(DAISY_CSI_DATA04)
	wpSD1.Ctl(ctl)

	wpSD2.Mode(USDHC2_WP_MODE)
	wpSD2.Select(DAISY_CSI_PIXCLK)
	wpSD2.Ctl(ctl)

	SD1.Init(SD1_BUS_WIDTH)
	SD2.Init(SD2_BUS_WIDTH)

	// Only SD1 supports 1.8V switching on this board.
	SD1.LowVoltage = func() bool {
		// No actual function is required as VEND_SPEC_VSELECT, already
		// set by the usdhc driver, is used on this board circuitry to
		// switch to LV.
		return true
	}
}
