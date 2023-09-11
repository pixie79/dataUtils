// Description: Datadog setup
// Author: Pixie79
// ============================================================================
// package datadog

package datadog

import (
	"strconv"

	"github.com/pixie79/data-utils/utils"
)

var (
	metricsBatchLength    string // metricsBatchLength is the number of metrics to send to Datadog in a single batch
	metricsBatchLengthInt int    // metricsBatchLengthInt is the number of metrics to send to Datadog in a single batch integer
	err                   error
)

// init sets the metricsBatchLengthInt variable
func init() {
	metricsBatchLength = utils.GetEnvDefault("METRICS_BATCH_LENGTH", "800")
	metricsBatchLengthInt, err = strconv.Atoi(metricsBatchLength)
	utils.MaybeDie(err, "cannot convert string metricsBatchLength to int")
}
