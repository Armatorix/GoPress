<div align="center">
    <img src="icon.png" alt="GoPress Icon" width="250px" height="250px" />
</div>

GoPress is a simple CSM (Content Management System) written in Go. It provides basic functionalities to manage content and users efficiently.

## Features

<!-- - Admin panel for content and user management
- Public API for accessing content
- In-memory content caching for improved performance -->

TODO:
admin pandel:
- [x] adding new content
- [x] changing self password

additional:
- [x] preview mode in editor
- [ ] add handling for tags
- [ ] page view statistics
- [ ] files upload handling to pg + hosting options + compress
- [ ] members handling - adding new members with password
- [ ] in memory content caching
- [ ] first publish date
- [ ] update readme with configs and how to use
- [ ] permalinks - named links in articles except ids
- [ ] organization page with simple user management 

## Options for docker image:
- [x] usage domain (for inner-links and CORS)
- [ ] GoPress domain (for analytic links generation) [or drop it and leave this job for GA?]
- [x] postgres connection uri
