# my-runc
demo for runc

## How To Use
1. build this project first
```shell
sh build.sh
```
the you will get an executable binary file:main
2. exec the main file 
```shell
sudo ./main run -ti /bin/bash
```
## Some Issue
if everything work well,you will start a shell terminal.
this terminal can be considered a process with new namespaces,
you can check namespace info at /proc filesystem.But there are some
unexpected places,for example,if u use "ps -ef",you will get all process
info from default pid namespace. The reason for this problem is that we use
mount command to mount the proc filesystem for  this process. I am trying to fix it.
