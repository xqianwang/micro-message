#!/bin/bash

oc delete is micro-message
oc delete route micro-message
oc delete service micro-message
oc delete dc micro-message