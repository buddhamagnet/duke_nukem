Hi,

Here are some tips on use cases you mention in your email. The following HTTP calls use the demo system we have setup for you. The demo system has license models configured for authenticated users and guest users.

In the demo system, when granting licenses (entitlement records), you will essentially grant rights to an asset called "NewsItem" from product package "pkg-1", which you will be able to see in some of the calls below. You will also find some further notes and comments below the examples:

Step #1: "login"

 * Do either an "OAuth login" or "direct login"
 * Direct login applies only to a subset of calls, which you can "justify" making without involving the knowledge of the user and the consumer application.
 * Direct login means that a user agent call the IdP login service explicitly and gets an authenticated session with the IdP. Note that this does not mean Single Sign-on.
 * You can get a user name - password pair at: http://econ-idp-poc.elasticbeanstalk.com/ --> register. For the purpose of the demo there are no enforcing permissions in play and you will be able to make any calls.

Example call for direct login:
http://econ-idp-poc.elasticbeanstalk.com/graph?operation=Login&userName=...&password=...

 * from the response you should keep cookies and check the response JSON to look like this:
{
"__objType": "LoginResult",
"resultCode": "Success",
"_loginName": "your-login-name@...com"
}

 * the recommended way is that you convert the above example to a HTTP POST with the parameters in the payload instead of the URL itself

OAuth flow summary
1. get OAuth request token from IdP at: http://econ-idp-poc.elasticbeanstalk.com/oauth/
2. redirect user agent to IdP authorization URL using the request token from step 1: http://econ-idp-poc.elasticbeanstalk.com/oauth?oauth_token=...
3. Handle OAuth result when the IdP redirects the user agent back to your service. At this stage the user is signed in with the IdP and you can equally grant an authenticated session to the user at your end.
4. Store the OAuth access token so that you may make OAuth signed requests on behalf of the user to the IdP / Entitlement service.
 * specific implementation details are dictated by the platform you use
 * The 10Duke SDK and our application stack includes this capability if you would consider it as a choice for implementing the components you are planning.


Use case #1 - Store entitlement record = grant licenses in 10Duke Entitlement service jargon

Example call:

http://econ-idp-poc.elasticbeanstalk.com/graph/ProductPackage[@title='pkg-1']?operation=InitializePersonalLicenses&initializeForProfileId=<id-of-profile-in-idp-you-are-granting-for>

Notes:
* If you omit parameter initializeForProfileId=<id-of-profile-in-idp-you-are-granting-for> --> then you are granting for the user logged in at the user agent your are calling from.
* If you have an OAuth signed request and omit initializeForProfileId=<id-of-profile-in-idp-you-are-granting-for> --> then you are granting for the user identity resolved from the OAuth session.

* You should get back a JSON object, which on first level looks like this:
{
"__objType": "InitializePersonalLicensesResult",
"initializeForProfile": {},
"resultCode": "Success",
"objects": {}
}

* keep an eye out for "resultCode":"Success"... anything else means doing further error handling
* the JSON is configured for maximum verbose level and you may study that freely.


Use case #2 - Call for an access decision

Example call:
http://econ-idp-poc.elasticbeanstalk.com/authz/?NewsItem
the same request asking for a JSON web token:
http://econ-idp-poc.elasticbeanstalk.com/authz/.jwt?NewsItem

Notes:
* These calls always apply for the logged in user, which means:
 a) with OAuth signed call --> applies for the user identity resolved by the OAuth session
 b) direct call (no OAUth) --> applies for the session identified by cookies used between the calling user agent and the service
* Of these two a) tends to be the correct one logically and b) in cases where server application would make a license check e.g. when starting.
* case a) requires that the OAuth login (SSO) has been implemented first

Other notes:
* you may freely study this entire flow in action by using e.g. Chrome's developer toolbar's network monitoring and the demo scripts in the deck we presented. The demo flow is based on direct login but shows all relevant API calls.

Any questions, I'm happy to help, just let me know...


Cheers,
Frej

On 05/25/2015 02:11 PM, Dave Goodchild wrote:
Hi Frej, what I am aiming to do for a 10duke POC initially is:

1. Store an entitlement record in 10duke.
2. Adapt our access control API to consult 10duke as well as our current access system for an access decision.