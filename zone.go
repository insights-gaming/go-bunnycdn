package bunnycdn

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type StorageZone struct {
	c    *Client
	name string
}

func (z *StorageZone) url(path, fileName string) string {
	return fmt.Sprintf("https://storage.bunnycdn.com/%s/%s/%s", z.name, path, fileName)
}

func (z *StorageZone) catch(res *http.Response) error {
	defer res.Body.Close()

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(res.Body); err != nil {
		return err
	}

	var err Error
	if err := json.Unmarshal(buf.Bytes(), &err); err != nil {
		return &UnexpectedResponseError{
			Code:   res.StatusCode,
			Status: res.Status,
			Body:   buf.Bytes(),
		}
	}

	return &err
}

func (z *StorageZone) Name() string {
	return z.name
}

func (z *StorageZone) Get(ctx context.Context, path, fileName string) (io.ReadCloser, error) {
	req, err := http.NewRequest(http.MethodGet, z.url(path, fileName), nil)
	if err != nil {
		return nil, err
	}

	res, err := z.c.hc.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusOK:
		return res.Body, nil
	case http.StatusNotFound:
		defer res.Body.Close()

		return nil, ErrNotFound
	}

	return nil, z.catch(res)
}

func (z *StorageZone) Put(ctx context.Context, path, fileName string, r io.Reader) error {
	req, err := http.NewRequest(http.MethodPut, z.url(path, fileName), r)
	if err != nil {
		return err
	}

	res, err := z.c.hc.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 201:
		return nil
	case 400:
		return ErrUploadFailed
	}

	return z.catch(res)
}

func (z *StorageZone) Delete(ctx context.Context, path, fileName string) error {
	req, err := http.NewRequest(http.MethodDelete, z.url(path, fileName), nil)
	if err != nil {
		return err
	}

	res, err := z.c.hc.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 200:
		return nil
	case 400:
		return ErrObjectDeleteFailed
	}

	return z.catch(res)
}