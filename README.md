openfaas-functions
=====

## filter-tweet

Filters out retweets.

Seal the secrets:

```sh
faas-cli cloud seal --name civo-slack-incoming-webhook-url \
    --literal incoming-webhook-url=https://hooks.slack.com/services/value-here
```

## Additional info

These functions are deployed to the [OpenFaaS Community Cluster](https://github.com/openfaas/community-cluster/) operated by OpenFaaS Ltd.

