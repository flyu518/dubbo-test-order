#! /bin/bash

# 需要执行：source export.sh，而不是 ./export.sh，直接执行是子进程中执行，环境变量不会生效！！！

NACOS_CONFIG="{\"address\":\"127.0.0.1:8848\",\"username\":\"\",\"password\":\"\"}"
CENTER_CONFIG="{\"namespace\":\"dev\",\"group\":\"order\",\"data_id\":\"server.yaml\"}"
CENTER_CONFIG_CLIENT="{\"namespace\":\"dev\",\"group\":\"order\",\"data_id\":\"client.yaml\"}"

# 设置环境变量
export DUBBOGO_NACOS_CONFIG="$NACOS_CONFIG"
export DUBBOGO_CENTER_CONFIG="$CENTER_CONFIG"
export DUBBOGO_CENTER_CONFIG_CLIENT="$CENTER_CONFIG_CLIENT"

echo "DUBBOGO_NACOS_CONFIG: $DUBBOGO_NACOS_CONFIG"
echo "DUBBOGO_CENTER_CONFIG: $DUBBOGO_CENTER_CONFIG"
echo "DUBBOGO_CENTER_CONFIG_CLIENT: $DUBBOGO_CENTER_CONFIG_CLIENT"
