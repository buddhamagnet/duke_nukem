#!/bin/sh
#
# -------------------------------------------------------------------------------------
# This example shows how to initialize licenses for an individual user. License initialization requires
# caller to define the product packgage name, and profile id.
# 
# Usage: ./initialize_personal_licenses.sh <product package name> <profile id>
# -------------------------------------------------------------------------------------
#
# -------------------------------------------------------------------------------------
# Required arguments
PRODUCT_PACKAGE_NAME=$1
PROFILE_ID=$2
# -------------------------------------------------------------------------------------
#
# -------------------------------------------------------------------------------------
# Include server.txt file for value of server / host to call
. ./server.txt
# -------------------------------------------------------------------------------------
#
# And finally the call to grant licenses to content of the
# product package for the user identified by PROFILE_ID
curl -b curl-cookies.txt "${SERVER}/graph/ProductPackage\[@title%3D'${PRODUCT_PACKAGE_NAME}'\]?operation=InitializePersonalLicenses&initializeForProfileId=${PROFILE_ID}"
