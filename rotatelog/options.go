package rotatelog

import (
	"fmt"
	"path/filepath"
	"time"
)

type Option func(*RotateLog)

func WithRotateTime(duration time.Duration) Option {
	return func(r *RotateLog) {
		r.rotateTime = duration
	}
}

func WithCurLogLinkPath(linkPath string) Option {
	return func(r *RotateLog) {
		r.curLogLinkPath = linkPath
	}
}

// WithDeleteExpiredFile Judge expired by last modify time
// cutoffTime = now - maxAge
// Only delete satisfying file wildcard filename
func WithDeleteExpiredFile(maxAge time.Duration, fileWildcard string) Option {
	return func(r *RotateLog) {
		r.maxAge = maxAge
		r.delFileWildcard = fmt.Sprintf("%s%s%s", filepath.Dir(r.logPath), string([]byte{filepath.Separator}), fileWildcard)
	}
}
