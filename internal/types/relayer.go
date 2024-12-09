package types

import "context"

type Relayer interface {
	Start(ctx context.Context)
	Shutdown()
	SetReadyToShutdown()
}
