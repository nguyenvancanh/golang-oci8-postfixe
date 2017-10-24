#!/bin/bash
set -e

service postfix start

exec $@
