package main

import "log"

func Run() {
	provider := GetProvider()
	if provider == nil {
		log.Fatalf("failed to get provider: %s", conf.Provider)
	}

	addr, err := provider.Query()

}
