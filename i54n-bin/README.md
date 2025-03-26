# i54n-bin

Welcome to the i54n-bin repository, which contains the bior binary for the
i54n project. The bior binary is a bioreactor management tool built with Go,
Cobra, and Mage to provide a clear and structured command-line interface.

Use `mage Build` to compile the bior binary and `mage Clean` to remove the
compiled binary. The bior tool includes subcommands such as `about` and
`license` to display project information and licensing details.

This repository supports the management of bioreactor schemas, the validation
of YAML definitions, and the execution of bioreactor control commands.
