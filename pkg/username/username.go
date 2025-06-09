package username

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gzipchrist/dont_at_me/pkg/social"
)

func CheckAvailabilitySerial(username string) error {
	for _, p := range social.Platforms {
		status := p.GetAvailability(username)
		fmt.Printf("    %s%s%s\n", p, strings.Repeat(" ", p.Spacer()), status.String())
	}

	return nil
}

func CheckAvailabilityConcurrent(username string) {
	wg := sync.WaitGroup{}
	results := make(chan string)

	for i := 0; i < len(social.Platforms); i++ {
		wg.Add(1)
		go func(p social.Platform) {
			defer wg.Done()
			status := p.GetAvailability(username)
			results <- fmt.Sprintf("%s %s%s (%s%s)\n", status, p, strings.Repeat(" ", p.Spacer()), p.URL, username)
		}(social.Platforms[i])
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("    %s", result)
	}

	return
}
