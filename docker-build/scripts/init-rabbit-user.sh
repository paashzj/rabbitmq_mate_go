#!/bin/bash

$RABBITMQ_HOME/sbin/rabbitmqctl add_user root root123
$RABBITMQ_HOME/sbin/rabbitmqctl set_permissions -p / root ".*" ".*" ".*"
$RABBITMQ_HOME/sbin/rabbitmqctl set_user_tags root administrator