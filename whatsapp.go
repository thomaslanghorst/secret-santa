package main

import (
	"errors"
	"fmt"
	"time"

	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/Rhymen/go-whatsapp"
)

type WhatsAppClientInterface interface {
	Login() error
	SendMessage(number, message string) error
	DeleteMessage(number, msgId string) error
	Logout() error
}

type WhatsAppClient struct {
	major int
	minor int
	patch int
	conn  *whatsapp.Conn
}

func NewWhatsAppClient(major, minor, patch int) *WhatsAppClient {
	return &WhatsAppClient{
		major: major,
		minor: minor,
		patch: patch,
	}
}

func (wac *WhatsAppClient) Login() error {

	conn, err := whatsapp.NewConn(5 * time.Second)
	conn.SetClientVersion(wac.major, wac.minor, wac.patch)
	if err != nil {
		fmt.Printf("error creating connection: %s\n", err.Error())
		return errors.New("error creating connection")
	}

	qr := make(chan string)
	go func() {
		terminal := qrcodeTerminal.New()
		terminal.Get(<-qr).Print()
	}()

	_, err = conn.Login(qr)
	if err != nil {
		fmt.Printf("error logging in: %s\n", err.Error())
		return errors.New("error logging in")
	}

	wac.conn = conn

	return nil
}

func (wac *WhatsAppClient) SendMessage(number, message string) (string, error) {

	if wac.conn == nil {
		return "", errors.New("not logged in")
	}

	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: fmt.Sprintf("%s@s.whatsapp.net", number),
		},
		Text: message,
	}

	msgId, err := wac.conn.Send(msg)
	if err != nil {
		fmt.Printf("error sending message: %s", err.Error())
		return "", errors.New("error sending message")
	}

	return msgId, nil
}

func (wac *WhatsAppClient) DeleteMessage(number, msgId string) error {
	err := wac.conn.DeleteMessage(fmt.Sprintf("%s@s.whatsapp.net", number), msgId, true)
	if err != nil {
		fmt.Printf("error deleting message: %s", err.Error())
		return errors.New("error deleting message")
	}

	return nil
}

func (wac *WhatsAppClient) Logout() error {
	err := wac.conn.Logout()
	if err != nil {
		fmt.Printf("error logging out: %s", err.Error())
		return errors.New("error logging out")
	}

	return nil
}
