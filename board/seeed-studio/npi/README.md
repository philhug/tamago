TamaGo - bare metal Go for ARM SoCs - NPI i.MX6ULL support
==========================================================

tamago | https://github.com/f-secure-foundry/tamago  

Copyright (c) F-Secure Corporation  
https://foundry.f-secure.com

![TamaGo gopher](https://github.com/f-secure-foundry/tamago/wiki/images/tamago.svg?sanitize=true)

Authors
=======

Andrea Barisani  
andrea.barisani@f-secure.com | andrea@inversepath.com  

Andrej Rosano  
andrej.rosano@f-secure.com   | andrej@inversepath.com  

Introduction
============

TamaGo is a framework that enables compilation and execution of unencumbered Go
applications on bare metal ARM System-on-Chip (SoC) components.

The [npi](https://github.com/f-secure-foundry/tamago/tree/master/board/seeed-studio/npi)
package provides support for the [NPi i.MX6ULL](https://wiki.seeedstudio.com/NPi-i.MX6ULL-Dev-Board-Linux-SBC/) board.

Documentation
=============

For more information about TamaGo see its
[repository](https://github.com/f-secure-foundry/tamago) and
[project wiki](https://github.com/f-secure-foundry/tamago/wiki).

For the underlying driver support for this board see package
[imx6](https://github.com/f-secure-foundry/tamago/tree/master/soc/imx6).

The package API documentation can be found on
[pkg.go.dev](https://pkg.go.dev/github.com/f-secure-foundry/tamago).

Supported hardware
==================

| SoC           | Board                                                                  | SoC package                                                         | Board package                                                                                   |
|---------------|------------------------------------------------------------------------|---------------------------------------------------------------------|-------------------------------------------------------------------------------------------------|
| NXP i.MX6ULL  | [NPi i.MX6ULL eMMC](https://wiki.seeedstudio.com/NPi-i.MX6ULL-Dev-Board-Linux-SBC/) | [imx6](https://github.com/f-secure-foundry/tamago/tree/master/soc/imx6) | [mx6ullevk](https://github.com/f-secure-foundry/tamago/tree/master/board/nxp/mx6ullevk) |
| NXP i.MX6ULL  | [NPi i.MX6ULL NAND](https://wiki.seeedstudio.com/NPi-i.MX6ULL-Dev-Board-Linux-SBC/) | [imx6](https://github.com/f-secure-foundry/tamago/tree/master/soc/imx6) | [mx6ullevk](https://github.com/f-secure-foundry/tamago/tree/master/board/nxp/mx6ullevk) |

Compiling
=========

Go applications are simply required to import, the relevant board package to
ensure that hardware initialization and runtime support takes place:

```golang
import (
	_ "github.com/f-secure-foundry/tamago/board/seeed-studio/npi"
)
```

Build the [TamaGo compiler](https://github.com/f-secure-foundry/tamago-go)
(or use the [latest binary release](https://github.com/f-secure-foundry/tamago-go/releases/latest)):

```
git clone https://github.com/f-secure-foundry/tamago-go -b latest
cd tamago-go/src && ./all.bash
cd ../bin && export TAMAGO=`pwd`/go
```

Go applications can be compiled as usual, using the compiler built in the
previous step, but with the addition of the following flags/variables and
ensuring that the required SoC and board packages are available in `GOPATH`:

```
GO_EXTLINK_ENABLED=0 CGO_ENABLED=0 GOOS=tamago GOARM=7 GOARCH=arm \
  ${TAMAGO} build -ldflags "-T 0x80010000  -E _rt0_arm_tamago -R 0x1000"
```

An example application, targeting the NPi i.MX6ULL platform,
is [available](https://github.com/f-secure-foundry/tamago-example).

Executing and debugging
=======================

The [example application](https://github.com/f-secure-foundry/tamago-example)
provides reference usage and a Makefile target for automatic creation of an ELF
as well as `imx` image for flashing.

Native hardware: imx image on microSD
-------------------------------------

Copy the built `imx` image to a microSD as follows:

```
sudo dd if=<path to imx file> of=/dev/sdX bs=1M conv=fsync
```
*IMPORTANT*: /dev/sdX must be replaced with your microSD device (not eventual
microSD partitions), ensure that you are specifying the correct one. Errors in
target specification will result in disk corruption.

Ensure the `SW5` boot switches on the processor board are set to microSD boot
mode:

| switch | position |
|--------|----------|
| 1      | OFF      |
| 2      | ON       |


Native hardware: existing bootloader
------------------------------------

Copy the built ELF binary on an external microSD card then launch it from the
U-Boot console as follows:

```
ext2load mmc 1:1 0x90000000 example
bootelf -p 0x90000000
```

For non-interactive execution modify the U-Boot configuration accordingly.

Standard output
---------------

The standard output can be accessed through the debug console found
on micro USB connector J1901 and the following `picocom` configuration:

```
picocom -b 115200 -eb /dev/ttyUSB0 --imap lfcrlf
```

Debugging
---------

The application can be debugged with GDB over JTAG using `openocd` and the
`imx6ull.cfg` and `gdbinit` debugging helpers published
[here](https://github.com/f-secure-foundry/tamago/tree/master/_dev).

```
# start openocd daemon
openocd -f interface/ftdi/jtagkey.cfg -f imx6ull.cfg

# connect to the OpenOCD command line
telnet localhost 4444

# debug with GDB
arm-none-eabi-gdb -x gdbinit example
```

Hardware breakpoints can be set in the usual way:

```
hb ecdsa.Verify
continue
```

License
=======

tamago | https://github.com/f-secure-foundry/tamago  
Copyright (c) F-Secure Corporation

This program is free software: you can redistribute it and/or modify it under
the terms of the GNU General Public License as published by the Free Software
Foundation under version 3 of the License.

This program is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE. See the GNU General Public License for more details.

See accompanying LICENSE file for full details.
