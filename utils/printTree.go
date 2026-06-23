package utils

import (
	"fmt"
	"io"
)

func PrintNode(writer io.Writer, node Node, prefix string, isLast bool, isRoot bool) {
	const (
		colorReset = "\033[0m"
		colorBold  = "\033[1m"

		colorRoot   = "\033[38;5;213m" // bright magenta
		colorBranch = "\033[38;5;117m" // sky blue
		colorLeaf   = "\033[38;5;156m" // soft green

		colorChrome = "\033[38;5;240m" // muted gray for connectors

		colorString = "\033[38;5;221m" // amber
		colorNumber = "\033[38;5;159m" // cyan
		colorNil    = "\033[38;5;203m" // red-ish

		connectorMid  = "├── "
		connectorLast = "└── "
		connectorPipe = "│   "
		connectorGap  = "    "

		shapeRoot   = "◈ "
		shapeBranch = "◆ "
		shapeLeaf   = "◇ "
	)

	// ── determine connector and next prefix ──────────────────────────────────
	var connector string
	var nextPrefix string
	if isRoot {
		connector = ""
		nextPrefix = ""
	} else if isLast {
		connector = colorChrome + connectorLast + colorReset
		nextPrefix = prefix + colorChrome + connectorGap + colorReset
	} else {
		connector = colorChrome + connectorMid + colorReset
		nextPrefix = prefix + colorChrome + connectorPipe + colorReset
	}

	// ── determine node role ───────────────────────────────────────────────────
	parent, isParent := node.(Node)
	hasKids := isParent && parent.HasChildren()

	var nodeColor, nodeShape string
	switch {
	case isRoot:
		nodeColor = colorRoot
		nodeShape = shapeRoot
	case hasKids:
		nodeColor = colorBranch
		nodeShape = shapeBranch
	default:
		nodeColor = colorLeaf
		nodeShape = shapeLeaf
	}

	// ── format the value ──────────────────────────────────────────────────────
	val := node.GetValue()
	var valueStr, valueColor string
	switch v := val.(type) {
	case nil:
		valueStr = "nil"
		valueColor = colorNil
	case string:
		valueStr = fmt.Sprintf("%q", v)
		valueColor = colorString
	case int, int8, int32, int64, float32, float64:
		valueStr = fmt.Sprintf("%v", v)
		valueColor = colorNumber
	default:
		valueStr = fmt.Sprintf("%v", v)
		valueColor = colorReset
	}

	// ── count children for annotation ─────────────────────────────────────────
	var annotation string
	if hasKids {
		n := len(parent.GetChildren())
		annotation = fmt.Sprintf(
			" %s[%d %s]%s",
			colorChrome, n,
			pluralize("child", n),
			colorReset,
		)
	}

	// ── print this node inline (Fixes alignment issues completely) ────────────
	fmt.Fprintf(writer,
		"%s%s%s%s%s%s%s%s%s\n",
		prefix,
		connector,
		nodeColor+colorBold, nodeShape, colorReset,
		valueColor, valueStr, colorReset,
		annotation,
	)

	// ── recurse into children ─────────────────────────────────────────────────
	if !hasKids {
		return
	}
	children := parent.GetChildren()
	for i, child := range children {
		PrintNode(writer, child, nextPrefix, i == len(children)-1, false)
	}
}

func pluralize(word string, n int) string {
	if n == 1 {
		return word
	}
	return word + "ren"
}

// stripANSI removes ANSI escape sequences to measure true visible length.
func stripANSI(s string) string {
	result := make([]byte, 0, len(s))
	inEscape := false
	for i := 0; i < len(s); i++ {
		if s[i] == '\033' {
			inEscape = true
			continue
		}
		if inEscape {
			if s[i] == 'm' {
				inEscape = false
			}
			continue
		}
		result = append(result, s[i])
	}
	return string(result)
}
