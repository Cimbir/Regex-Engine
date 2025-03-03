package parseregex

import "fmt"

type PointType uint8

const (
	MATCHING   PointType = iota
	ANCHOR     PointType = iota
	QUANTIFIER PointType = iota
	GROUP      PointType = iota
	OR         PointType = iota
	LITERAL    PointType = iota
)

type ParsedPoint struct {
	Type  PointType
	Value interface{}
}

type ParsedRegex struct {
	Pos    int
	Points []ParsedPoint
}

func ParseRegex(regex string) (ParsedRegex, error) {
	parsedRegex := ParsedRegex{0, make([]ParsedPoint, 0)}
	for parsedRegex.Pos < len(regex) {
		err := parseCharacter(regex, &parsedRegex)
		if err != nil {
			return parsedRegex, err
		}
	}
	return parsedRegex, nil
}

func parseCharacter(regex string, parsed *ParsedRegex) error {
	char := regex[parsed.Pos]
	if IsAnchor(char) {
		// Parse anchor
	} else if IsQuantifier(char) {
		// Parse quantifier
	} else if IsGroup(char) {
		// Parse group
	} else if IsOr(char) {
		// Parse or
	} else if IsMatching(char) {
		// Parse matching
	} else if IsLiteral(char) {
		// Parse literal
	} else {
		return fmt.Errorf("Invalid character: %c", char)
	}
	return nil
}
