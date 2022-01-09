#!/bin/bash

mkdir $RABBITMQ_HOME/logs
nohup $RABBITMQ_HOME/mate/rabbitmq_mate >>$RABBITMQ_HOME/logs/rabbitmq_mate.stdout.log 2>>$RABBITMQ_HOME/rabbitmq_mate.stderr.log

