package vibe

import (
	"github.com/spf13/cast"
	"time"
)

//GetString returns the value associated with the key as a string.
func GetString(key string) string { return v.GetString(key) }

func (v *Vibe) GetString(key string) string {
	return cast.ToString(v.Get(key))
}

//GetBool returns the value associated with the key as a boolean.
func GetBool(key string) bool { return v.GetBool(key) }

func (v *Vibe) GetBool(key string) bool {
	return cast.ToBool(v.Get(key))
}

//GetInt returns the value associated with the key as an integer.
func GetInt(key string) int { return v.GetInt(key) }

func (v *Vibe) GetInt(key string) int {
	return cast.ToInt(v.Get(key))
}

//GetInt32 returns the value associated with the key as an integer.
func GetInt32(key string) int32 { return v.GetInt32(key) }

func (v *Vibe) GetInt32(key string) int32 {
	return cast.ToInt32(v.Get(key))
}

//GetInt64 returns the value associated with the key as an integer.
func GetInt64(key string) int64 { return v.GetInt64(key) }

func (v *Vibe) GetInt64(key string) int64 {
	return cast.ToInt64(v.Get(key))
}

//GetUint returns the value associated with the key as an unsigned integer.
func GetUint(key string) uint { return v.GetUint(key) }

func (v *Vibe) GetUint(key string) uint {
	return cast.ToUint(v.Get(key))
}

//GetUint32 returns the value associated with the key as an unsigned integer.
func GetUint32(key string) uint32 {
	return v.GetUint32(key)
}

func (v *Vibe) GetUint32(key string) uint32 {
	return cast.ToUint32(v.Get(key))
}

//GetUint64 returns the value associated with the key as an unsigned integer.
func GetUint64(key string) uint64 {
	return v.GetUint64(key)
}

func (v *Vibe) GetUint64(key string) uint64 {
	return cast.ToUint64(v.Get(key))
}

//GetFloat64 returns the value associated with the key as a float64.
func GetFloat64(key string) float64 {
	return v.GetFloat64(key)
}

func (v *Vibe) GetFloat64(key string) float64 {
	return cast.ToFloat64(v.Get(key))
}

//GetTime returns the value associated with the key as time.
func GetTime(key string) time.Time {
	return v.GetTime(key)
}

func (v *Vibe) GetTime(key string) time.Time {
	return cast.ToTime(v.Get(key))
}
