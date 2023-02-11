# Chinde Mail Indexer

This project indexes the content of Enron emails in ZincSearch. It is built using the go code.

## Indexation

### Download the Enron Corp Mails

The mails can be downloaded from this [link](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz).

Download the tgz file and unzip it to your preferred local path. The path is necessary for the future step.

### Set the following OS variables

The path to the root folder is from the previous step, the public DNS corresponds to the public IP obtained in the terraform deployment, [see this](../README.md).

```sh
export ENRON_ROOT_FOLDER="/path/to/enronmails/root/folder/"
export AWS_PUBLIC_DNS="ec2-aws-public-dns"
```

Set the ZincSearch credentials, as well:

```sh
export ZS_USER="zincsearch user"
export ZS_PASSWORD="zincsearch password"
```

### Create firts index

For better performance, create index firts with mapping configuration.

```sh
curl -X POST "http://${aws_public_dns}:4080/api/index" \
-H "Content-Type: application/json" \
-d '{
  "name": "enronmails",
  "shard_num": 1,
  "storage_type": "disk",
  "mappings": {
    "properties": {
      "@timestamp": {
        "type": "date",
        "index": true,
        "store": false,
        "sortable": true,
        "aggregatable": true,
        "highlightable": false
      },
      "Bcc": {
        "type": "text",
        "index": true,
        "store": false,
        "sortable": false,
        "aggregatable": false,
        "highlightable": false
      },
      "Body": {
        "type": "text",
        "index": true,
        "store": false,
        "sortable": false,
        "aggregatable": false,
        "highlightable": true
      },
      "Cc": {
        "type": "text",
        "index": true,
        "store": false,
        "sortable": false,
        "aggregatable": false,
        "highlightable": false
      },
      "Date": {
        "type": "text",
        "index": true,
        "store": false,
        "sortable": false,
        "aggregatable": false,
        "highlightable": false
      },
      "From": {
        "type": "text",
        "index": true,
        "store": false,
        "sortable": false,
        "aggregatable": false,
        "highlightable": false
      },
      "MessageID": {
        "type": "text",
        "index": true,
        "store": false,
        "sortable": false,
        "aggregatable": false,
        "highlightable": false
      },
      "Subject": {
        "type": "text",
        "index": true,
        "store": false,
        "sortable": false,
        "aggregatable": false,
        "highlightable": false
      },
      "To": {
        "type": "text",
        "index": true,
        "store": false,
        "sortable": false,
        "aggregatable": false,
        "highlightable": false
      },
      "_id": {
        "type": "keyword",
        "index": true,
        "store": false,
        "sortable": true,
        "aggregatable": true,
        "highlightable": false
      }
    }
  }
}' \
-u "zincsearch-user:zinc-search-password"
```

The successful response is something like this:

```json
{ "message": "ok", "index": "enronmails", "storage_type": "disk" }
```

### Start the indexation

Exec the following command in the root folder of mail-indexer project:

```sh
go run .
```
