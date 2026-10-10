package htnstratum

import "strings"

// plainNodeError turns a node RPC failure into a log line a miner can read.
// The raw error stays available to the caller. This only names the case.
func plainNodeError(err error) string {
	if err == nil {
		return "node returned an empty error"
	}
	msg := err.Error()
	switch {
	case strings.Contains(msg, "not synced") || strings.Contains(msg, "IBD running"):
		return "node still syncing, block not submitted"
	case strings.Contains(msg, "no peers") || strings.Contains(msg, "nearly synced"):
		return "node has no peers yet, block not submitted"
	case strings.Contains(msg, "ErrBadPruningPointUTXOSet"):
		return "pruning-point set does not match the header, block not submitted"
	case strings.Contains(msg, "ErrDuplicateBlock"):
		return "block already known"
	case strings.Contains(msg, "ErrInvalidPoW"):
		return "proof of work does not match"
	case strings.Contains(msg, "Could not decode address"):
		return "miner address is not a valid hoosat address"
	case strings.Contains(msg, "failed fetching new block template"):
		return "node did not return a block template"
	default:
		return "block rejected by node"
	}
}
