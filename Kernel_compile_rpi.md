

# Cross compile raspberry pi kernel

## Prerequisites

 - Raspberry pi device
 - Empty SD card
 - Preflashed Raspberry Pi OS
   - https://www.raspberrypi.com/software/operating-systems/
   - https://projects.raspberrypi.org/en/projects/raspberry-pi-setting-up/2

## Destination device

 - Raspberry Pi Zero W
 - SD card: SanDisk 32GB 
 - Raspberry Pi OS (Legacy) Lite 32bit
   - https://downloads.raspberrypi.com/raspios_oldstable_full_armhf/release_notes.txt

## Download kernel source

https://www.raspberrypi.com/documentation/computers/linux_kernel.html

List of available raspberry pi kernel branches:

 - https://github.com/raspberrypi/linux


Prepare and clone raspberry pi kernel:

``` bash
cd /media/kanc/20210324_4TB1/Software/raspberry_pi/

git clone --depth=1 https://github.com/raspberrypi/linux
```

To cross compile raspberry pi kernel (debian) we need suatible host machine (debian, or ubuntu).
For this assigment we will use Ubuntu 

```bash
uname -a
```

`Linux kanc-home 6.8.0-40-generic #40~22.04.3-Ubuntu SMP PREEMPT_DYNAMIC Tue Jul 30 17:30:19 UTC 2 x86_64 x86_64 x86_64 GNU/Linux`

## Install host machine dependencies

```bash
sudo apt install bc bison flex libssl-dev make libc6-dev libncurses5-dev
```

Since our device is Rpi Zero W, which is 32bit, we will use 32bit toolchain

 - Device list and architecture is available at official raspberry pi site: https://www.raspberrypi.com/documentation/computers/linux_kernel.html#cross-compile-the-kernel

```bash
sudo apt install crossbuild-essential-armhf
```

Make kernel sources

```bash

cd /media/kanc/20210324_4TB1/Software/raspberry_pi/linux

make ARCH=arm CROSS_COMPILE=arm-linux-gnueabihf- bcmrpi_defconfig
```


Adjust current/local kernel version:

Locate `CONFIG_LOCALVERSION` variable.

```bash
nano .config

CONFIG_LOCALVERSION="rpi-zero-w-kanc-v1"
```

Build

Check current host machine 

 - CPU core count
```bash
nproc
```


```bash
make -j16 ARCH=arm CROSS_COMPILE=arm-linux-gnueabihf- zImage modules dtbs
```

You can compile modules separetaly:

 - zImage
 - modules
 - dtbs

(Sometimes for different device i got error for different reasons.)

```bash
make -j16 ARCH=arm CROSS_COMPILE=arm-linux-gnueabihf- zImage
make -j16 ARCH=arm CROSS_COMPILE=arm-linux-gnueabihf- modules
make -j16 ARCH=arm CROSS_COMPILE=arm-linux-gnueabihf- dtbs
```

First two steps will take from minute to five minutes depending of a system you are compiling on.

## Prepare kernel files

Create separate directory that will hold all relevant kernel files.

```bash
mkdir /media/kanc/20210324_4TB1/Software/raspberry_pi/kernel-prepared

mkdir /media/kanc/20210324_4TB1/Software/raspberry_pi/kernel-prepared/boot
mkdir /media/kanc/20210324_4TB1/Software/raspberry_pi/kernel-prepared/boot/overlays
```

### Copy boot image

Copy previously built zImage boot image file to boot dir.
```
cp /media/kanc/20210324_4TB1/Software/raspberry_pi/linux/arch/arm/boot/zImage /media/kanc/20210324_4TB1/Software/raspberry_pi/kernel-prepared/boot/kernel-k1.img
```
### Copy device tree blobs

Following the instructions on RaspberryPi site, for ARM 32bit, and kernel > 6.5,
copy all compiled device tree files from broadcom folder.

```bash
cd /media/kanc/20210324_4TB1/Software/raspberry_pi2/linux/arch/arm/boot/dts/broadcom 
cp *.dtb /media/kanc/20210324_4TB1/Software/raspberry_pi2/kernel-prepared/boot/
```
### Copy overlays

Copy all DTBO files to overlays dir.

```bash
cd /media/kanc/20210324_4TB1/Software/raspberry_pi2/linux/arch/arm/boot/dts/overlays
cp *.dtb* /media/kanc/20210324_4TB1/Software/raspberry_pi2/kernel-prepared/boot/overlays/
```

### Install modules

Lib dir
From kernel linux dir:

```bash
make -j16 ARCH=arm CROSS_COMPILE=arm-linux-gnueabihf- INSTALL_MOD_PATH=/media/kanc/20210324_4TB1/Software/raspberry_pi2/kernel-prepared/ modules_install
```

## Prepare and Copy kernel files to Raspberry pi.

### Prepare folder 

On your raspberry pi prepare folder:

```bash
mkdir rpi-kernel-001; cd rpi-kernel-001
kanc@raspberrypi:~/rpi-kernel-001
```
From prepared kernel dir remove build dir

```bash
rm -r /media/kanc/20210324_4TB1/Software/raspberry_pi2/kernel-prepared/lib/modules/6.6.51rpi-zero-w-kanc-v1+/build
```
### Copy kernel files

From host machine copy prepared kernel files:

```bash
cd /media/kanc/20210324_4TB1/Software/raspberry_pi2/kernel-prepared
scp -r * kanc@192.168.53.121:/home/kanc/rpi-kernel-001/
```

### Backup current rpi kernel

```bash
mkdir /home/kanc/bkp_kernel-6.1.21
cp -r /boot /home/kanc/bkp_kernel-6.1.21/

# lib dir is actualy symlink of a /usr/lib dir.
cp -r /usr/lib /home/kanc/bkp_kernel-6.1.21/

```

### Check current kernel verison

`kanc@raspberrypi:~/rpi-kernel-001 $ uname -a`

`Linux raspberrypi 6.1.21+ #1642 Mon Apr  3 17:19:14 BST 2023 armv6l GNU/Linux`


### Replace current rpi kernel

Replace boot, add kernel modules.

```bash
cd /home/kanc/rpi-kernel-001/
sudo cp -r /home/kanc/rpi-kernel-001/boot/* /boot

sudo cp -r /home/kanc/rpi-kernel-001/lib/modules/6.6.51rpi-zero-w-kanc-v1+/ /lib/modules

sudo shutdown -r 0
```

After restart, check kernel version

```bash
$ uname -a
```

`Linux raspberrypi 6.6.51rpi-zero-w-kanc-v1+ #1 Wed Sep 18 22:29:38 CEST 2024 armv6l GNU/Linux`











## Extras

RPI kernel offers GUI menuconfig tool
https://www.raspberrypi.com/documentation/computers/linux_kernel.html#configure-the-kernel

``` bash
make ARCH=arm CROSS_COMPILE=arm-linux-gnueabihf- menuconfig
```
