/*
 * File:    grapigo_storaging_test.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	23-5-22
 *
 * Summary of File:
 *
 *  This program tests grapigo_storaging.go and its functions 
 *  
 *
 */

package grapigo_storaging

import (
    "testing"
    "os"
    "fmt"
    "io/fs"
)

var WhiteNotPassed  = "\033[97m NOT PASSED: "
var NotPassed       = "\033[31m NOT PASSED: \033[97m"
var Passed          = "\033[32m PASSED: \033[97m"
var TotalPassed     = 0
var TotalNotPassed  = 0