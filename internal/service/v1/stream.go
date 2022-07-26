//Author: Ivan Kistochkin
//Desc: Сервис для обменивания SDP между стримером и зрителем

package servicev1

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
		mx sync.Mutex

		requestOffer chan struct{}

		receivedOffers  chan *StreamMessage
		receivedAnswers chan *StreamMessage
	}
)

func newStreamService() *streamService {
	return &streamService{
		requestOffer:    make(chan struct{}),
		receivedOffers:  make(chan *StreamMessage),
		receivedAnswers: make(chan *StreamMessage),
	}
}

func (s *streamService) getOffer() *StreamMessage {
	s.requestOffer <- struct{}{}
	return <-s.receivedOffers
}

func (s *streamService) sendAnswer(a *StreamMessage) {
	s.receivedAnswers <- a
}

func (s *streamService) Create(ctx context.Context, owner StreamConnection) {
	if ok := s.mx.TryLock(); !ok {
		owner.WriteError(errors.New("stream already running"))
	}
	defer s.mx.Unlock()

	defer func() {
		close(s.requestOffer)
		s.requestOffer = make(chan struct{})

		close(s.receivedAnswers)
		s.receivedAnswers = make(chan *StreamMessage)
	}()

	//отправитель запросов на получение WebRTC offer
	go func() {
		for range s.requestOffer {
			err := owner.Write(NewStreamMessage("", "", ""))
			if err != nil {
				log.Println(err)
			}
		}
	}()

	//отправляет WebRTC answer
	go func() {
		for answer := range s.receivedAnswers {
			err := owner.Write(answer)
			if err != nil {
				log.Println(err)
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
	offer := s.getOffer()

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
