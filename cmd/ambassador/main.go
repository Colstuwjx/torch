package main

// Ambassador is an agent lives with the prometheus instance.

// It will rerender the `config.yml` and `target.json`(or maybe the other sd) by torch sd.
// BTW, it also do healthcheck on the prometheus instance side,
// and do service registration/unregister towards the torch nodes.

// TODO: works with the consul agent, and render the `config.yml` and `target.json`
// reload the prometheus instance while it is needed, and do healthchecks to maintain
// the service registration and unregister.
