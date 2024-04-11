package Konata_client

type DefaultClient struct {
	opt *Options
	process
	onClose func() error // hook called when client is closed
}

// NewClient returns a client to the Redis Server specified by Options.
func NewClient(opt *Options) *DefaultClient {
	opt.init()

	c := DefaultClient{
		opt: opt,
	}
	c.init()

	return &c
}

func (c *DefaultClient) init() {
	c.process = c.
		c.initHooks(hooks{
		dial:       c.baseClient.dial,
		process:    c.baseClient.process,
		pipeline:   c.baseClient.processPipeline,
		txPipeline: c.baseClient.processTxPipeline,
	})
}
