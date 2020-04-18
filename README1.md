# ELASTIC SEARCH NOTES

The basic as we already know elastic search is a powerful database which is most famous for it's search indexing capabilities. In this notes i will be referencing The complete guide to elastic search.

## Shards:

 Pieces of an index(a table) are called Shards.

## Why is it necessary?

Sharding is necessary when you  you have an index that is too big to be stored on a single cluster or nodes(a group of nodes where your elastic search is hosted). You simply split the node into shards and store it into multiple clusters or nodes.

You might do a thing called over-sharding that is creating too many unnecessary shards.

## Replication?

To handle  and manage a system it is  important that we make it fault tolerant. There might be a chance that our index crashes due to some reason. To safeguard us from this, we create replicas of indices(more than one index). Event shards are duplicated they are called *replica shards*  thus making the shards which has been replicated the ***Primary Shards***

![ELASTIC%20SEARCH%20NOTES/Screenshot_(172).png](ELASTIC%20SEARCH%20NOTES/Screenshot_(172).png)

but again the question arises what if the whole node fails ? we solve this by storing shards and replica shards on different nodes

![ELASTIC%20SEARCH%20NOTES/Screenshot_(173).png](ELASTIC%20SEARCH%20NOTES/Screenshot_(173).png)

- if for some reason you are facing  bottleneck in the number of queries in the index you can always use the shard as it is a fully functioning index on its own.

# CHEAT SHEET:

[Elasticsearch Cheatsheet : Example API usage of using Elasticsearch with curl](https://gist.github.com/ruanbekker/e8a09604b14f37e8d2f743a87b930f93)

## Scripted Updates:

    POST /products/_update/100
    {
       "script": {
        "source": "ctx._source.in_stock-=params.quantity",
        "params": {
          "quantity":4
        }
      }
      
    }
    

## UPSERT:

only runs if the doc exists else a new upsert doc is created:

    POST /products/_update/101
    {
       "script": {
        "source": "ctx._source.quantity--"
      }
      , "upsert": {
        "name": "Sandwich maker",
        "Price": 122,
        "quantity":100
      }
      
    }
