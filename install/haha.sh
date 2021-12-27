#!/bin/bash

curl -XPOST -H "content-type:application/json" http://x.x.xxx.xxx:xxxx/common/loan/v1/login -d '
{
  "userName":"xxxxx",
  "password":"xxxxx",
  "flag":"loan"
}'
