#!/bin/bash

gogio -target android -version $1 -minsdk 29 .
apksigner sign --ks ~/.android/sign.keystore treman.apk