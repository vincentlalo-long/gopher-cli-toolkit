package models

import "strings"

// FilterConfig
type FilterConfig struct {
	Root    string
	Ext     string
	Query   string
	MinSize int64
}

// Check match condition
func (c *FilterConfig) Matches(fileName string, fileSize int64, fileExt string) bool {
	// filter extension
	if c.Ext != "" && fileExt != c.Ext {
		return false
	}
	// minSize filter
	if fileSize < c.MinSize {
		return false
	}
	// Key name filter
	if c.Query != "" && !contains(fileName, c.Query) {
		return false
	}
	return true
}

func contains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
