# go-edit-cmdline
## What is it?
It is really cumbersome to edit the kernel parameters in cmdline.txt via ansible with a template engine.

So this small go binary let's you allow to add, remove or edit existing paramters via cmdline flags :)

## How to use


```
# This will create a new file in tmp
# if you wish to edit it directly remove both environment values
export cmdline="/boot/cmdline.txt"
export target="/tmp/cmdline2.txt"
# - := removes paramter
# + := adds or edit parameter

# change fsck.repair to no, remove rw, add ro add non existing ip parameter with dhcp
go run main.go +fsck.repair=no -rw +ro +ip=dhcp
```


## State
I've just coded it in 2h and just tested it. Use with caution :)