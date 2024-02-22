package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/google/generative-ai-go/genai"
)

type Session struct {
	Sender      int64
	ChatSession *genai.ChatSession
}

type Sessions struct {
	m  map[int64]*Session
	mu sync.Mutex
}

func (s *Sessions) add(session *Session) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.m == nil {
		s.m = map[int64]*Session{}
	}
	s.m[session.Sender] = session
}

func (s *Sessions) get(sender int64) *Session {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.m[sender]
}

func (s *Sessions) delete(sender int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, sender)
}

var sessions = &Sessions{}

func (sessions *Sessions) Ask(sender int64, content string) (string, error) {
	session := sessions.get(sender)
	if session == nil {
		session = &Session{
			Sender:      sender,
			ChatSession: gemini().StartChat(),
		}
		sessions.add(session)
	}
	if len(session.ChatSession.History) > 11 {
		session.ChatSession.History = session.ChatSession.History[:10]
	}
	resp, err := session.ChatSession.SendMessage(context.Background(), genai.Text(content))
	log.Printf("[gemini] ask:%s, resp:%+v, err:%+v\n", content, resp, err)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", resp.Candidates[0].Content.Parts[0]), nil
}
