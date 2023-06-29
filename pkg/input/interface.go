package input

import (
	"github.com/liuxinwang/go-mysql-starrocks/pkg/channel"
	"github.com/liuxinwang/go-mysql-starrocks/pkg/position"
	"regexp"
)

type Plugin interface {
	NewInput(config interface{}, ruleRegex []string)
	StartInput(pos position.Position, syncChan *channel.SyncChannel) position.Position
	StartMetrics()
	Close()
	SetIncludeTableRegex(map[string]interface{}) (*regexp.Regexp, error)    // for add rule
	RemoveIncludeTableRegex(map[string]interface{}) (*regexp.Regexp, error) // for delete rule
}
