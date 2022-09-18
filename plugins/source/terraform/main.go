package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/terraform/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://8a66aa7c550b46258f2391cbae261fe2@o1396617.ingest.sentry.io/6747630"

func main() {
	serve.Serve(serve.Options{
		SourcePlugin: plugin.Plugin(),
		SentryDsn:    sentryDSN,
	})
}
