package play

func approach(tar, curr, step float32) float32 {
	if curr > tar {
		curr -= step
		if curr < tar {
			return tar
		}
	}
	if curr < tar {
		curr += step
		if curr > tar {
			return tar
		}
	}
	return curr
}
