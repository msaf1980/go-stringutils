# stringutils

Some string utils

`UnsafeString(ptr *byte, length int) string` return unsafe string from bytes slice pointer indirectly (without allocation)

`Builder` very simular to strings.Builder, but has better perfomance (at golang 1.14).
