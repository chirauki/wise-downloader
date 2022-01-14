package cmd

const (
	dateLayout = "2006-01-02"

	flagToken          = "api-token"
	flagProfile        = "profile"
	flagStartDate      = "start-date"
	flagEndDate        = "end-date"
	flagCurrency       = "currency"
	flagLimit          = "limit"
	flagOffset         = "offset"
	flagStatus         = "status"
	flagSourceCurrency = "source-currency"
	flagTargetCurrency = "target-currency"
	flagTransfer       = "transfer"
	flagOutput         = "output"
)

var (
	token          string
	profile        int
	startDate      string
	endDate        string
	currency       string
	limit          int
	offset         int
	status         string
	sourceCurrency string
	targetCurrency string
	transfer       int
	output         string
)
