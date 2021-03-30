package sample1

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/multierr"
)

// PriceService is a service that we can use to get prices for the items
// Calls to this service are expensive (they take time)
type PriceService interface {
	GetPriceFor(itemCode string) (float64, error)
}

// TransparentCache is a cache that wraps the actual service
// The cache will remember prices we ask for, so that we don't have to wait on every call
// Cache should only return a price if it is not older than "maxAge", so that we don't get stale prices
type TransparentCache struct {
	actualPriceService PriceService
	maxAge             time.Duration
	prices             sync.Map
}

// NewTransparentCache returns a new TransparentCache.
func NewTransparentCache(actualPriceService PriceService, maxAge time.Duration) *TransparentCache {
	return &TransparentCache{
		actualPriceService: actualPriceService,
		maxAge:             maxAge,
	}
}

// GetPriceFor gets the price for the item, either from the cache or the actual service if it was not cached or too old
func (c *TransparentCache) GetPriceFor(itemCode string) (float64, error) {
	price, ok := c.load(itemCode)
	if ok {
		return price, nil
	}
	price, err := c.actualPriceService.GetPriceFor(itemCode)
	if err != nil {
		return 0, fmt.Errorf("getting price from service : %v", err.Error())
	}
	c.store(price, itemCode)
	return price, nil
}

// GetPricesFor gets the prices for several items at once, some might be found in the cache, others might not
// If any of the operations returns an error, it should return an error as well
func (c *TransparentCache) GetPricesFor(itemCodes ...string) ([]float64, error) {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)
	errors := []error{}
	results := []float64{}

	wg.Add(len(itemCodes))
	for _, itemCode := range itemCodes {
		go func(itemCode string) {
			defer wg.Done()

			price, err := c.GetPriceFor(itemCode)

			// using mutex to avoid race condition
			mu.Lock()
			defer mu.Unlock()
			if err != nil {
				errors = append(errors, err)
				return
			}
			results = append(results, price)
		}(itemCode)
	}
	wg.Wait()

	if len(errors) != 0 {
		return []float64{}, multierr.Combine(errors...)
	}
	return results, nil
}

// load the price and checks if was retrieved less than "maxAge" ago!
func (c *TransparentCache) load(itemCode string) (float64, bool) {
	item, ok := c.prices.Load(itemCode)
	if ok {
		price := item.(Price)
		if price.IsValid(c.maxAge) {
			return price.Value(), true
		}
	}
	return 0.0, false
}

// store the price in the prices mapping
func (c *TransparentCache) store(price float64, itemCode string) {
	c.prices.Store(itemCode, NewPrice(price))
}
