package duke_nukem

/* LATEST VALUES AS CONFIRMED BY 10DUKE

operation=RegisterUser
&/Profile[@id='{randomUuid,profile}']/@id={randomUuid,profile}&/Profile[@id='{randomUuid,profile}']/~ManyToOne/Person[@id='{randomUuid,person}']/@firstName={{.FirstName}}
&/Profile[@id='{randomUuid,profile}']/~ManyToOne/Person[@id='{randomUuid,person}']/@lastName={{.LastName}}
&/Profile[@id='{randomUuid,profile}']/~OneToMany/ContactInformation[@id='{randomUuid,contactInformation}']/~OneToMany/EmailAddress[@id='{randomUuid,emailAddress}']/@value={{.Email}}
&/Profile[@id='{randomUuid,profile}']/~ManyToOne/Person[@id='{randomUuid,person}']/~OneToMany/EmailAndPassword[@id='{randomUuid,emailAndPassword}']/@userName={{.Email}}
&/Profile[@id='{randomUuid,profile}']/~ManyToOne/Person[@id='{randomUuid,person}']/~OneToMany/EmailAndPassword[@id='{randomUuid,emailAndPassword}']/@passwordPlain={{.Password}}
&confirmedPassword={{.Password}}
&/Profile[@id='{randomUuid,profile}']/~OneToMany/Account[@id='{randomUuid,account}']/@id={randomUuid,account}&/Profile[@id='{randomUuid,profile}']/~OneToMany/Account[@id='{randomUuid,account}']/@accountType=personal
&/Profile[@id='{randomUuid,profile}']/~OneToMany/ContactInformation[@id='{randomUuid,contactInformation}']/~OneToMany/PostalAddress[@id='{randomUuid,postalAddress}']/@countryCode={{.Country}}
&acceptsTsAndCs=true
&accountType=personal

*/

var (
	userProfile     = "/Profile[@id='{randomUuid,profile}']/@id"
	userFirstName   = "/Profile[@id='{randomUuid,profile}']/~ManyToOne/Person[@id='{randomUuid,person}']/@firstName"
	userLastName    = "/Profile[@id='{randomUuid,profile}']/~ManyToOne/Person[@id='{randomUuid,person}']/@lastName"
	userEmail       = "/Profile[@id='{randomUuid,profile}']/~OneToMany/ContactInformation[@id='{randomUuid,contactInformation}']/~OneToMany/EmailAddress[@id='{randomUuid,emailAddress}']/@value"
	userName        = "/Profile[@id='{randomUuid,profile}']/~ManyToOne/Person[@id='{randomUuid,person}']/~OneToMany/EmailAndPassword[@id='{randomUuid,emailAndPassword}']/@userName"
	userPassword    = "/Profile[@id='{randomUuid,profile}']/~ManyToOne/Person[@id='{randomUuid,person}']/~OneToMany/EmailAndPassword[@id='{randomUuid,emailAndPassword}']/@passwordPlain"
	userConfirm     = "confirmedPassword"
	userAccountType = "/Profile[@id='{randomUuid,profile}']/~OneToMany/Account[@id='{randomUuid,account}']/@id"
	userAccount     = "/Profile[@id='{randomUuid,profile}']/~OneToMany/Account[@id='{randomUuid,account}']/@accountType"
	userTsAndCs     = "acceptsTsAndCs"
)
