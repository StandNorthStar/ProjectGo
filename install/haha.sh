#!/bin/bash

curl -XPOST -H "content-type:application/json" http://12.1.216.132:8012/common/loan/v1/login -d '
{
  "userName":"2651888",
  "password":"123456",
  "flag":"loan"
}'
