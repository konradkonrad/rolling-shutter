// evtype declares the the different event types sent by shuttermint
package evtype

var (
	Accusation     = "shutter.accusation-registered"
	Apology        = "shutter.apology-registered"
	BatchConfig    = "shutter.batch-config"
	CheckIn        = "shutter.check-in"
	EonStarted     = "shutter.eon-started"
	PolyCommitment = "shutter.poly-commitment-registered"
	PolyEval       = "shutter.poly-eval-registered"
)
