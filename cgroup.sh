#!/bin/bash

## make a new cgroup file system
sudo mkdir cgroup_test
sudo mount -t cgroup -o none,name=cgroup_test cgroup_test ./cgroup_test
## this operation will create a root file system of cgroup
## create new sub-filesystem from this cgroup
sudo mkdir cgroup-1
sudo mkdir cgroup-2

## use default hierarchy
/sys/fs/cgroup/
