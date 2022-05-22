/*
 * File:    grapigo_storaging.go
 *
 * Author:      Ruymán Borges R. (ruyman21@gmail.com)
 * Date:        21-05-22
 * 
 * Version:     1.00
 * VersionDate: 21-05-22
 * 
 *
 * Summary of File:
 *
 *  This program stores and retrieve data structures into a localfile for persistance
 *  It works alongside grapigo_api and grapigo main programs
 *  
 *  Documentation is also provided to operate it.
 * 
 *  It runs on a unix system (i.e. ubuntu)
 *
 *  Copyright (C) Ruymán Borges Rodríguez - All Rights Reserved
 *  Unauthorized copying of this file, via any medium is strictly prohibited
 *  Proprietary and confidential
 *  Written by Ruymán Borges Rodríguez <ruyman21@gmail.com>, May 2022
 */

package grapigo_storaging

import {
    "io/ioutil"
    "os"
    "fmt"
    "log"
    "time"
}

var localStorageFilePath = './storage/answers.csv'

// checkData checks if there is a a file containing previously storaged data in a given path
//
//  Takes: 1 argument
//      1: path string
// 
//  Returns:
//      1: bool                 (true for valid data in a given path)
//      2: error                (error, if any, in file opening operation)
//
func checkData(path string) (bool, error) {
    // 

    f, err := os.Open(path)

    defer f.Close()

    if err != nil {
        return false, err
    }

    stat, _ := f.Stat()

    if (stat.Size() > 0) {
        return true, _
    }
    
    return false, _
}

// loadData loads a file from the localStorageFilePath path in local persistance storage
// 
//  Returns:
//          ioutil.ReadFile     (localStorageFilePath data as []byte)
//          error               (error if any, in ioutil.ReadFile)
//
func loadData() (ioutil.ReadFile, error) {
    // ... code that loads the storaged data in the 'localStorageFilePath' var path
    f, err := ioutil.ReadFile(localStorageFilePath)

    if err != nil {
        return _, err    
    }
    
    return f, _
}



func AnsByteToStruct(memAnswersByte []byte) Answers {

    // ... code that handles and process the bytes into Answers struct format

    //return answersStruct
}

// initData sets a base empty array of Answer struct inside the Answers struct dataset
func initData() {
    // ... code that inits a variable for the Answers struct
}

func Init() {
    // Check if there is any and in that case load previously storaged Answers into memory for this session

    pathExistsNotEmpty, err     := grapigo_storaging.checkData(localStorageFilePath)

    if (pathExistsNotEmpty) {
        memAnswersByte, err     := grapigo_storaging.loadData()

        if err != nil {
            memAnswers          := grapigo_storaging.byteAnsToStruct(memAnswersByte)    
        }
        
    } else {
        memAnswers, err         := initData()
    }

    if err != nil {
        // Handle preloading or init storaged answers error
        panic(err)
    }
}