package as

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// timeformats contains the supported time formats
// for the convertion to time.Time.
var timeformats = []string{
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	time.Kitchen,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
	"02.01.06",
	"01/02/06",
	"2006-01-02",
	"2006/01/02",
	"01/02/2006",
	"02.01.2006",
	"01/02/06 15:04",
	"2006-01-02 15:04",
	"2006-01-02T15:04",
	"01/02/2006 15:04",
	"02.01.06 15:04:05",
	"01/02/06 15:04:05",
	"01/02/2006 15:04:05",
	"2006-01-02 15:04:05",
	"2006-01-02T15:04:05",
	"02.01.2006 15:04:05",
}

// Bool returns a boolean value. 
// It mainly depends on the output of strconv.ParseBool, 
// but also checks for integer values.
func Bool(value interface{}) bool {
	if Int(value) > 0 {
		return true
	}
	b, _ := strconv.ParseBool(String(value))
	return b
}

// Bytes returns a slice of bytes.
func Bytes(value interface{}) []byte {
	if value == nil {
		return []byte{}
	}

	switch val := value.(type) {
	case bool:
		if val == true {
			return []byte("true")
		}
		return []byte("false")
	case string:
		return []byte(val)
	case []byte:
		return val
	default:
		return []byte(fmt.Sprintf("%v", value))
	}
}

// Duration converts input values to time.Duration.
// It mainly depends on time.ParseDuration.
func Duration(value interface{}) time.Duration {
	switch value.(type) {
	case int, int8, int16, int32, int64:
		return time.Duration(Int(value))
	case uint, uint8, uint16, uint32, uint64:
		return time.Duration(Int(value))
	case float32, float64:
		return time.Duration(Int(value))
	default:
		dur, _ := time.ParseDuration(String(value))
		return dur
	}
}

// Float converts it's input to type float64.
// int, uint and float gets converted as expected,
// time is transformed to a float of the corresponding timestamp.
// strings and byte slices gets converted via strconv.ParseFloat.
func Float(value interface{}) float64 {
	switch val := value.(type) {
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case uint:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)
	case float32:
		return float64(val)
	case float64:
		return float64(val)
	case time.Time:
		return float64(val.Unix())
	case bool:
		if val == true {
			return float64(1)
		}
		return float64(0)
	default:
		f, _ := strconv.ParseFloat(String(value), 64)
		return float64(f)
	}
}

// FloatFromXString converts strings to float64.
// Most values can be converted to float via Float(),
// but floats as strings in e.g. german spelling 
// should be converted with this function.
func FloatFromXString(value string) float64 {
	value = strings.Trim(value, "\t\n\r¢§$€ ")
	var float float64
	c := strings.Count(value, ",")
	p := strings.Count(value, ".")
	fc := strings.Index(value, ",")
	fp := strings.Index(value, ".")
	if c == 0 && p == 1 {
		float, _ = strconv.ParseFloat(value, 64)
	} else if c == 1 && p == 0 {
		value = strings.Replace(value, ",", ".", 1)
		float, _ = strconv.ParseFloat(value, 64)
	} else if c == 0 && p == 0 {
		intx, _ := strconv.ParseInt(value, 0, 64)
		float = float64(intx)
	} else if c > 1 && p < 2 {
		value = strings.Replace(value, ",", "", -1)
		float, _ = strconv.ParseFloat(value, 64)
	} else if c < 2 && p > 1 {
		value = strings.Replace(value, ".", "", -1)
		value = strings.Replace(value, ",", ".", 1)
		float, _ = strconv.ParseFloat(value, 64)
	} else if c == 1 && p == 1 {
		if fp < fc {
			value = strings.Replace(value, ".", "", -1)
			value = strings.Replace(value, ",", ".", 1)
		} else {
			value = strings.Replace(value, ",", "", -1)
		}
		float, _ = strconv.ParseFloat(value, 64)
	} else {
		value = "0"
		float, _ = strconv.ParseFloat(value, 64)
	}
	return float64(float)
}

// Int returns an int64 of the input value.
// Float values and float values in strings will be rounded via 
// "round half towards positive infinity".
// strings get converted via strconv.ParseFloat.
func Int(value interface{}) int64 {
	switch val := value.(type) {
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return int64(val)
	case uint:
		return int64(val)
	case uint8:
		return int64(val)
	case uint16:
		return int64(val)
	case uint32:
		return int64(val)
	case uint64:
		return int64(val)
	case float32:
		return int64(val + 0.5)
	case float64:
		return int64(val + 0.5)
	case time.Time:
		return int64(val.Unix())
	case bool:
		if val == true {
			return int64(1)
		}
		return int64(0)
	default:
		i, _ := strconv.ParseFloat(String(value), 64)
		return int64(i + 0.5)
	}
}

// String converts input values to string.
// Time and Duration gets converted via standard functions.
// Most types gets "converted" via fmt.Sprintf.
func String(value interface{}) string {
	if value == nil {
		return ""
	}

	switch val := value.(type) {
	case bool:
		if value.(bool) == true {
			return "true"
		}
		return "false"
	case time.Duration:
		return string(val.String())
	case time.Time:
		return string(val.Format(time.RFC3339))
	case string:
		return string(val)
	case []byte:
		return string(val)
	default:
		return fmt.Sprintf("%v", val)
	}
}

// Time converts inputs values to time.Time.
// Time formats in the variable timeformats can be used.
func Time(value interface{}) time.Time {
	s := String(value)
	for _, format := range timeformats {
		r, err := time.Parse(format, s)
		if err == nil {
			return r
		}
	}
	return time.Time{}
}

// Uint returns an uint64 of the input value.
// Float values and float values in strings will be rounded via 
// "round half towards positive infinity".
// strings get converted via strconv.ParseFloat.
func Uint(value interface{}) uint64 {

	switch val := value.(type) {
	case int:
		return uint64(val)
	case int8:
		return uint64(val)
	case int16:
		return uint64(val)
	case int32:
		return uint64(val)
	case int64:
		return uint64(val)
	case uint:
		return uint64(val)
	case uint8:
		return uint64(val)
	case uint16:
		return uint64(val)
	case uint32:
		return uint64(val)
	case uint64:
		return uint64(val)
	case float32:
		return uint64(val + 0.5)
	case float64:
		return uint64(val + 0.5)
	case time.Time:
		return uint64(val.Unix())
	case bool:
		if val == true {
			return uint64(1)
		}
		return uint64(0)
	default:
		i, _ := strconv.ParseFloat(String(value), 64)
		return uint64(i + 0.5)
	}
}
