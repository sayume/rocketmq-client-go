package consumer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zjykzk/rocketmq-client-go/log"
	"github.com/zjykzk/rocketmq-client-go/message"
)

type mockMessagePuller struct {
	runPull bool
}

func (p *mockMessagePuller) pull(r *pullRequest) { p.runPull = true }

func TestNewPullService(t *testing.T) {
	_, err := newPullService(pullServiceConfig{})
	assert.NotNil(t, err)
	_, err = newPullService(pullServiceConfig{messagePuller: &mockMessagePuller{}})
	assert.NotNil(t, err)
	ps, err := newPullService(pullServiceConfig{
		messagePuller: &mockMessagePuller{},
		logger:        log.MockLogger{},
	})
	assert.Nil(t, err)
	assert.NotNil(t, ps)
	assert.Equal(t, defaultRequestBufferSize, ps.requestBufferSize)
}

func TestPullService(t *testing.T) {
	ps, err := newPullService(pullServiceConfig{
		messagePuller: &mockMessagePuller{},
		logger:        log.MockLogger{},
	})
	if err != nil {
		t.Fatal(err)
	}

	r := &pullRequest{
		messageQueue: &message.Queue{},
	}

	ps.submitRequestImmediately(r)
	count := func() int {
		c := 0
		ps.queuesOfMessageQueue.Range(func(_, _ interface{}) bool {
			c++
			return true
		})
		return c
	}
	assert.Equal(t, 1, count())
	ps.submitRequestImmediately(r)
	assert.Equal(t, 1, count())

	ps.shutdown()
}
