#!/bin/sh
#
# -------------------------------------------------------------------------------------------------------------------------
# This script demonstrates reading "me" Profile object from the identity provider. Authentication and corresponding profile
# identification relies on an authenticated user session, which is created by a executing login-form.sh first.
#
# Usage: ./cookie_auth_query_me.sh
#
# NOTE: for success case test you must run login-form.sh with successful login first.
#       that will produce a curl cookie jar with name curl-cookies.txt
# -------------------------------------------------------------------------------------------------------------------------
#
# -------------------------------------------------------------------------------------------------------------------------
# Include server.txt file for value of server / host to call
. ./server.txt
# -------------------------------------------------------------------------------------------------------------------------
#
curl -b curl-cookies.txt ${SERVER}/graph/me

