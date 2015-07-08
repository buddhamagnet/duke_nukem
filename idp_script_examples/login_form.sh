#!/bin/sh
#
# -----------------------------------------------------------------------------
# This script demonstrates user login by form submit to the identity provider.
#
# Usage: ./login_form.sh <user name> <password>
#
# This script will produce a curl cookie jar with name curl-cookies.txt
# The cookie jar file will contain user session cookies if login is successful.
# -----------------------------------------------------------------------------
#
# -----------------------------------------------------------------------------
# Include server.txt file for value of server / host to call
. ./server.txt
# -----------------------------------------------------------------------------
#
# And finally the call to login
curl -c curl-cookies.txt --data "operation=Login" --data "userName=$1" --data "password=$2" ${SERVER}/graph/

