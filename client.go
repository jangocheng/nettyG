package gonet

import "net"

type Client struct {

}

func NewClient() *Client{
	return  &Client{}
}

func (c *Client) Connect(proto string,addr string){
	conn,err:= net.Dial(proto,addr)
	if err!=nil{
		return
	}

	r:=&Reactor{pipeline:NewPipeline()}
	r.el.run(r.pipeline)

	go handle(conn,r)
}