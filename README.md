# Grace Hopper OS Example

This repository serves as an example of an Open Source repo. The code is a simple 

Slides: https://docs.google.com/presentation/d/1ZCUBTZErugAegwAl-1VWpSKtqah9sfOjXpZnOvmHkuk/edit?usp=sharing

Link to PR that enabled that feature

## Initial set-up/Bootstrapping:
- [ ] README.md
- [ ] CLA or DCO bot
   - [DCO](https://github.com/organizations/oss-workshop/settings/installations/11858404)
   - https://opensource.com/article/18/3/cla-vs-dco-whats-difference
- [ ] [Docker](https://github.com/oss-workshop/grace-hopper-example/commit/40d1e13390f9dc9c0a8e29a7c207a2af4a19cc99)
    - Github Container Registry
    - Github action to create image for each release
- [ ] Road Map
    - Give users an idea of where project is going
- [ ] [Security](https://github.com/oss-workshop/grace-hopper-example/blob/master/SECURITY.md)
    - Allow users to report security issues without compromising project
- [Stale Bot](https://github.com/organizations/oss-workshop/settings/installations/11858383)
    - Deal with old/abandoned PRs
- Linter *(mentioned in demo)
- [Semantic Pull Requests](https://github.com/organizations/oss-workshop/settings/installations/11858319)
    - Easily filter commits for CHANGELOG.md
- Documentation structure (mkdocs/markdown)
- Code coverage tracker    
 
For an OS repo, you need to set up security, create a roadmap for users to follow, etc.

# Maintenance
### Daily Stand-Up Meeting:
- Necessary to regularly review new issues/PRs
    - Review OSS issues or PRs opened in the last working day, so we do not miss anything.
    - Review owned issues, maybe reassign or disown.
    - Review open PRs, so we agree that it should be done or do not forget about them.

### Testing
- [ ] Create robust unit/e2e tests
    - Will improve quality of contributions
    - Ensure that other users' changes meet necessary standards

### Contribution
- [ ] Require certain % code coverage from contributors
- [ ] Must provide unit/e2e tests
   
