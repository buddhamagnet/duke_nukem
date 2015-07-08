#!/bin/sh
#
# -------------------------------------------------------------------------------------
# This script demonstrates making an access decision (consuming a license) on behalf of
# a user. Consuming a license means actually using and reserving it for use, where check
# license semantics means only checks if the license would be available for consumption.
#
# Usage: ./check_license.sh <item name> <article> <on behalf of profile id>
# -------------------------------------------------------------------------------------
#
# -------------------------------------------------------------------------------------
# Required arguments:
LICENSED_ITEM=$1
ARTICLE=$2
ON_BEHALF_OF=$3
# -------------------------------------------------------------------------------------
#
# -------------------------------------------------------------------------------------
#
# Include server.txt file for value of server / host to call
. ./server.txt
# -------------------------------------------------------------------------------------
#
# And finally the call to consume the license.
# NOTE: you may change the response format by adding .json or .jwt after /authz/, e.g. like this: ${SERVER}/authz/.json
curl -b curl-cookies.txt -b curl-cookies.txt --data "&articleId=${ARTICLE}" --data "onBehalfOfId=${ON_BEHALF_OF}" ${SERVER}/authz/?${LICENSED_ITEM}

