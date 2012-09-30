package main

// Response status codes
const SUCCESS_STATUS = 1
const FAILURE_STATUS = -1
const WARNING_STATUS = -2

// Response errors status codes
const TYPE_ERROR = "0"
const KEY_ERROR = "1"
const VALUE_ERROR = "2"
const INDEX_ERROR = "3"
const RUNTIME_ERROR = "4"
const OS_ERROR = "5"

// Signals
const BATCH_PUT_SIGNAL = "1"
const BATCH_DELETE_SIGNAL = "0"