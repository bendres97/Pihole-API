#!/bin/bash

docker build build/ -t gobuild

docker run --rm -v $PWD:/build:z gobuild