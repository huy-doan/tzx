package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/jszwec/csvutil"
	"github.com/test-tzs/nomraeite/internal/pkg/logger"
	"github.com/test-tzs/nomraeite/internal/pkg/utils/messages"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type CSVReaderOptions struct {
	TrimEmptyRows     bool
	CustomDateFormats []string
}

func defaultCSVReaderOptions() *CSVReaderOptions {
	return &CSVReaderOptions{
		TrimEmptyRows: true,
		CustomDateFormats: []string{
			"2006/01/02 15:04", // Standard format with leading zeros
			"2006/1/2 15:04",   // Without leading zeros for month/day
			"2006/1/2 15:4",    // Without leading zeros for month/day/hour
			"2006/1/2 0:0",     // Without leading zeros for month/day and with zero hour/minute
			"2006/01/02 0:00",  // With leading zeros for month/day and with zero hour/minute
			"2006/1/02 0:00",   // Mixed format
			"2006/01/2 0:00",   // Mixed format
		},
	}
}

type CSVReader[T any] struct {
	Reader  io.Reader
	Decoder *csvutil.Decoder
	Options *CSVReaderOptions
	logger  logger.Logger
}

func NewCSVReader[T any](r io.Reader, options *CSVReaderOptions) (*CSVReader[T], error) {
	shiftJisReader := transform.NewReader(r, japanese.ShiftJIS.NewDecoder())
	csvReader := csv.NewReader(shiftJisReader)
	decoder, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		return nil, fmt.Errorf(messages.MsgFailedToCreateCSVDecoder+": %w", err)
	}
	if options == nil {
		options = defaultCSVReaderOptions()
	}
	if len(options.CustomDateFormats) > 0 {
		decoder.WithUnmarshalers(
			csvutil.UnmarshalFunc(func(data []byte, t *time.Time) error {
				if len(data) == 0 {
					return nil
				}
				var parseErr error
				for _, format := range options.CustomDateFormats {
					tt, err := time.Parse(format, string(data))
					if err == nil {
						*t = tt
						return nil
					}
					parseErr = err
				}
				logger.GetLogger().Error(messages.MsgFailedToParseDate+"\n", map[string]any{
					"data":  string(data),
					"error": parseErr,
				})
				return parseErr
			}),
		)
	}
	decoder.Map = func(field, col string, v any) string {
		// Handle 1E+02 notation for floats/integers
		if num, err := strconv.ParseFloat(field, 64); err == nil {
			return strconv.FormatFloat(num, 'f', -1, 64)
		}

		return field
	}
	return &CSVReader[T]{
		Reader:  r,
		Decoder: decoder,
		Options: options,
		logger:  logger.GetLogger(),
	}, nil
}

func (r *CSVReader[T]) ReadEach(callback func(record T) error) error {
	for {
		record := new(T)
		if err := r.Decoder.Decode(record); err != nil {
			if err == io.EOF {
				break
			} else {
				if r.Options.TrimEmptyRows && isEmptyRecord(r.Decoder.Record()) {
					continue
				}
				return fmt.Errorf(messages.MsgFailedToDecodeCSVRecord+": %w", err)
			}
		}
		if err := callback(*record); err != nil {
			return fmt.Errorf(messages.MsgCSVCallbackError+": %w", err)
		}
	}
	return nil
}

func (r *CSVReader[T]) ReadAll() ([]T, error) {
	var records []T
	err := r.ReadEach(func(record T) error {
		records = append(records, record)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CSVReader[T]) HasHeader(headers []string) bool {
	if len(r.Decoder.Header()) == 0 {
		return false
	}
	headersMap := make(map[string]bool)
	for _, header := range r.Decoder.Header() {
		headersMap[header] = true
	}
	for _, header := range headers {
		if _, exists := headersMap[header]; !exists {
			r.logger.Error(fmt.Sprintf("Header '%s' not found in CSV headers: %v\n", header, r.Decoder.Header()), nil)
			return false
		}
	}
	return true
}

func (r *CSVReader[T]) Header() []string {
	return r.Decoder.Header()
}

func isEmptyRecord(record []string) bool {
	for _, field := range record {
		if field != "" {
			return false
		}
	}
	return true
}
