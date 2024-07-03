#!/bin/bash

CMD="mysql -u sqlsheet -psqlsheet sqlsheet"

sudo mysql < createuser.sql
$CMD < dump.sql
