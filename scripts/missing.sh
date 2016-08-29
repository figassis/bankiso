#!/bin/bash

find ../iso20022 -type f | xargs grep -H -c 'String() (result string, ok bool)' | grep 0$ | cut -d':' -f1 > missing.txt