openfaas-functions
=====

## filter-tweet

Filters out retweets.

### Seal the secrets for the incoming webhook URL

Seal the secrets:

```sh
faas-cli cloud seal --name civo-slack-incoming-webhook-url \
    --literal incoming-webhook-url=https://hooks.slack.com/services/value-here
```

### Filter out a term

Edit [here](https://github.com/civo/openfaas-functions/blob/master/filter-tweet/handler.go#L28)

### Endpoint

For use with IFTTT and other integrations:

```
https://civo.o6s.io/filter-tweet
```

## Additional info

These functions are deployed to the [OpenFaaS Community Cluster](https://github.com/openfaas/community-cluster/) operated by OpenFaaS Ltd.

