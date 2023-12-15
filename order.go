package main

import (
  "fmt"
)

type OrderPriority int

const (
  OP_URGENT OrderPriority = iota + 1
  OP_LOW
)

func PriorityString(op OrderPriority) string {
  switch(op) {
  case OP_URGENT: 
  return "URGENT"
  case OP_LOW: 
  return "LOW"
  }

  return "SYSTEM ERROR - FATAL"
}

type Order struct {
  VerificationCode string
  Priority OrderPriority
}

func (o *Order) Generate() string {
  const PRIORITY_MARKER = "***"
  header := fmt.Sprintf("%s %s %s", 
    PRIORITY_MARKER, 
    PriorityString(o.Priority), 
    PRIORITY_MARKER)

  const ORDER_HEADER_MARK = "====="
  const NEW_LINE = "\n"
  const END_BODY = "\n\n\n\n"

  return header + END_BODY
}
