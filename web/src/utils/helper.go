package utils

import "strings"

// matchesApplicable prüft, ob deviceVersion zu applicableVersions passt.
// Regeln:
// - "", "all"  → immer true
// - Kommagetrennte Liste → irgendein Eintrag muss passen
// - Optional: Sternchen * als Wildcard (Prefix/Suffix/Enthält)
func MatchesApplicable(deviceVersion string, applicableVersions string) bool {
	av := strings.TrimSpace(strings.ToLower(applicableVersions))
	dv := strings.TrimSpace(strings.ToLower(deviceVersion))

	if av == "" || av == "all" {
		return true
	}

	for _, token := range strings.Split(av, ",") {
		pat := strings.TrimSpace(strings.ToLower(token))
		if pat == "" {
			continue
		}
		// Exakte Übereinstimmung
		if pat == dv {
			return true
		}
		// Einfache Wildcards erlauben (optional):
		// "*x", "x*", "*x*", "x*y" etc.
		if wildcardMatch(dv, pat) {
			return true
		}
	}
	return false
}

// wildcardMatch unterstützt '*' als Wildcard.
// Beispiel: "3.*" matcht "3.2.1", "*rc" matcht "1.0rc", "*3.2*" matcht "v3.2-build42".
func wildcardMatch(s, pattern string) bool {
	// Kein Wildcard? dann exakter Vergleich
	if !strings.Contains(pattern, "*") {
		return s == pattern
	}

	// Aufsplitten an '*', alle Teile müssen in s in Reihenfolge vorkommen
	parts := strings.Split(pattern, "*")
	idx := 0
	// Falls Pattern nicht mit '*' beginnt, muss s mit erstem Teil beginnen
	if parts[0] != "" {
		if !strings.HasPrefix(s, parts[0]) {
			return false
		}
		idx = len(parts[0])
	}

	// Mittelteile: der Reihe nach finden
	for i := 1; i < len(parts)-1; i++ {
		p := parts[i]
		if p == "" {
			continue
		}
		j := strings.Index(s[idx:], p)
		if j < 0 {
			return false
		}
		idx += j + len(p)
	}

	// Letzter Teil: falls Pattern nicht mit '*' endet, muss s damit enden
	last := parts[len(parts)-1]
	if last != "" {
		return strings.HasSuffix(s, last)
	}
	return true
}

