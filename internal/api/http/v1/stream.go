package apictrlv1

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebsocketMessageData struct {
	BroadcastID string `json:"broadcastID"`
	Type        string `json:"type"`
	SDP         string `json:"sdp"`
}

type StreamController struct {
	mx sync.Mutex

	upgrader  *websocket.Upgrader
	ownerConn *websocket.Conn

	offersChan       chan *WebsocketMessageData
	requestOfferChan chan struct{}
}

func NewStreamController() *StreamController {
	controller := &StreamController{
		upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			}},

		offersChan:       make(chan *WebsocketMessageData),
		requestOfferChan: make(chan struct{}),
	}

	return controller
}

func (c *StreamController) getOwnerOffer() *WebsocketMessageData {
	c.requestOfferChan <- struct{}{}
	return <-c.offersChan
}

func (c *StreamController) getWatcherAnswer(watcherConn *websocket.Conn, offer *WebsocketMessageData) *WebsocketMessageData {
	watcherConn.WriteJSON(newResponse(offer, nil))

	var answer WebsocketMessageData
	if err := watcherConn.ReadJSON(&answer); err != nil {
		watcherConn.WriteJSON(newResponse(nil, err))
		return nil
	}

	return &answer
}

func (c *StreamController) WSCreate(ctx *gin.Context) {
	conn, err := c.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newResponse(nil, err))
		return
	}
	defer conn.Close()

	if !c.mx.TryLock() {
		conn.WriteJSON(newResponse(nil, fmt.Errorf("stream already exists")))
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

	go func() {
		for range c.requestOfferChan {
			//send empty object `{}` to get offer
			conn.WriteJSON(newResponse(gin.H{}, nil))
		}
	}()

	for {
		var offer WebsocketMessageData

		if err := c.ownerConn.ReadJSON(&offer); err != nil {
			log.Println(err)
			return
		}

		c.offersChan <- &offer
	}
}

func (c *StreamController) WSWatch(ctx *gin.Context) {
	conn, err := c.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newResponse(nil, err))
		return
	}
	defer conn.Close()

	offer := c.getOwnerOffer()

	if answer := c.getWatcherAnswer(conn, offer); answer != nil {
		if c.ownerConn != nil {
			msg := newResponse(WebsocketMessageData{
				BroadcastID: answer.BroadcastID,
				Type:        answer.Type,
				SDP:         answer.SDP,
			}, nil)

			c.ownerConn.WriteJSON(msg)
		}
	}
}
