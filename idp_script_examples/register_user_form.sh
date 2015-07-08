#!/bin/sh
#
# ---------------------------------------------------------------------------
# This example registers a new user with the identity provider.
#
# Usage: ./register_user_form.sh <first name> <last name> <email> <password>
# ---------------------------------------------------------------------------
#
# ---------------------------------------------------------------------------
# The Curl form submit below is ~equivalent to the following HTML form:
#
# <form method="POST" id="registerForm" action="/register/registerHandler.vsl">
#    <input name="operation" value="RegisterUser">
#    <input name="/Profile[@id='{randomUuid,profile}']/@id" value="{randomUuid,profile}">
#    <input name="/Profile[@id='{randomUuid,profile}']/~ManyToOne/Person[@id='{randomUuid,person}']/@firstName" type="text" value="">
#    <input name="/Profile[@id='{randomUuid,profile}']/~ManyToOne/Person[@id='{randomUuid,person}']/@lastName" type="text" value="">
#    <input name="email" type="text" value="" oninvalid="setCustomValidity('Please enter a valid email');">
#    <input name="/Profile[@id='{randomUuid,profile}']/~OneToMany/ContactInformation[@id='{randomUuid,contactInformation}']/~OneToMany/EmailAddress[@id='{randomUuid,emailAddress}']/@value" value="{$email}">
#    <input name="/Profile[@id='{randomUuid,profile}']/~ManyToOne/Person[@id='{randomUuid,person}']/~OneToMany/EmailAndPassword[@id='{randomUuid,emailAndPassword}']/@userName" value="{$email}">
#    <input type="password" name="/Profile[@id='{randomUuid,profile}']/~ManyToOne/Person[@id='{randomUuid,person}']/~OneToMany/EmailAndPassword[@id='{randomUuid,emailAndPassword}']/@passwordPlain" value="">
#    <input type="password" name="confirmedPassword" value="">
#    <input name="/Profile[@id='{randomUuid,profile}']/~OneToMany/Account[@id='{randomUuid,account}']/@id" value="{randomUuid,account}">
#    <input name="/Profile[@id='{randomUuid,profile}']/~OneToMany/Account[@id='{randomUuid,account}']/@accountType" value="personal">
#    <input type="checkbox" name="acceptsTsAndCs" value="true">
# </form>
# ---------------------------------------------------------------------------
#
#
# ---------------------------------------------------------------------------
# Include server.txt file for value of server / host to call
. ./server.txt
# ---------------------------------------------------------------------------
#
# And finally the call to register a user
curl --data "operation=RegisterUser" --data-urlencode "/Profile[@id%3D'{randomUuid, profile}']/@id%3D{randomUuid, profile}" --data-urlencode "/Profile[@id%3D'{randomUuid, profile}']/~ManyToOne/Person[@id%3D'{randomUuid, person}']/@firstName=$1" --data-urlencode "/Profile[@id%3D'{randomUuid, profile}']/~ManyToOne/Person[@id%3D'{randomUuid, person}']/@lastName=$2" --data-urlencode "Profile[@id%3D'{randomUuid, profile}']/~OneToMany/ContactInformation/~OneToMany/EmailAddress/@value%3D$3" --data-urlencode "/Profile[@id%3D'{randomUuid, profile}']/~ManyToOne/Person[@id%3D'{randomUuid, person}']/~OneToMany/EmailAndPassword[@id%3D'{randomUuid, emailAndPassword}']/@userName=$3" --data-urlencode "/Profile[@id%3D'{randomUuid, profile}']/~ManyToOne/Person[@id%3D'{randomUuid, person}']/~OneToMany/EmailAndPassword[@id%3D'{randomUuid, emailAndPassword}']/@passwordPlain=$4" --data-urlencode "/Profile[@id%3D'{randomUuid,profile}']/~OneToMany/Account[@id%3D'{randomUuid,account}']/@id%3D{randomUuid,account}" --data-urlencode "/Profile[@id%3D'{randomUuid,profile}']/~OneToMany/Account[@id%3D'{randomUuid,account}']/@accountType=personal" --data "confirmedPassword=$4" --data "acceptsTsAndCs=true" --data "accountType=personal" ${SERVER}/graph/

# ECONOMIST NOTES: this beast is what comes back as a successful registration: https://gist.github.com/buddhamagnet/e4f18b2356f8724b2862