package utils

import (
	"gopkg.in/urfave/cli.v1"
)

var (
	// ----------------------------
	// ABCI Flags

	TendermintAddrFlag = cli.StringFlag{
		Name:  "tendermint_addr",
		Value: "tcp://localhost:46657",
		Usage: "This is the address that ethermint will use to connect to the tendermint core node. Please provide a port.",
	}

	ABCIAddrFlag = cli.StringFlag{
		Name:  "abci_laddr",
		Value: "tcp://0.0.0.0:46658",
		Usage: "This is the address that the ABCI server will use to listen to incoming connection from tendermint core.",
	}

	ABCIProtocolFlag = cli.StringFlag{
		Name:  "abci_protocol",
		Value: "socket",
		Usage: "socket | grpc",
	}

	VerbosityFlag = cli.IntFlag{
		Name:  "verbosity",
		Value: 3,
		Usage: "Logging verbosity: 0=silent, 1=error, 2=warn, 3=info, 4=core, 5=debug, 6=detail",
	}

	ConfigFileFlag = cli.StringFlag{
		Name:  "config",
		Usage: "TOML configuration file",
	}
)
