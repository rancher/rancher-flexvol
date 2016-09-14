#!/bin/bash

mkdir -p /usr/libexec/kubernetes/kubelet-plugins/volume/exec/rancher-secrets
cp /usr/bin/rancher-flexvol /usr/libexec/kubernetes/kubelet-plugins/volume/exec/rancher-secrets/rancher-secrets

/bin/cat
