#!/bin/sh
#
# -------------------------------------------------------------------------------------
# This script demonstrates making an access decision check (check a license) on behalf of
# a user. Check license = check if the license would be granted, which does not consume it yet.
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
curl -b curl-cookies.txt -b curl-cookies.txt --data "&articleId=${ARTICLE}" --data "onBehalfOfId=${ON_BEHALF_OF}" --data "doConsume=false" ${SERVER}/authz/?${LICENSED_ITEM}

