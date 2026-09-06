# ETE CICD pipeline

So, here is the deal, idea is to setup things in such a way that, entire ci and cd is automated completely.

TODO

- Improve Poller loop to check harbour and then build for next version, covers the cases where it checked a version before but the build didnt go through completly. currently there is mismatch between harbour and what poller thinks in existence of image versions

- dev/staging/prod env with promotation flow




Features worth seeing

- ETE project build and artifact storage
- CVE scanning in harbor