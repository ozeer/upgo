#!/bin/bash

# 检查的 YAML 文件路径和名称
YAML_FILE="conf/config.yml"

# 获取标签对应的版本号
LATEST_TAG=$(git describe --tags --abbrev=0)

# 获取 YAML 文件中的 app.version
YAML_VERSION=$(awk '/app:/ {flag=1;next} flag && /version:/ {print $2; exit}' "$YAML_FILE")

echo "Latest tag: $LATEST_TAG"
echo "Yaml tag: $YAML_VERSION"

# 检查版本号是否与标签一致
if [[ "$LATEST_TAG" == "$YAML_VERSION" ]]; then
  echo "YAML 文件中的 app.version 与标签一致，可以推送标签。"
  exit 0
else
  echo "YAML 文件中的 app.version 与标签不一致，请确保文件中的版本号与标签对应。"
  exit 1
fi
