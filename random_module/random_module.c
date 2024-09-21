#include <linux/init.h>
#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/fs.h>
#include <linux/uaccess.h>

#define DEVICE_NAME "random_module"

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Kanc");
MODULE_DESCRIPTION("A simple random number generator module");

static int dev_open(struct inode*, struct file*);
static int dev_release(struct inode*, struct file*);
static ssize_t dev_read(struct file*, char*, size_t, loff_t*);
static ssize_t dev_write(struct file*, const char*, size_t, loff_t*);

static struct file_operations fops = {
   .open = dev_open,
   .read = dev_read,
   .write = dev_write,
   .release = dev_release,
};

static int major;

static int __init random_number_init(void) {
    major = register_chrdev(0, DEVICE_NAME, &fops);

    if (major < 0) {
        printk(KERN_ALERT "random_module: load failed\n");
        return major;
    }

    printk(KERN_INFO "random_module: module has been loaded: %d\n", major);
    return 0;
}

static void __exit random_number_exit(void)
{
    unregister_chrdev(major, DEVICE_NAME);
    printk(KERN_INFO "random_module: exit");
}

static int dev_open(struct inode *inodep, struct file *filep) {
   printk(KERN_INFO "random_module: device opened\n");

   return 0;
}

static ssize_t dev_write(struct file *filep, const char *buffer,
                         size_t len, loff_t *offset) {

   printk(KERN_INFO "random_module: read only\n");
   return -EFAULT;
}

static int dev_release(struct inode *inodep, struct file *filep) {
   printk(KERN_INFO "random_module: device closed\n");

   return 0;
}

static ssize_t dev_read(struct file *filep, char *buffer, size_t len, loff_t *offset) {

    // unsigned int seed = ktime_get_real_ns(); //time(NULL);
    unsigned int number = 0;
    char *message = "1234567890";
    int errors;
    int message_length;

    get_random_bytes(&number, sizeof(number)-1);
    sprintf(message, "%u\n", number);

    message_length = strlen(message);

    errors = copy_to_user(buffer, message, message_length);

    return errors == 0 ? message_length : -EFAULT;
}

// register init and exit functions
module_init(random_number_init);
module_exit(random_number_exit);
