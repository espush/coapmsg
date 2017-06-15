package coapmsg

import (
	"errors"
	"math/rand"
	"net"
	"sync"
	"time"
)

// CurrentMessageID stores the current message id used/generated for messages
var CurrentMessageID = 0
var MESSAGEID_MUTEX *sync.Mutex

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	CurrentMessageID = rand.Intn(65535)

	MESSAGEID_MUTEX = &sync.Mutex{}
}

var GENERATE_ID uint16 = 0

const UDP = "udp"

// Types of Messages
const (
	MessageConfirmable    = 0
	MessageNonConfirmable = 1
	MessageAcknowledgment = 2
	MessageReset          = 3
)

// Fragments/parts of a CoAP Message packet
const (
	DataHeader     = 0
	DataCode       = 1
	DataMsgIDStart = 2
	DataMsgIDEnd   = 4
	DataTokenStart = 4
)

// OptionCode type represents a valid CoAP Option Code
type OptionCode int

const (
	// OptionIfMatch request-header field is used with a method to make it conditional.
	// A client that has one or more entities previously obtained from the resource can verify
	// that one of those entities is current by including a list of their associated entity tags
	// in the If-Match header field.
	OptionIfMatch OptionCode = 1

	OptionURIHost       OptionCode = 3
	OptionEtag          OptionCode = 4
	OptionIfNoneMatch   OptionCode = 5
	OptionObserve       OptionCode = 6
	OptionURIPort       OptionCode = 7
	OptionLocationPath  OptionCode = 8
	OptionURIPath       OptionCode = 11
	OptionContentFormat OptionCode = 12
	OptionMaxAge        OptionCode = 14
	OptionURIQuery      OptionCode = 15
	OptionAccept        OptionCode = 17
	OptionLocationQuery OptionCode = 20
	OptionBlock2        OptionCode = 23
	OptionBlock1        OptionCode = 27
	OptionSize2         OptionCode = 28
	OptionProxyURI      OptionCode = 35
	OptionProxyScheme   OptionCode = 39
	OptionSize1         OptionCode = 60
)

// CoapCode defines a valid CoAP Code Type
type CoapCode uint8

const (
	Get    CoapCode = 1
	Post   CoapCode = 2
	Put    CoapCode = 3
	Delete CoapCode = 4

	// 2.x
	CoapCodeEmpty    CoapCode = 0
	CoapCodeCreated  CoapCode = 65 // 2.01
	CoapCodeDeleted  CoapCode = 66 // 2.02
	CoapCodeValid    CoapCode = 67 // 2.03
	CoapCodeChanged  CoapCode = 68 // 2.04
	CoapCodeContent  CoapCode = 69 // 2.05
	CoapCodeContinue CoapCode = 95 // 2.31

	// 4.x
	CoapCodeBadRequest               CoapCode = 128 // 4.00
	CoapCodeUnauthorized             CoapCode = 129 // 4.01
	CoapCodeBadOption                CoapCode = 130 // 4.02
	CoapCodeForbidden                CoapCode = 131 // 4.03
	CoapCodeNotFound                 CoapCode = 132 // 4.04
	CoapCodeMethodNotAllowed         CoapCode = 133 // 4.05
	CoapCodeNotAcceptable            CoapCode = 134 // 4.06
	CoapCodeRequestEntityIncomplete  CoapCode = 136 // 4.08
	CoapCodeConflict                 CoapCode = 137 // 4.09
	CoapCodePreconditionFailed       CoapCode = 140 // 4.12
	CoapCodeRequestEntityTooLarge    CoapCode = 141 // 4.13
	CoapCodeUnsupportedContentFormat CoapCode = 143 // 4.15

	// 5.x
	CoapCodeInternalServerError  CoapCode = 160 // 5.00
	CoapCodeNotImplemented       CoapCode = 161 // 5.01
	CoapCodeBadGateway           CoapCode = 162 // 5.02
	CoapCodeServiceUnavailable   CoapCode = 163 // 5.03
	CoapCodeGatewayTimeout       CoapCode = 164 // 5.04
	CoapCodeProxyingNotSupported CoapCode = 165 // 5.05
)

const DefaultAckTimeout = 2
const DefaultAckRandomFactor = 1.5
const DefaultMaxRetransmit = 4
const DefaultNStart = 1
const DefaultLeisure = 5
const DefaultProbingRate = 1

const CoapDefaultHost = ""
const CoapDefaultPort = 5683
const CoapsDefaultPort = 5684

const PayloadMarker = 0xff
const MaxPacketSize = 1500

// MessageIDPurgeDuration defines the number of seconds before a MessageID Purge is initiated
const MessageIDPurgeDuration = 60

// Proxy Filter
type ProxyFilter func(Message, net.Addr) bool

type MediaType int

const (
	MediaTypeTextPlain                  MediaType = 0
	MediaTypeTextXML                    MediaType = 1
	MediaTypeTextCsv                    MediaType = 2
	MediaTypeTextHTML                   MediaType = 3
	MediaTypeImageGif                   MediaType = 21
	MediaTypeImageJpeg                  MediaType = 22
	MediaTypeImagePng                   MediaType = 23
	MediaTypeImageTiff                  MediaType = 24
	MediaTypeAudioRaw                   MediaType = 25
	MediaTypeVideoRaw                   MediaType = 26
	MediaTypeApplicationLinkFormat      MediaType = 40
	MediaTypeApplicationXML             MediaType = 41
	MediaTypeApplicationOctetStream     MediaType = 42
	MediaTypeApplicationRdfXML          MediaType = 43
	MediaTypeApplicationSoapXML         MediaType = 44
	MediaTypeApplicationAtomXML         MediaType = 45
	MediaTypeApplicationXmppXML         MediaType = 46
	MediaTypeApplicationExi             MediaType = 47
	MediaTypeApplicationFastInfoSet     MediaType = 48
	MediaTypeApplicationSoapFastInfoSet MediaType = 49
	MediaTypeApplicationJSON            MediaType = 50
	MediaTypeApplicationXObitBinary     MediaType = 51
	MediaTypeTextPlainVndOmaLwm2m       MediaType = 1541
	MediaTypeTlvVndOmaLwm2m             MediaType = 1542
	MediaTypeJSONVndOmaLwm2m            MediaType = 1543
	MediaTypeOpaqueVndOmaLwm2m          MediaType = 1544
)

const (
	MethodGet     = "GET"
	MethodPut     = "PUT"
	MethodPost    = "POST"
	MethodDelete  = "DELETE"
	MethodOptions = "OPTIONS"
	MethodPatch   = "PATCH"
)

type BlockSizeType byte

const (
	BlockSize16   BlockSizeType = 0
	BlockSize32   BlockSizeType = 1
	BlockSize64   BlockSizeType = 2
	BlockSize128  BlockSizeType = 3
	BlockSize256  BlockSizeType = 4
	BlockSize512  BlockSizeType = 5
	BlockSize1024 BlockSizeType = 6
)

// Errors
var ErrPacketLengthLessThan4 = errors.New("Packet length less than 4 bytes")
var ErrInvalidCoapVersion = errors.New("Invalid CoAP version. Should be 1.")
var ErrOptionLengthUsesValue15 = errors.New(("Message format error. Option length has reserved value of 15"))
var ErrOptionDeltaUsesValue15 = errors.New(("Message format error. Option delta has reserved value of 15"))
var ErrUnknownMessageType = errors.New("Unknown message type")
var ErrInvalidTokenLength = errors.New("Invalid Token Length ( > 8)")
var ErrUnknownCriticalOption = errors.New("Unknown critical option encountered")
var ErrUnsupportedMethod = errors.New("Unsupported Method")
var ErrNoMatchingRoute = errors.New("No matching route found")
var ErrUnsupportedContentFormat = errors.New("Unsupported Content-Format")
var ErrNoMatchingMethod = errors.New("No matching method")
var ErrNilMessage = errors.New("Message is nil")
var ErrNilConn = errors.New("Connection object is nil")
var ErrNilAddr = errors.New("Address cannot be nil")
var ErrMessageSizeTooLongBlockOptionValNotSet = errors.New("Message is too long, block option or value not set")

type Option interface {
	Name() string
	IsElective() bool
	IsCritical() bool
	StringValue() string
	IntValue() int
	GetCode() OptionCode
	GetValue() interface{}
}

// Represents the payload/content of a CoAP Message
type MessagePayload interface {
	GetBytes() []byte
	Length() int
	String() string
}

type Message interface {
	GetToken() []byte
	GetMessageId() uint16
	GetMessageType() uint8
	GetAcceptedContent() MediaType
	GetCodeString() string
	GetCode() CoapCode
	GetMethod() uint8
	GetTokenLength() uint8
	GetTokenString() string
	GetOptions(id OptionCode) []Option
	GetOption(id OptionCode) Option
	GetAllOptions() []Option
	GetOptionsAsString(id OptionCode) []string
	GetLocationPath() string
	GetURIPath() string
	GetPayload() MessagePayload

	SetToken([]byte)
	SetMessageId(uint16)
	SetMessageType(uint8)
	SetBlock1Option(opt Option)
	SetStringPayload(s string)
	SetPayload(MessagePayload)

	AddOption(code OptionCode, value interface{})
	AddOptions(opts []Option)
	CloneOptions(cm Message, opts ...OptionCode)
	ReplaceOptions(code OptionCode, opts []Option)
	RemoveOptions(id OptionCode)
}

type BlockMessage interface {
}
