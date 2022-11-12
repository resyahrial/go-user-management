package exception

import (
	"fmt"
	"log"
	"strings"
)

const (
	BaseModule = "BASE"
)

type Base struct {
	Code              int
	Message           string
	Module            string
	CollectionMessage map[string][]string
}

func (exc *Base) Error() string {
	if exc.CollectionMessage != nil && len(exc.CollectionMessage) != 0 {
		return fmt.Sprintf("%v", exc.CollectionMessage)
	}
	return exc.Message
}

func (exc *Base) LogError() {
	if exc.CollectionMessage != nil && len(exc.CollectionMessage) != 0 {
		log.Printf("[%s] %v", exc.Module, exc.CollectionMessage)
		return
	}
	log.Printf("[%s] %s", exc.Module, exc.Message)
}

func (exc *Base) SetMessage(msg string) *Base {
	exc.Message = msg
	exc.CollectionMessage = nil
	return exc
}

func (exc *Base) SetCollectionMessage(msg map[string][]string) *Base {
	exc.Message = ""
	exc.CollectionMessage = msg
	return exc
}

func (exc *Base) SetModule(moduleName string) *Base {
	if strings.TrimSpace(moduleName) != "" {
		exc.Module = moduleName
	}
	return exc
}
