package fnapiio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ShopOpts struct {
	// Language for the response.
	//
	// API's default: english
	Lang Lang

	// Set to true if you want to get the render data for each display asset.
	//
	// API's default: false
	IncludeRenderData bool

	// Allows you to customize the list of fields to return for
	// each cosmetics. (comma separated)
	//
	// API's default:
	Fields string

	// Date of the shop (available from Jan 1st 2022) - Format: yyyy-mm-dd
	//
	// API's default:
	Date string
}

// Shop fetches /v2/shop endpoint and returns the deserialized response.
func (c *Client) Shop(opts ShopOpts) (*Shop, error) {
	// use the default language if nothing's set.
	if opts.Lang == LangNone {
		opts.Lang = c.lang
	}

	// initialize the http request
	req, err := http.NewRequest(http.MethodGet, Host+"/v2/shop", nil)
	if err != nil {
		return nil, err
	}

	// set the query parameters
	query := req.URL.Query()
	query.Set("lang", string(opts.Lang))
	query.Set("includeRenderData", strconv.FormatBool(opts.IncludeRenderData))
	query.Set("fields", opts.Fields)
	query.Set("date", opts.Date)
	req.URL.RawQuery = query.Encode()

	// set the authorization header
	req.Header.Set("Authorization", c.token)

	// Send the http request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// ToDo: handle the errors.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch > status code %d", resp.StatusCode)
	}

	shop := &Shop{}

	// deserialize the response.
	if err = json.NewDecoder(resp.Body).Decode(shop); err != nil {
		return nil, err
	}
	return shop, nil
}
