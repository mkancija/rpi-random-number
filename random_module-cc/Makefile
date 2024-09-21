CR_C := arm-linux-gnueabihf-
MODULES := /media/kanc/20210324_4TB1/Software/raspberry_pi/linux/
obj-m := random_module.o

all:
	make ARCH=arm CROSS_COMPILE=$(CR_C) -C $(MODULES) M=$(shell pwd) modules

clean:
	make ARCH=arm CROSS_COMPILE=$(CR_C) -C $(MODULES) M=$(shell pwd) clean
