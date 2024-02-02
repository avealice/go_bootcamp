package db

import (
	"bytes"
	"context"
	"day03/internal/db/csvreader"
	"day03/internal/types"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

func (es ElasticsearchStore) CreateIndex(query string) {
	var (
		res *esapi.Response
		err error
	)

	if res, err = es.client.Indices.Delete([]string{es.indexName}, es.client.Indices.Delete.WithIgnoreUnavailable(true)); err != nil || res.IsError() {
		log.Fatalf("Cannot delete index: %s", err)
	}

	res.Body.Close()

	res, err = es.client.Indices.Create(es.indexName, es.client.Indices.Create.WithBody(strings.NewReader(query)))

	if err != nil {
		log.Fatalf("Cannot create index: %s", err)
	}
	if res.IsError() {
		log.Fatalf("Cannot create index: %s", res)
	}

	res.Body.Close()
}

func (es *ElasticsearchStore) BulkIndex(csvPath string) {
	var (
		places          []types.Place
		countSuccessful uint64
	)

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:  es.indexName,
		Client: es.client,
	})

	if err != nil {
		log.Fatalf("Error creating the indexer: %s", err)
	}

	places = csvreader.LoadDataFromCSV(csvPath)

	start := time.Now().UTC()

	for _, p := range places {
		data, err := json.Marshal(p)
		if err != nil {
			log.Fatalf("Cannot encode place %s: %s", p.Name, err)
		}
		err = bi.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: strconv.Itoa(p.ID),
				Body:       bytes.NewReader(data),
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					atomic.AddUint64(&countSuccessful, 1)
				},
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						log.Printf("ERROR: %s", err)
					} else {
						log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
		if err != nil {
			log.Fatalf("Unexpected error: %s", err)
		}
	}

	if err := bi.Close(context.Background()); err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}

	biStats := bi.Stats()

	log.Println(strings.Repeat("▔", 65))

	dur := time.Since(start)

	if biStats.NumFailed > 0 {
		log.Fatalf(
			"Indexed [%s] documents with [%s] errors in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			humanize.Comma(int64(biStats.NumFailed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
	} else {
		log.Printf(
			"Sucessfuly indexed [%s] documents in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
	}

	log.Println(strings.Repeat("▁", 65))
}
