package log

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
)

// Config  log config
type Config struct {
	FileName         string        `yaml:"FileName"`
	LinkName         string        `yaml:"LinkName"`
	Level            zerolog.Level `yaml:"Level"`
	MaxAgeDay        int           `yaml:"MaxAgeDay"`
	RotationTimeHour int           `yaml:"RotationTimeHour"`
}

// Init  .
func (p *Config) Init(ctx context.Context) error {
	options := make([]rotatelogs.Option, 0, 6)
	options = append(options,
		rotatelogs.WithMaxAge(time.Hour*time.Duration(24*p.MaxAgeDay)),
		rotatelogs.WithRotationTime(time.Hour*time.Duration(p.RotationTimeHour)),
		rotatelogs.ForceNewFile(),
	)
	if p.LinkName != "" {
		options = append(options, rotatelogs.WithLinkName(p.LinkName))
	}

	logFile, err := rotatelogs.New(p.FileName, options...)
	if err != nil {
		return fmt.Errorf("open or create log file[%s] failed, err: %v", p.FileName, err)
	}

	logger := zerolog.Ctx(ctx).Level(p.Level).With().Timestamp().Logger()

	switch logger.GetLevel() {
	case zerolog.DebugLevel, zerolog.TraceLevel:
		logger = logger.Output(io.MultiWriter(logFile, os.Stdout)).With().Logger()
	default:
		logger = logger.Output(logFile)
	}

	Logger = logger

	return nil
}
