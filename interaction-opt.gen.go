// retrieve options from interaction data
package corde

// String returns the option with key k of type string
func (o OptionsInteractions) String(k string) string {
	v, _ := o[k].(string)
	return v
}

// Int returns the option with key k of type int
func (o OptionsInteractions) Int(k string) int {
	v, _ := o[k].(int)
	return v
}

// Int64 returns the option with key k of type int64
func (o OptionsInteractions) Int64(k string) int64 {
	v, _ := o[k].(int64)
	return v
}

// Uint returns the option with key k of type uint
func (o OptionsInteractions) Uint(k string) uint {
	v, _ := o[k].(uint)
	return v
}

// Uint64 returns the option with key k of type uint64
func (o OptionsInteractions) Uint64(k string) uint64 {
	v, _ := o[k].(uint64)
	return v
}

// Float32 returns the option with key k of type float32
func (o OptionsInteractions) Float32(k string) float32 {
	v, _ := o[k].(float32)
	return v
}

// Float64 returns the option with key k of type float64
func (o OptionsInteractions) Float64(k string) float64 {
	v, _ := o[k].(float64)
	return v
}

// Bool returns the option with key k of type bool
func (o OptionsInteractions) Bool(k string) bool {
	v, _ := o[k].(bool)
	return v
}

// Snowflake returns the option with key k of type Snowflake
func (o OptionsInteractions) Snowflake(k string) Snowflake {
	v, _ := o[k].(Snowflake)
	return v
}
