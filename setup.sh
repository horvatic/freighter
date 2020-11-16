#!/bin/bash

export API_KEY="123"

if [ "$1" == 'run-test-env' ]
then
  nohup bin/freighter & >/dev/null &
  sleep 5
  bin/freighterTest & >/dev/null
  echo
  echo
  echo "****Process Running****"
  ps -a | grep "freighter"
elif [ "$1" == 'run' ]
then
  nohup bin/freighter & >/dev/null &
  echo
  echo
  echo "****Process Running****"
  ps -a | grep "freighter"
fi
