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
    

UPSERT:

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

### HOW ELASTIC SEARCH ROUTES AND READS DATA?

![ELASTIC%20SEARCH%20NOTES/Screenshot_(178).png](ELASTIC%20SEARCH%20NOTES/Screenshot_(178).png)

### HOW IT WRITES DATA?

![ELASTIC%20SEARCH%20NOTES/Screenshot_(182).png](ELASTIC%20SEARCH%20NOTES/Screenshot_(182).png)

A very common issue can come up when while copying data between a replica group. Say the primary shard just shared the data to the first shard and failed . Then between the remaining two shards one of them is promoted to be the primary shard. 

### Primary terms:

**How many times shards have changed(default 1)**

### Sequence Number:

Primary shards increase this when they process a write request.

**They help shards recover from failure.**

![ELASTIC%20SEARCH%20NOTES/Screenshot_(184).png](ELASTIC%20SEARCH%20NOTES/Screenshot_(184).png)

### Versioning:

every time a modification is made the ***_version*** tag value is increased by one. this is also really helpful when it comes to optimized concurrency control in elastic search queries.

### UPDATE BY QUERY:

This API is used when we need to update say N documents inside an index. It does so by creating a snapshot of the index(why? it is explained later). after that it fetches documents from all the shards.

    POST /products/_update_by_query
    {
    	"script":{
    				"source": "ctx._source.in_stock--"
    	},
    	"query":{
    		"match_all":{}
    	}
    		
    
    }

### DELETE BY QUERY:

    POST /products/_delete_by_query
    {
    	"query":{
    		"match_all":{}
    	}
    }

### DOING BULK TASKS:

CREATE UPDATE DELETE TO MASSES ALL AT ONCE:

    POST /_bulk
    {"index":{"_name": "name of index","_id":200} }
    {"name": "Espresso Machine","price":199,"inc_stock":5 }
    {"create":{"_index":"products","_id":201}} 
    {"name":"Milk Powder", "price": 149,"in_stock":12 }
    
    POST /_bulk
    {"update":{"_index":"products","_id":201} }
    {"doc":{"price":129} }
    {"delete":{"_index":"products","_id":200}}
    ***Only the delete function doesn't expects a second line***

Now if you want to update bulk actions to a single index :

    POST /products/_bulk
    {"update":{"_index":"products","_id":201} }
    {"doc":{"price":129} }
    {"delete":{"_index":"products","_id":200}}

### THINGS TO KEEP IN MIND

![ELASTIC%20SEARCH%20NOTES/Screenshot_(185).png](ELASTIC%20SEARCH%20NOTES/Screenshot_(185).png)

![ELASTIC%20SEARCH%20NOTES/Screenshot_(186).png](ELASTIC%20SEARCH%20NOTES/Screenshot_(186).png)

![ELASTIC%20SEARCH%20NOTES/Screenshot_(187).png](ELASTIC%20SEARCH%20NOTES/Screenshot_(187).png)

img

![ELASTIC%20SEARCH%20NOTES/Screenshot_(188).png](ELASTIC%20SEARCH%20NOTES/Screenshot_(188).png)
