#!/bin/bash

if ! which conversion-gen &>/dev/null; then
    # TODO automatically install conversion-gen
    echo "Install conversion-gen from https://github.com/kubernetes/code-generator first."
    exit 1
fi

conversion-gen --input-dirs github.com/coreos/prometheus-operator/pkg/client/monitoring/v1,github.com/coreos/prometheus-operator/pkg/client/monitoring/v2 -O zz_generated.conversion
