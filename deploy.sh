#!/bin/sh

set -o errexit
set -o pipefail

VENDOR=bnavetta
DRIVER=nvidiaDrivers

if [ $# -ne 1 ]; then
  echo "Usage: $0 <volume plugin directory>" >&2
  exit 1
fi

plugin_dir="$1"

driver_dir="$plugin_dir/$VENDOR${VENDOR:+"~"}${DRIVER}"
if [ ! -d "$driver_dir" ]; then
    mkdir "$driver_dir"
fi

# Use a copy then a move so the driver is installed atomically
cp "/$DRIVER" "$driver_dir/.$DRIVER"
mv -f "$driver_dir/.$DRIVER" "$driver_dir/$DRIVER"

while true; do sleep 3600; done