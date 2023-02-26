#!/bin/bash
# demo for aufs
## mkdir some dirs and use mount
sudo mount -t aufs -o dirs=./container_layer:./image_layer1:./image_layer2:./image_layer3:./image_layer4 none ./mnt