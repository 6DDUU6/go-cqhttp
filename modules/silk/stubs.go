// Package silk Silk编码核心模块
package silk

import (
	"github.com/6DDUU6/go-cqhttp/internal/base"
)

func init() {
	base.EncodeSilk = encode
	base.ResampleSilk = resample
}
