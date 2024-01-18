## Labels



Below is the list of labels and their descriptions used in Gossamer repository.


- **Epic** - Issue used to track development status of a complex feature, aggregates several issues.
- **Feature-branch** - pull request from feature branch to origin.
- **Release** - pull request with next release changes.
- **`A-`**  Action needed label. Marks that there is a specific action needed for issue/PR
  - **A-tooBig** - issue or PR needs to be broken down to smaller parts.
  - **A-stale** - issue or PR is deprecated and needs to be closed. 
  - **A-blocked** - issue or PR is blocked  until something else changes.
  - **A-triage** - issue description needs refactor and/or labeled.
  - **A-debug** - issue requires detective debug work to figure out what's going wrong.
  - **A-design** - issue requires design work to think about how it would best be accomplished.
- **`T-`** Describes the type of issue or pull request.
  - **T-bug** - this issue covers unexpected and/or wrong behaviour. 
  - **T-feat** - this issue/pr is a new feature or functionality.
  - **T-enhancement** - this issue/pr covers improvement of existing functionality. 
  - **T-refactor** - this issue/pr covers refactoring of existing code.  
  - **T-security** - this issue/pr covers security sensitive problem. 
- **`C-`** Complexity label. We operate only 3 complexity levels.
  - **C-simple** -  Minor changes changes, no additional research needed. Good first issue/review.
  - **C-complex** - Complex changes across multiple modules. Possibly will require additional research.
  - **C-chaotic** - Unpredictable nature of this task/changes makes its chaotic.
- **`P-`** Priority level. We only have 3 priority levels, everything else is average by default. 
  - **P-critical** - this must be fixed immediately or contributors or users will be severely impacted.
  - **P-high** - this should be addressed ASAP. **Colour #FBCA04**
  - **P-low** - this is mostly nice to have. **Colour #0E8A16**
- **`S-`** Scope this work related to, could be multiple, but usually this means that task needs to be break down.
  - **S-sync-[westend | kusama | polkadot | paseo]** -  related to particular network syncing.
  - **S-tests** - issue related to adding new tests.  
  - **S-doc** - documentation related.
  - **S-cli** - issue related to Gossamer CLI.
  - **S-ci** - issue related to continuous integration tasks or piplelines.
  - **S-crypto** - issues related to the lib/crypto package.
  - **S-grandpa** - issues related to block finality.
  - **S-babe** - issues related to block production functionality. 
  - **S-runtime**- issues related to the lib/runtime package
  - **S-telemetry** - issue related to node telemetry and metrics reports. 
  - **S-rpc** - issues related to the dot/rpc package.
  - **S-scale** - issues related to the pkg/scale package.
  - **S-utils** - issues related to all other lib packages.
  - **S-network** - issues related to the dot/network package.
  - **S-state** - issues related to dot/state package.
  - **S-subsystems-overseer** - issues related to polkadot host overseer functionality.
  - **S-subsystems-collator** - issues related to polkadot host collator subsystem functionality.
  - **S-subsystems-backing** -  issues related to polkadot host backing subsystem functionality.
  - **S-subsystems-availability** - issues related to polkadot host availability subsystem functionality.
  - **S-subsystems-disputes** - issues related to polkadot host disputes subsystem functionality.