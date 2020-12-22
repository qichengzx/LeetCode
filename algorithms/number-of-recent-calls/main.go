package main

type RecentCounter struct {
	t []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.t = append(this.t, t)
	threshold := t - 3000
	for this.t[0] < threshold {
		this.t = this.t[1:]
	}

	return len(this.t)
}