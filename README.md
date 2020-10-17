# stringutils

Some string utils

`UnsafeString([]byte) string` return unsafe string from bytes slice indirectly (without allocation)
`UnsafeStringFromPtr(*byte, length) string` return unsafe string from bytes slice pointer indirectly (without allocation)
`UnsafeStringBytes(*string) []bytes` return unsafe string bytes indirectly (without allocation)

`Builder` very simular to strings.Builder, but has better perfomance (at golang 1.14).
