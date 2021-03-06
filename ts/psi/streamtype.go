package psi

type StreamType byte

const (
	ZeroStreamType StreamType = iota
	MPEG1Video
	MPEG2Video
	MPEG1Audio
	MPEG2Audio
	PrivSect
	PrivPES
	MHEG
	DSMCC
	H222_1
	DSMCC_A
	DSMCC_B
	DSMCC_C
	DSMCC_D
	MPEG2Aux
	AAC
	MPEG4Video
	MPEG4Audio
	DSMCC_SDP
	SPSPES
	SPSSect
	MetaDataPES
	MetaDataSect
	MetaDataDC
	MetaDataOC
	MetaDataDL
	MPEG2IPMP
	H264Video
)

var streamTypes = []string{
	"ZeroStreamType",
	"MPEG1Video",
	"MPEG2Video",
	"MPEG1Audio",
	"MPEG2Audio",
	"PrivSect",
	"PrivPES",
	"MHEG",
	"DSMCC",
	"H222_1",
	"DSMCCA",
	"DSMCCB",
	"DSMCCC",
	"DSMCCD",
	"MPEG2Aux",
	"AAC",
	"MPEG4Video",
	"MPEG4Audio",
	"DSMCC_SDP",
	"SPSPES",
	"SPSSect",
	"MetaDataPES",
	"MetaDataSect",
	"MetaDataDC",
	"MetaDataOC",
	"MetaDataDL",
	"MPEG2IPMP",
	"H264Video",
}

func (t StreamType) String() string {
	if int(t) >= len(streamTypes) {
		return "unknown"
	}
	return streamTypes[t]
}

func ParseStreamType(s string) StreamType {
	for i, t := range streamTypes {
		if t == s {
			return StreamType(i)
		}
	}
	return ZeroStreamType
}
