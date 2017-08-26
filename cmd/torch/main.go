package main

// Torch is a gate for prometheus

// It communicates with the upstream agents(known as `ambassador`),
// handles the `cluster_name->replica->shard`, `persistent hashing`, `rebalance` etc.
// Thus, ambassadors will communicate with torch nodes,
// and fetch its own target groups.

// What's more, torch will also act as a query service gate on the front of prometheus instances.
// It chooses the target prometheus instance via the request metric name,
// and compat with the prometheus official API.

// TODO: works with the consul servers, maintain the service discovery data and shard state via ambassadors.

// TODO: implement a way to handle `the persistent hashing strategy`, and maintain the sharding information.
// Rebalance the sd while there are any shard exited or added.

// TODO: add an API to return the prometheus data, which is compat with the prometheus official API.
