```
   ________ _____ ___
  / ____|_   _|  ____|
 | |      | | | |__   
 | |      | | |  __|  
 | |____ _| |_| |____ 
  \_____|_____|______|
```


<div>
  <img alt="Repo Size" src="https://img.shields.io/github/repo-size/michalszmidt/cie" />
  <img alt="Lines of code" src="https://sloc.xyz/github/michalszmidt/cie?category=code" />
  <img alt="Last Commit" src="https://img.shields.io/github/last-commit/michalszmidt/cie" />
  <img alt="Assets Downloads" src="https://img.shields.io/github/downloads/michalszmidt/cie/total" />
</div>

# About
cie is a simple CLI ICal editor managed by yaml


# Installation
Download from release page

# Usage scenario
You have fixed schedule exported to ics but you made some unofficial modifications to it that are not covered by the issuer.
So you set up cron job to:
1. fetch ics
2. `cie` thath makes modifications 

it produces another ics that you expose a path to on your server so, your clients can fetch modified one.
