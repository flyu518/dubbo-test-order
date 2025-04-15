#! /bin/bash

# 需要执行：source export.sh，而不是 ./export.sh，直接执行是子进程中执行，环境变量不会生效！！！

LM_CENTER_CONFIG="{\"address\":\"127.0.0.1:8848\",\"username\":\"\",\"password\":\"\",\"namespace\":\"dev\",\"group\":\"lingmou\",\"data_id\":\"order-server.yaml\"}"

# 设置环境变量
export LM_CENTER_CONFIG="$LM_CENTER_CONFIG"

echo "LM_CENTER_CONFIG: $LM_CENTER_CONFIG"
