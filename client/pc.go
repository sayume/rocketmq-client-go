package client

import (
	"sync"

	"github.com/zjykzk/rocketmq-client-go/route"
)

// producer interface needed by reblance
type producer interface {
	Group() string
	PublishTopics() []string
	UpdateTopicPublish(topic string, router *route.TopicRouter)
	NeedUpdateTopicPublish(topic string) bool
}

type producerColl struct {
	sync.RWMutex
	eles map[string]producer // key: group name, NOTE donot modify directly
}

func (pc *producerColl) coll() []producer {
	pc.RLock()
	coll, i := make([]producer, len(pc.eles)), 0
	for _, p := range pc.eles {
		coll[i] = p
		i++
	}
	pc.RUnlock()
	return coll
}

func (pc *producerColl) putIfAbsent(group string, p producer) (prev producer, suc bool) {
	pc.Lock()
	prev, exist := pc.eles[group]
	if !exist {
		pc.eles[group] = p
		suc = true
	}
	pc.Unlock()
	return
}

func (pc *producerColl) contains(group string) bool {
	pc.RLock()
	_, b := pc.eles[group]
	pc.RUnlock()
	return b
}

func (pc *producerColl) delete(group string) {
	pc.Lock()
	delete(pc.eles, group)
	pc.Unlock()
}

func (pc *producerColl) size() int {
	pc.RLock()
	sz := len(pc.eles)
	pc.RUnlock()
	return sz
}

// RunningInfo consumer running information
type RunningInfo struct {
	Properties    map[string]string `json:"properties"`
	Subscriptions []*Data           `json:"subscriptionSet"`
	// MQTable map[string]*ProcessQueueInfo TODO
	// Statuses map[string]ConsumerStatus TODO
}

// consumer interface needed by reblance
type consumer interface {
	Group() string
	SubscribeTopics() []string
	UpdateTopicSubscribe(topic string, router *route.TopicRouter)
	NeedUpdateTopicSubscribe(topic string) bool
	ConsumeFromWhere() string
	Model() string
	Type() string
	UnitMode() bool
	Subscriptions() []*Data
	ReblanceQueue()
	RunningInfo() RunningInfo
}

type consumerColl struct {
	sync.RWMutex
	eles map[string]consumer // key: group name, NOTE: donot modify directly
}

func (cc *consumerColl) coll() []consumer {
	cc.RLock()
	coll, i := make([]consumer, len(cc.eles)), 0
	for _, c := range cc.eles {
		coll[i] = c
		i++
	}
	cc.RUnlock()
	return coll
}

func (cc *consumerColl) putIfAbsent(group string, c consumer) (prev consumer, suc bool) {
	cc.Lock()
	prev, exist := cc.eles[group]
	if !exist {
		cc.eles[group] = c
		suc = true
	}
	cc.Unlock()
	return
}

func (cc *consumerColl) contains(group string) bool {
	cc.RLock()
	_, b := cc.eles[group]
	cc.RUnlock()
	return b
}

func (cc *consumerColl) get(group string) consumer {
	cc.RLock()
	c := cc.eles[group]
	cc.RUnlock()
	return c
}

func (cc *consumerColl) delete(group string) {
	cc.Lock()
	delete(cc.eles, group)
	cc.Unlock()
}

func (cc *consumerColl) size() int {
	cc.RLock()
	sz := len(cc.eles)
	cc.RUnlock()
	return sz
}

type admin interface {
	Group() string
}

type adminColl struct {
	sync.RWMutex
	eles map[string]admin // key: group name, NOTE: donot modify directly
}

func (ac *adminColl) putIfAbsent(group string, c admin) (prev admin, suc bool) {
	ac.Lock()
	prev, exist := ac.eles[group]
	if !exist {
		ac.eles[group] = c
		suc = true
	}
	ac.Unlock()
	return
}

func (ac *adminColl) delete(group string) {
	ac.Lock()
	delete(ac.eles, group)
	ac.Unlock()
}

func (ac *adminColl) size() int {
	ac.RLock()
	sz := len(ac.eles)
	ac.RUnlock()
	return sz
}
