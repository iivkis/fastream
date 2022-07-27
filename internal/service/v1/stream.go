//Author: Ivan Kistochkin
//Desc: Сервис для обменивания SDP между стримером и зрителем

package service

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/pion/webrtc/v3"
)

type StreamMessage struct {
	BroadcastID string
	Session     *webrtc.SessionDescription
}

func NewStreamMessage(broadcastID, sdpType, sdp string) *StreamMessage {
	return &StreamMessage{
		BroadcastID: broadcastID,
		Session: &webrtc.SessionDescription{
			Type: webrtc.NewSDPType(sdpType),
			SDP:  sdp,
		},
	}
}

type (
	StreamConnection interface {
		Read() (*StreamMessage, error)
		Write(*StreamMessage) error
		WriteError(error) error
	}

	streamService struct {
		mx              sync.Mutex
		requestOffer    chan struct{}
		receivedOffers  chan *StreamMessage
		receivedAnswers chan *StreamMessage
	}
)

func newStreamService() stream {
	return &streamService{
		requestOffer:    make(chan struct{}),
		receivedOffers:  make(chan *StreamMessage),
		receivedAnswers: make(chan *StreamMessage),
	}
}

func (s *streamService) getOffer() chan *StreamMessage {
	s.requestOffer <- struct{}{}
	return s.receivedOffers
}

func (s *streamService) sendAnswer(answer *StreamMessage) {
	s.receivedAnswers <- answer
}

func (s *streamService) Create(ctx context.Context, owner StreamConnection) {
	if ok := s.mx.TryLock(); !ok {
		owner.WriteError(errors.New("stream already running"))
	}
	defer s.mx.Unlock()

	//отправитель запросов на получение WebRTC offer
	go func() {
		for {
			select {
			case <-s.requestOffer:
				err := owner.Write(NewStreamMessage("", "", ""))
				if err != nil {
					log.Println(err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	//отправляет WebRTC answer
	go func() {
		for {
			select {
			case answer := <-s.receivedAnswers:
				err := owner.Write(answer)
				if err != nil {
					log.Println(err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	for {
		offer, err := owner.Read()
		if err != nil {
			log.Println(err)
			return
		}
		s.receivedOffers <- offer
	}
}

func (s *streamService) Watch(ctx context.Context, watcher StreamConnection) {
	next := func(offer *StreamMessage) {
		err := watcher.Write(offer)
		if err != nil {
			log.Println(err)
			return
		}

		answer, err := watcher.Read()
		if err != nil {
			log.Println(err)
			return
		}

		s.sendAnswer(answer)
	}

	select {
	case offer := <-s.getOffer():
		next(offer)
	case <-ctx.Done():
		return
	}
}
