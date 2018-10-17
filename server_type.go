package a2s

type ServerType int

const (
	ServerType_Unknown ServerType = iota
	ServerType_Dedicated
	ServerType_NonDedicated
	ServerType_SourceTV
)

func ParseServerType(servertype uint8) ServerType {
	switch servertype {
	case uint8('d'):
		return ServerType_Dedicated
	case uint8('l'):
		return ServerType_NonDedicated
	case uint8('p'):
		return ServerType_SourceTV
	}

	return ServerType_Unknown
}
