package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	schema "github.com/gorilla/Schema"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/volatiletech/sqlboiler/boil"
)

type Validator interface {
	Validate() error
}

type TransactionExecutorParams struct {
	ctx  context.Context
	conn boil.ContextExecutor
	r    *http.Request
	w    http.ResponseWriter
}

func (p *TransactionExecutorParams) SetHeaders(values map[string]string) {
	for k, v := range values {
		p.w.Header().Set(k, v)
	}
}

func (p *TransactionExecutorParams) Param(name string) string {
	return helpers.QueryStringValue(p.r, name)
}

func (p *TransactionExecutorParams) ParamBool(name string) *bool {
	value := p.Param(name)
	val, err := strconv.ParseBool(value)
	if err != nil {
		return nil
	}
	return &val
}

func (p *TransactionExecutorParams) DecodeQuery(v interface{}) error {
	decoder := schema.NewDecoder()
	// should not ignore wrong query string fields by default
	//decoder.IgnoreUnknownKeys(true)
	if err := decoder.Decode(v, p.r.URL.Query()); err != nil {
		return err
	}
	if validator, ok := v.(Validator); ok {
		return validator.Validate()
	}
	return nil
}

func (p *TransactionExecutorParams) Decode(v interface{}) error {
	if p.r.Body == nil {
		return errors.New("missing body, cannot decode")
	}
	return json.NewDecoder(p.r.Body).Decode(v)
}

type TransactionExecutor func(TransactionExecutorParams) (status int, result interface{}, err error)

type PromiseExecutorParameter struct {
	ctx    context.Context
	status int
	result interface{}
	err    error
}

// PromiseExecutor should be used for post db operations like publishing to a message broker
type PromiseExecutor func(parameter PromiseExecutorParameter)
