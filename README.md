# torch
Torch brings the light to the world.

## Design

Torch is a simple gate to solve the prometheus single node problem. 
It can be integrated with prometheus well, and makes prometheus horizontal scalable.

torch's whole architecture is shown below:

![](https://github.com/Colstuwjx/torch/blob/master/arch/torch.png)

### Torch SD

Each torch node can act as a SD for the backend prometheus instance, it maintains the
hashing strategy with upstreams, and tells the prometheus instance which targets it need to scrape.

### Maintain the instance status

BTW, torch also do health check with the upstreams, and maintains their health state, render
the target groups and route the query service requests unless the prometheus instance is healthy.

### Gate for query service

As we mentioned in the previous section, torch nodes also act as the gate for query services,
and route the queries to a healthy and matched shard(the backend prometheus instance, here we called shard).

### Works well with prometheus

Torch didn't and won't break any prometheus features, as the sd is compat with the official discovery solution and each backend prometheus instance still could eat the community benifits.

### Drawback

Torch can NOT solve the prometheus storage engine side problems, such as crash recovery, chunk ops, cold store etc.

it also can NOT map the rule evalution since each rule is still evaluated on the each prometheus instance side, what's more, each shard in the same replica maybe has inconsistent data in risk, as each shard do scrape by it's own target scraper.
