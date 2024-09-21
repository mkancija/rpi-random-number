KERNELVER := $(shell uname -r)
KERNELDIR := /lib/modules/$(KERNELVER)/build

obj-m += random_module.o

all:
    make -C $(KERNELDIR) M=$(PWD) modules


clean:
    make -C $(KERNELDIR) M=$(PWD) clean