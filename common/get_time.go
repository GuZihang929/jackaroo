package common

import "time"

// int64转换成时间
func GetTime(sou int64) time.Time {
	ts := sou
	if ts < 1e10 { //如果是秒
		return time.Unix(ts, 0)
	} else if ts < 1e14 { // 如果ts是豪秒，则将其转换为纳秒
		return time.Unix(0, ts*1e6)
	} else if ts < 1e16 { // 如果ts是微秒，则将其转换为纳秒
		return time.Unix(0, ts*1e3)
	} else if ts < 9e18 { //纳秒
		return time.Unix(0, ts)
	}
	return time.Time{}
}
