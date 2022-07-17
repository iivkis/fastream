package apihv1

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
)

type StreamController struct {
	mx sync.Mutex

	upgrader  *websocket.Upgrader
	ownerConn *websocket.Conn

	offersChan       chan *webrtc.SessionDescription
	requestOfferChan chan struct{}
}

func NewStreamController() *StreamController {
	controller := &StreamController{
		upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			}},

		offersChan:       make(chan *webrtc.SessionDescription),
		requestOfferChan: make(chan struct{}),
	}

	return controller
}

func (c *StreamController) getOwnerOffer() *webrtc.SessionDescription {
	c.requestOfferChan <- struct{}{}
	return <-c.offersChan
}

func (c *StreamController) getWatcherAnswer(watherConn *websocket.Conn, offer *webrtc.SessionDescription) *webrtc.SessionDescription {
	watherConn.WriteJSON(offer)

	var answer webrtc.SessionDescription
	if err := watherConn.ReadJSON(&answer); err != nil {
		watherConn.WriteJSON(NewResponse(nil, err))
		return nil
	}

	return &answer
}

func (c *StreamController) WSCreate(ctx *gin.Context) {
	conn, err := c.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewResponse(nil, err))
		return
	}
	defer conn.Close()

	if !c.mx.TryLock() {
		conn.WriteJSON(NewResponse(nil, fmt.Errorf("stream already exists")))
		return
	}
	defer c.mx.Unlock()

	defer func() {
		c.ownerConn = nil

		close(c.requestOfferChan)
		c.requestOfferChan = make(chan struct{})

		log.Println("owner websocket connection was closed")
	}()

	c.ownerConn = conn

	//send empty object {} to get offer
	go func() {
		for range c.requestOfferChan {
			conn.WriteJSON(gin.H{})
		}
	}()

	//catch offer
	{
		var offer webrtc.SessionDescription

		for {
			if err := c.ownerConn.ReadJSON(&offer); err != nil {
				log.Println(err)
				return
			}

			c.offersChan <- &offer
		}
	}
}

func (c *StreamController) WSWatch(ctx *gin.Context) {
	conn, err := c.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewResponse(nil, err))
		return
	}
	defer conn.Close()

	offer := c.getOwnerOffer()

	if answer := c.getWatcherAnswer(conn, offer); answer != nil {
		if c.ownerConn != nil {
			c.ownerConn.WriteJSON(answer)
		}
	}
}
