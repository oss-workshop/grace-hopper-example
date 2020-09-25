# Grace Hopper OS Example

[Presentation Slides](https://docs.google.com/presentation/d/1ZCUBTZErugAegwAl-1VWpSKtqah9sfOjXpZnOvmHkuk/edit?usp=sharing)

## Initial Set-Up/Bootstrapping:
- [ ] Documentation
    - [ ] ReadMe
        - Provides information/instructions about project
    - [ ] [Owners](https://github.com/oss-workshop/grace-hopper-example/blob/master/OWNERS)
        - Roles that certain users have in repository
    - [ ] [Contributing](https://github.com/oss-workshop/grace-hopper-example/blob/master/docs/CONTRIBUTING.md)
        - Guidelines for contributions
    - [ ] [Changelog](https://github.com/oss-workshop/grace-hopper-example/blob/master/CHANGELOG.md)
        - Record of important changes made to project
    - [ ] [RoadMap](https://github.com/oss-workshop/grace-hopper-example/blob/master/docs/ROADMAP.md)
        - Give users an idea of where project is going
    - [ ] [Security](https://github.com/oss-workshop/grace-hopper-example/blob/master/SECURITY.md)
        - Allow users to report security issues without compromising project
    - [ ] [License](https://choosealicense.com/)
    - [ ] Documentation Structure
- [ ] [CLA or DCO](https://opensource.com/article/18/3/cla-vs-dco-whats-difference)
    - Contributors must sign agreement allowing project to use their contributions
- [ ] [Docker](https://github.com/oss-workshop/grace-hopper-example/commit/40d1e13390f9dc9c0a8e29a7c207a2af4a19cc99)
    - Github Container Registry
    - Github action to create image for each release
- [ ] [Code Coverage](https://github.com/oss-workshop/grace-hopper-example/actions?query=workflow%3A%22Code+Coverage+Workflow%22)
    - Determines number of lines in code validated by existing tests
    - Used to enforce testing protocols
- [ ] [Bots](https://github.com/organizations/oss-workshop/settings/installations)
    - Stale Bot
         - Deal with old/abandoned PRs
    - Linter
        - Enforces code formatting
    - Semantic Pull Requests
        - Easily filter commits for CHANGELOG.md
    - DCO Bot
        - Enforce that DCO is signed
 
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
