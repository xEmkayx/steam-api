package config

type OutputFormat int64

const (
	Json = iota
	Xml
	Vdf
)

func (o OutputFormat) String() string {
	switch o {
	case Json:
		return "json"
	case Xml:
		return "xml"
	case Vdf:
		return "vdf"
	default:
		return "json"
	}
}

// -------------------------------------

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

func (l LogLevel) String() string {
	switch l {
	case Debug:
		return "Debug"
	case Info:
		return "Info"
	case Warning:
		return "Warning"
	case Error:
		return "Error"
	default:
		return "Unknown Level"
	}
}

// -------------------------------------

type FriendsListRelationship string

const (
	All    = "all"
	Friend = "friend"
)

func (f FriendsListRelationship) String() string {
	switch f {
	case All:
		return "all"
	case Friend:
		return "friend"
	default:
		return "all"
	}
}

// -------------------------------------

type Language string

// todo: add languages
const (
	English = "english"
	German  = "german"
)

func (l Language) String() string {
	switch l {
	case English:
		return "english"
	case German:
		return "german"
	default:
		return "english"
	}
}
