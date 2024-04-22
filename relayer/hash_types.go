package relayer

type ExtraData struct {
	ExtraData []byte `ssz-max:"32"`
}

type LogsBloom struct {
	LogsBloom [256]byte `ssz-size:"256"`
}
