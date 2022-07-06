#!/bin/bash

$RABBITMQ_HOME/sbin/rabbitmq-server -detached
sleep 10
rabbitmq-plugins enable rabbitmq_amqp1_0
