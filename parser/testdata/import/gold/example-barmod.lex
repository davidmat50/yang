module "module"
[ident] "example-ba"...
{ "{"
namespace "namespace"
[string] "\"http://ex"...
; ";"
prefix "prefix"
[string] "\"barmod\""
; ";"
import "import"
[ident] "example-fo"...
{ "{"
prefix "prefix"
[string] "\"foomod\""
; ";"
} "}"
augment "augment"
[string] "\"/foomod:t"...
{ "{"
leaf "leaf"
[ident] "bar"
{ "{"
type "type"
[ident] "boolean"
; ";"
} "}"
} "}"
} "}"
