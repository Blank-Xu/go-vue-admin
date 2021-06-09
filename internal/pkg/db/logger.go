package db

import (
	"fmt"

	"github.com/rs/zerolog"
	"xorm.io/xorm/log"
)

type Logger struct {
	logger   *zerolog.Logger
	database string
	level    log.LogLevel
	showSQL  bool
}

func NewLogger(logger *zerolog.Logger, database string, logLevel log.LogLevel) *Logger {
	return &Logger{
		logger:   logger,
		database: database,
		level:    logLevel,
	}
}

// Debug implement log.Logger
func (p *Logger) Debug(v ...interface{}) {
	if p.level <= log.LOG_DEBUG {
		p.logger.Debug().Str("database", p.database).Msg(fmt.Sprintln(v...))
	}
}

// Debugf implement log.Logger
func (p *Logger) Debugf(format string, v ...interface{}) {
	if p.level <= log.LOG_DEBUG {
		p.logger.Debug().Str("database", p.database).Msgf(format, v...)
	}
}

// Error implement log.Logger
func (p *Logger) Error(v ...interface{}) {
	if p.level <= log.LOG_ERR {
		p.logger.Error().Str("database", p.database).Msg(fmt.Sprintln(v...))
	}
}

// Errorf implement log.Logger
func (p *Logger) Errorf(format string, v ...interface{}) {
	if p.level <= log.LOG_ERR {
		p.logger.Error().Str("database", p.database).Msgf(format, v...)
	}
}

// Info implement log.Logger
func (p *Logger) Info(v ...interface{}) {
	if p.level <= log.LOG_INFO {
		p.logger.Info().Str("database", p.database).Msg(fmt.Sprintln(v...))
	}
}

// Infof implement log.Logger
func (p *Logger) Infof(format string, v ...interface{}) {
	if p.level <= log.LOG_INFO {
		p.logger.Info().Str("database", p.database).Msgf(format, v...)
	}
}

// Warn implement log.Logger
func (p *Logger) Warn(v ...interface{}) {
	if p.level <= log.LOG_WARNING {
		p.logger.Warn().Str("database", p.database).Msg(fmt.Sprintln(v...))
	}
}

// Warnf implement log.Logger
func (p *Logger) Warnf(format string, v ...interface{}) {
	if p.level <= log.LOG_WARNING {
		p.logger.Warn().Str("database", p.database).Msgf(format, v...)
	}
}

// Level implement log.Logger
func (p *Logger) Level() log.LogLevel {
	return p.level
}

// SetLevel implement log.Logger
func (p *Logger) SetLevel(l log.LogLevel) {
	p.level = l
	return
}

// ShowSQL implement log.Logger
func (p *Logger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		p.showSQL = true
		return
	}
	p.showSQL = show[0]
}

// IsShowSQL implement log.Logger
func (p *Logger) IsShowSQL() bool {
	return p.showSQL
}
