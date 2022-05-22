/*
 * File:    main.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	21-5-22
 *
 * Summary of File:
 *
 *  This program runs a service that exposes an graphQL API which allows people to create, update, delete and retrieve answers as key-value pairs.
 * 	The answers are stored in a persistent storage medium so they can handle service restarts.
 *  
 *
 */


/*
 * If a user saves the same key multiple times (using update), every answer should be saved. When retrieving an answer, it should return the latest answer.
 *
 * If a user tries to create an answer that already exists - the request should fail and an adequate message or code should be returned.
 *
 * If an answer doesn't exist or has been deleted, an adequate message or code should be returned.
 *
 * When returning history, only mutating events (create, update, delete) should be returned. The "get" events should not be recorded.
 *
 * It is possible to create a key after it has been deleted. However, it is not possible to update a deleted key. For example the following event sequences are allowed:
 *
 * create → delete → create → update
 *
 * create → update → delete → create → update
 *
 * However, the following should not be allowed:
 *
 * create → delete → update
 *
 * create → create
 * 
 * Additional questions:
 *
 * How would you support multiple users?
 * How would you support answers with types other than string
 * What are the main bottlenecks of your solution?
 * How would you scale the service to cope with thousands of requests?
 *
 */

package main

import (
    "github.com/ruymanbr/grapigo/pkg/grapigo_api"
    "github.com/ruymanbr/grapigo/pkg/grapigo_storaging"
)

func main() {
    
    grapigo_storaging.Init()

    grapigo_api.Start()
    
}