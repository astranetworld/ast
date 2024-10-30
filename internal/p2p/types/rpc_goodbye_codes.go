package p2ptypes

import ssztype "github.com/N42world/ast/common/types/ssz"

// RPCGoodbyeCode represents goodbye code, used in sync package.
type RPCGoodbyeCode = ssztype.SSZUint64

const (
	// Spec defined codes.
	GoodbyeCodeClientShutdown RPCGoodbyeCode = iota + 1
	GoodbyeCodeWrongNetwork
	GoodbyeCodeGenericError

	// Teku specific codes
	GoodbyeCodeUnableToVerifyNetwork = RPCGoodbyeCode(128)

	// Lighthouse specific codes
	GoodbyeCodeTooManyPeers = RPCGoodbyeCode(129)
	GoodbyeCodeBadScore     = RPCGoodbyeCode(250)
	GoodbyeCodeBanned       = RPCGoodbyeCode(251)
)

// GoodbyeCodeMessages defines a mapping between goodbye codes and string messages.
var GoodbyeCodeMessages = map[RPCGoodbyeCode]string{
	GoodbyeCodeClientShutdown:        "client shutdown",
	GoodbyeCodeWrongNetwork:          "irrelevant network",
	GoodbyeCodeGenericError:          "fault/error",
	GoodbyeCodeUnableToVerifyNetwork: "unable to verify network",
	GoodbyeCodeTooManyPeers:          "client has too many peers",
	GoodbyeCodeBadScore:              "peer score too low",
	GoodbyeCodeBanned:                "client banned this node",
}

// ErrToGoodbyeCode converts given error to RPC goodbye code.
func ErrToGoodbyeCode(err error) RPCGoodbyeCode {
	switch err {
	case ErrWrongForkDigestVersion:
		return GoodbyeCodeWrongNetwork
	default:
		return GoodbyeCodeGenericError
	}
}
