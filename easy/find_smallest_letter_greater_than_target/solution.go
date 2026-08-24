// https://leetcode.com/problems/find-smallest-letter-greater-than-target
package findsmallestlettergreaterthantarget

func nextGreatestLetter(letters []byte, target byte) byte {
	const bigger = byte('z' + 1)
	smallerGreater := bigger
	lowerLetter := letters[0]
	for _, l := range letters {
		if l < smallerGreater && l > target {
			smallerGreater = l
		}
		if l < lowerLetter {
			lowerLetter = l
		}
	}

	if bigger == smallerGreater {
		return lowerLetter
	}

	return smallerGreater
}

var NextGreatestLetter func(letters []byte, target byte) byte = nextGreatestLetter
