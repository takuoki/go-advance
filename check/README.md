# Check Utils

```go
package sample 
import (
	"github.com/fcfcqloow/go-advance/check"
	"github.com/fcfcqloow/go-advance/convert"
)
func main () {
    
    check.IsNumber(uint8(1))
    //true
    check.IsPtr(convert.StringPtr("aaa"))
    // true
    check.IsNil(nil)
    // true
    check.IsBool(false)
    // true

    _t := T{Test: "test"}
    check.AreEqualJson(t1, convert.CompactJson(t1))
    // true
    check.AreEqualJson(t1, map[string]interface{}{"Test": "test", "Sample": "sample"}) 
    // true
}

```
