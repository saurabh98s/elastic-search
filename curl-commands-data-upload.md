## DATA UPLOAD

use the file products-bulk.json
`curl _H "Content-Type:application/x-ndjson" -XPOST http://localhost:9200/products/_bulk --data-binary "@products-bulk.json"`

check out the shards partition using 
`GET /_cat/shards?v`

`{"took":1992,"errors":false,"items":[{"index":{"_index":"json_products","_type":"_doc","_id":"1","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":0,"_primary_term":1,"status":201}},{"index":{"_ind
ex":"json_products","_type":"_doc","_id":"2","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":1,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"3","_versi
on":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":2,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"4","_version":1,"result":"created","_shards":{"total":2,"succe
ssful":1,"failed":0},"_seq_no":3,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"5","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":4,"_primary_term":1,"
status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"6","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":5,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_t
ype":"_doc","_id":"7","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":6,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"8","_version":1,"result":"created
","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":7,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"9","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_
seq_no":8,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"10","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":9,"_primary_term":1,"status":201}},{"index"
:{"_index":"json_products","_type":"_doc","_id":"11","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":10,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"1
2","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":11,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"13","_version":1,"result":"created","_shards":{"tot
al":2,"successful":1,"failed":0},"_seq_no":12,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"14","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":13,"_pr
imary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"15","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":14,"_primary_term":1,"status":201}},{"index":{"_index":"j
son_products","_type":"_doc","_id":"16","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":15,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"17","_version"
:1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":16,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"18","_version":1,"result":"created","_shards":{"total":2,"succes
sful":1,"failed":0},"_seq_no":17,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"19","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":18,"_primary_term":1
,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"20","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":19,"_primary_term":1,"status":201}},{"index":{"_index":"json_products"
,"_type":"_doc","_id":"21","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":20,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"22","_version":1,"result":"
created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":21,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"23","_version":1,"result":"created","_shards":{"total":2,"successful":1,"fail
ed":0},"_seq_no":22,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"24","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":23,"_primary_term":1,"status":201
}},{"index":{"_index":"json_products","_type":"_doc","_id":"25","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":24,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_do
c","_id":"26","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":25,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"27","_version":1,"result":"created","_sh
ards":{"total":2,"successful":1,"failed":0},"_seq_no":26,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"28","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_
no":27,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"29","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":28,"_primary_term":1,"status":201}},{"index":{
"_index":"json_products","_type":"_doc","_id":"30","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":29,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"31"
,"_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":30,"_primary_term":1,"status":201}},{"index":{"_index":"json_products","_type":"_doc","_id":"32","_version":1,"result":"created","_shards":{"total` <br>

above you can see a sample response