#!/bin/bash
bash <<< "cat <<< "\""$(cat $@)"\"""
