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


package main

import (
    "github.com/ruymanbr/grapigo/pkg/grapigo_api"
    "github.com/ruymanbr/grapigo/pkg/grapigo_storaging"
)

func main() {
    
    grapigo_storaging.Init()

    grapigo_api.Start()
    
}