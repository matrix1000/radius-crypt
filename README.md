# radius-crypt

Small script in Go for creating the following 
password hashes for freeradius.

SMD5-Password
SSHA2-256-Password
SSHA2-512-Password

More information: https://freeradius.org/radiusd/man/rlm_pap.html

Usage: go run radius-crypt.go PASSWORD
Salt: Ob05RHt9JezkCQIh_fyG
{SMD5-Password}: xpZSntJlkZO70ZeVufm8o09iMDVSSHQ5SmV6a0NRSWhfZnlH
{SSHA2-256-Password}: Lg5HjeGH5+jXsajN6pNt6FJbwSKK6NrT+i66nOwQQL1PYjA1Ukh0OUplemtDUUloX2Z5Rw==
{SSHA2-512-Password}: 9cc0QNc/MgVnUAr+RHCNrPLME6X6b8shrkmNkrW/NIhPilyLE78DFXNVh78ivjuTvn7KOz9MuSStzBuwnpxZ1E9iMDVSSHQ5SmV6a0NRSWhfZnlH
