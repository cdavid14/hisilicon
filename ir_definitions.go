package hisilicon

/* Structs Definitions */
type HI_UNF_KEY_STATUS_E int32

const (
	HI_UNF_KEY_STATUS_DOWN HI_UNF_KEY_STATUS_E = iota
	HI_UNF_KEY_STATUS_HOLD HI_UNF_KEY_STATUS_E = iota
	HI_UNF_KEY_STATUS_UP   HI_UNF_KEY_STATUS_E = iota
	HI_UNF_KEY_STATUS_BUTT HI_UNF_KEY_STATUS_E = iota
)

type HI_UNF_IR_CODE_E int32

const (
	HI_UNF_IR_CODE_NEC_SIMPLE HI_UNF_IR_CODE_E = iota
	HI_UNF_IR_CODE_TC9012     HI_UNF_IR_CODE_E = iota
	HI_UNF_IR_CODE_NEC_FULL   HI_UNF_IR_CODE_E = iota
	HI_UNF_IR_CODE_SONY_12BIT HI_UNF_IR_CODE_E = iota
	HI_UNF_IR_CODE_RAW        HI_UNF_IR_CODE_E = iota
	HI_UNF_IR_CODE_BUTT       HI_UNF_IR_CODE_E = iota
)

type HI_UNF_IR_PROTOCOL_E int32

const (
	HI_UNF_IR_NEC  HI_UNF_IR_PROTOCOL_E = 0
	HI_UNF_IR_RC6A HI_UNF_IR_PROTOCOL_E = iota + 10
	HI_UNF_IR_RC5
	HI_UNF_IR_LOW_LATENCY_PROTOCOL
	HI_UNF_IR_RC6_MODE0
	HI_UNF_IR_RCMM
	HI_UNF_IR_RUWIDO
	HI_UNF_IR_RCRF8
	HI_UNF_IR_MULTIPLE
	HI_UNF_IR_RMAP
	HI_UNF_IR_RSTEP
	HI_UNF_IR_RMAP_DOUBLEBIT
	HI_UNF_IR_LOW_LATENCY_PRO_PROTOCOL
	HI_UNF_IR_XMP
	HI_UNF_IR_USER_DEFINED
	HI_UNF_IR_PROTOCOL_BUTT
)

type KeyAttr struct {
	/* upper 16bit data under key mode */
	Upper uint64
	/* lower 32bit data under key mode
	 * or symbol value under symbol mode
	 */
	Lower        uint64
	ProtocolName [32]byte
	/* indentify key status. */
	StatusKey HI_UNF_KEY_STATUS_E
}
