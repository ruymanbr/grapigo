/*
 * File:    grapigo_api.go
 *
 * Author:      Ruymán Borges R. (ruyman21@gmail.com)
 * Date:        22-05-22
 * 
 * Version:     1.00
 * VersionDate: 22-05-22
 * 
 *
 * Summary of File:
 *
 *  This program handles requests using a graphQL environment
 *  It also interacts with grapigo_storaging pkg to store data and retrieve 
 *  previously saved data in a persistant way
 *  
 *  API documentation is also provided to operate it.
 * 
 *  It runs on a unix system (i.e. ubuntu)
 *
 *  Copyright (C) Ruymán Borges Rodríguez - All Rights Reserved
 *  Unauthorized copying of this file, via any medium is strictly prohibited
 *  Proprietary and confidential
 *  Written by Ruymán Borges Rodríguez <ruyman21@gmail.com>, May 2022
 */

package grapigo_api

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
    "github.com/graphql-go/graphql"
)

/*
An answer can be defined as:

key: string
value: string
e.g. in JSON:

{
    "key" : "name",
    "value" : "John"
}
*/

type Answers struct {
    AStruct         []Answer        `json:"ans_struct_array"`
}

type Answer struct {
    Key             string          `json:"ans_key"`
    Value           string          `json:"ans_value"`
    Answer_events   []Event         `json:"ans_event_history"`
}

/* 
An event can be defined as:
event: string
data: answer
e.g. in JSON:

{
    "event" : "create",
    "data" : {
        "key": "name",
        "value": "John"
    }
}
*/

type Event struct {
    Event           string          `json:"event_action"`
    Data            Answer          `json:"event_data"`
    Time            time.Time       `json:"event_time"`
}

const (
    EmptyStr string = ""
)

// loadGraphAPI sets up a graphQL environment to handle server requests
func loadGraphAPI() {
    // ... code that properly sets up a graphQL environment and variables to handle requests
    
}

// Start initiates graphQL API
func Start() {

    loadGraphAPI()

    // Load Graph API to handle requests
    //loadGraphAPI()
}