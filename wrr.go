package wrr

import "sync"

type Balancer struct {
	sync.RWMutex
	list map[string]*Item
}

type Item struct {
	key    string
	weight Weight
}

type Weight struct {
	init      int
	current   int
	effective int
}

func Init() *Balancer {
	b := &Balancer{}
	b.list = make(map[string]*Item)
	return b
}

func (b *Balancer) Add(key string, weight int) {
	b.Lock()
	defer b.Unlock()

	b.list[key] = &Item{
		key:    key,
		weight: Weight{init: weight, effective: weight},
	}
}

func (b *Balancer) Next() string {
	b.Lock()
	defer b.Unlock()

	if len(b.list) == 0 {
		return ""
	}

	var total int
	var best *Item
	for _, item := range b.list {
		total += item.weight.effective
		item.weight.current += item.weight.effective
		if best == nil || item.weight.current > best.weight.current {
			best = item
		}
	}
	best.weight.current -= total
	return best.key
}

func (b *Balancer) Remove(key string) {
	b.Lock()
	defer b.Unlock()

	delete(b.list, key)
}

func (b *Balancer) IncWeight(key string) {
	b.Lock()
	defer b.Unlock()

	item := b.list[key]
	if item == nil || item.weight.effective >= item.weight.init {
		return
	}
	item.weight.effective++
}

func (b *Balancer) DecWeight(key string) {
	b.Lock()
	defer b.Unlock()

	item := b.list[key]
	if item == nil || item.weight.effective <= 0 {
		return
	}
	item.weight.effective--
}
